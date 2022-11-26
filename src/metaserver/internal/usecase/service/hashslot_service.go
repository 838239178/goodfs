package service

import (
	"common/hashslot"
	"common/logs"
	"common/pb"
	"common/util"
	"context"
	"errors"
	"fmt"
	"metaserver/config"
	"metaserver/internal/entity"
	"metaserver/internal/usecase"
	"metaserver/internal/usecase/db"
	"metaserver/internal/usecase/pool"
	"strings"
	"time"

	"github.com/tinylib/msgp/msgp"
	"google.golang.org/grpc"
)

type HashSlotService struct {
	Store        *db.HashSlotDB
	Serivce      usecase.IMetadataService
	Cfg          *config.HashSlotConfig
	startReceive func()
}

func NewHashSlotService(st *db.HashSlotDB, serv usecase.IMetadataService, cfg *config.HashSlotConfig) *HashSlotService {
	return &HashSlotService{
		Store:        st,
		Serivce:      serv,
		Cfg:          cfg,
		startReceive: func() {},
	}
}

func (h *HashSlotService) OnLeaderChanged(isLeader bool) {
	if isLeader {
		var (
			info  *hashslot.SlotInfo
			exist bool
			err   error
		)
		// if not exist, read from configuration
		if info, exist, err = h.Store.Get(h.Cfg.StoreID); !exist {
			if err != nil {
				logs.Std().Errorf("get slot-info when leader changed: %s", err)
				return
			}
			info = &hashslot.SlotInfo{Slots: h.Cfg.Slots, GroupID: h.Cfg.StoreID}
			logs.Std().Infof("no exist slots, init from config: id=%s, slots=%s", info.GroupID, info.Slots)
		}
		info.Location = pool.HttpHostPort
		if err := h.Store.Save(h.Cfg.StoreID, info); err != nil {
			logs.Std().Error(err)
			return
		}
	}
}

func (h *HashSlotService) GetCurrentSlots(reload bool) (map[string][]string, error) {
	prov, err := h.Store.GetEdgeProvider(reload)
	if err != nil {
		return nil, err
	}
	res := make(map[string][]string)
	for _, v := range hashslot.CopyOfEdges("", prov) {
		res[v.Value] = append(res[v.Value], fmt.Sprint(v.Start, "-", v.End))
	}
	return res, nil
}

func (h *HashSlotService) PrepareMigrationTo(loc *pb.LocationInfo, slots []string) error {
	// validate slots
	provider, err := h.Store.GetEdgeProvider(false)
	if err != nil {
		return err
	}
	edges, err := hashslot.WrapSlotsToEdges(slots, pool.HttpHostPort)
	if err != nil {
		return err
	}
	// ensure all slots is in this server
	for _, edge := range edges {
		if !hashslot.IsValidEdge(edge, provider) {
			return fmt.Errorf("slot %s is not in this server", edge)
		}
	}
	// send prepare rpc to target
	cc, err := grpc.Dial(fmt.Sprint(loc.GetHost(), ":", loc.GetRpcPort()), grpc.WithInsecure())
	if err != nil {
		return err
	}
	client := pb.NewHashSlotClient(cc)
	resp, err := client.PrepareMigration(context.Background(), &pb.PrepareReq{
		Id: h.Cfg.StoreID,
		Location: &pb.LocationInfo{
			Host:     util.GetHost(),
			RpcPort:  pool.Config.Cluster.Port,
			HttpPort: pool.Config.Port,
		},
		Slots: slots,
	})
	if err != nil {
		return err
	}
	if !resp.GetSuccess() {
		return errors.New(resp.GetMessage())
	}
	// change status to migrate-to
	return h.Store.ReadyMigrateTo(loc.GetHost(), slots)
}

// PrepareMigrationFrom Change into migrate-from. Status will change back if timeout
func (h *HashSlotService) PrepareMigrationFrom(loc *pb.LocationInfo, slots []string) error {
	// validate slots
	provider, err := h.Store.GetEdgeProvider(false)
	if err != nil {
		return err
	}
	edges, err := hashslot.WrapSlotsToEdges(slots, pool.HttpHostPort)
	if err != nil {
		return err
	}
	// ensure all slots is not in this server
	for _, edge := range edges {
		if hashslot.IsValidEdge(edge, provider) {
			return fmt.Errorf("slot %s is currently in this server", edge)
		}
	}
	// change status to migrate-from
	if err := h.Store.ReadyMigrateFrom(loc.GetHost(), slots); err != nil {
		return err
	}
	cancelCh := make(chan struct{})
	h.startReceive = func() {
		close(cancelCh)
		h.startReceive = func() {}
	}
	go func() {
		select {
		case <-cancelCh:
			logs.Std().Debug("migration-from timeout-ctx canceled")
		case <-time.NewTicker(h.Cfg.PrepareTimeout).C:
			h.startReceive()
			_ = h.Store.FinishMigrateFrom()
			logs.Std().Errorf("timeout migrating from %s", loc.GetHost())
		}
	}()
	return nil
}

func (h *HashSlotService) FinishReceiveItem(success bool) error {
	h.startReceive()
	var (
		newSlots []string
		fromHost string
		ok       bool
	)
	if ok, fromHost, newSlots = h.Store.GetMigrateFrom(); !ok {
		return fmt.Errorf("get received slots fails: server is not in migrate-from")
	}
	if err := h.Store.FinishMigrateFrom(); err != nil {
		util.LogErr(err)
	}
	if !success {
		return nil
	}
	info, _, err := h.Store.Get(h.Cfg.StoreID)
	if err != nil {
		return fmt.Errorf("update slot fails after finish migration: %w", err)
	}
	// combine to new slots, ignore error because both have been validated
	curEdges, _ := hashslot.WrapSlotsToEdges(info.Slots, info.Location)
	newEdges, _ := hashslot.WrapSlotsToEdges(newSlots, info.Location)
	info.Slots = hashslot.CombineEdges(curEdges, newEdges).Strings()
	// save new slot-info
	if err = h.Store.Save(h.Cfg.StoreID, info); err != nil {
		return fmt.Errorf("save new slot-info fails after finsih migrateion: %w", err)
	}
	logs.Std().Debugf("finish migration from %s success", fromHost)
	return nil
}

func (h *HashSlotService) ReceiveItem(item *pb.MigrationItem) error {
	h.startReceive()
	logData := &entity.RaftData{
		Version:  &entity.Version{},
		Metadata: &entity.Metadata{},
	}
	if err := util.DecodeMsgp(
		util.IfElse[msgp.Unmarshaler](item.IsVersion, logData.Version, logData.Metadata),
		item.Data,
	); err != nil {
		return err
	}
	if item.IsVersion {
		if err := h.Serivce.ReceiveVersion(item.Name, logData.Version); err != nil {
			if errors.Is(err, usecase.ErrExists) {
				return nil
			}
			return err
		}
	} else if err := h.Serivce.AddMetadata(logData.Metadata); err != nil {
		if errors.Is(err, usecase.ErrExists) {
			return nil
		}
		return err
	}
	return nil
}

// AutoMigrate migrate data
//TODO(perf): multi goroutine
func (h *HashSlotService) AutoMigrate(toLoc *pb.LocationInfo, slots []string) error {
	logger := logs.New("hash-slot-migration")
	if ok, host, _ := h.Store.GetMigrateTo(); !ok || host != toLoc.GetHost() {
		return fmt.Errorf("no ready to migrate to %s", toLoc.GetHost())
	}
	// connect to target
	cc, err := grpc.Dial(fmt.Sprint(toLoc.Host, ":", toLoc.RpcPort), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer cc.Close()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := pb.NewHashSlotClient(cc).StreamingReceive(ctx)
	if err != nil {
		return err
	}
	delEdges, _ := hashslot.WrapSlotsToEdges(slots, "")
	migKeys := h.Serivce.FilterKeys(func(s string) bool {
		return hashslot.IsSlotInEdges(hashslot.CalcBytesSlot(util.StrToBytes(s)), delEdges)
	})
	logger.Debugf("migration keys:\n %v", migKeys)
	var errs []error
	var sucNum int
	for _, key := range migKeys {
		// get data
		data, err := h.Serivce.GetMetadataBytes(key)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		// send data
		if err := stream.Send(&pb.MigrationItem{
			Name: key,
			Data: data,
		}); err != nil {
			logger.Debugf("send metadata %s err: %s", key, err)
			errs = append(errs, err)
			continue
		}
		// recv response
		resp, err := stream.Recv()
		if err != nil {
			logger.Debugf("recv send-metadata %s response err: %s", key, err)
			errs = append(errs, err)
			continue
		} else if !resp.Success {
			logger.Debugf("send-metadata %s recv failure resposne: %s", key, resp.Message)
			errs = append(errs, errors.New(resp.Message))
			continue
		}
		// start send versions
		allVersionSuccess := true
		h.Serivce.ForeachVersionBytes(key, func(b []byte) bool {
			// send version
			if err := stream.Send(&pb.MigrationItem{
				Name:      key,
				Data:      b,
				IsVersion: true,
			}); err != nil {
				errs = append(errs, err)
				allVersionSuccess = false
				logger.Debugf("send-metadata-version %s err: %s", key, err)
			}
			// recv response
			resp, err := stream.Recv()
			if err != nil {
				errs = append(errs, err)
				allVersionSuccess = false
				logger.Debugf("send-metadata-version %s recv err: %s", key, err)
			} else if !resp.Success {
				errs = append(errs, errors.New(resp.Message))
				allVersionSuccess = false
				logger.Debugf("send-metadata-version %s recv failure resposne: %s", key, resp.Message)
			}
			return true
		})
		// delete if all success
		if allVersionSuccess {
			if err := h.Serivce.RemoveMetadata(key); err != nil {
				errs = append(errs, err)
				logger.Debugf("delete-metadata %s err: %s", key, err)
			} else {
				sucNum++
			}
		}
	}
	if err := h.Store.FinishMigrateTo(); err != nil {
		errs = append(errs, err)
		logger.Debugf("switch status to normal err: %s", err)
	}
	logger.Infof("migration totally %d metadata and successed %d verions", len(migKeys), sucNum)
	if len(errs) > 0 {
		sb := strings.Builder{}
		sb.WriteString("occurred errors:")
		for _, err := range errs {
			sb.WriteRune('\n')
			sb.WriteString(err.Error())
		}
		logger.Error(sb.String())
		return errors.New("migrate partly fails, please retry again")
	}
	// all migrate success
	// remove slots from current slot-info
	info, _, err := h.Store.Get(h.Cfg.StoreID)
	if err != nil {
		return fmt.Errorf("update slot fails after finish migration: %w", err)
	}
	curEdges, _ := hashslot.WrapSlotsToEdges(info.Slots, info.Location)
	info.Slots = hashslot.RemoveEdges(curEdges, delEdges).Strings()
	// save new slot-info
	if err = h.Store.Save(h.Cfg.StoreID, info); err != nil {
		return fmt.Errorf("save new slot-info fails after finsih migrateion: %w", err)
	}
	// close stream as success
	if err := stream.CloseSend(); err != nil {
		logger.Error(err)
	}
	logger.Infof("finish migration to %s success", toLoc.Host)
	return nil
}
