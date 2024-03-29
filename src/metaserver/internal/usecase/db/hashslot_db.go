package db

import (
	"common/hashslot"
	"common/logs"
	"common/util"
	"context"
	"errors"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"sync/atomic"
)

var (
	hsLog         = logs.New("hash-slot-db")
	MaxExpireUnix = int64(5 * time.Minute.Seconds())
)

const (
	StatusNormal int32 = 1 << iota
	StatusMigrateTo
	StatusMigrateFrom
	StatusClosed
)

func StatusDesc(status int32) (string, error) {
	var desc string
	switch status {
	case StatusMigrateFrom:
		desc = "migrate-from"
	case StatusMigrateTo:
		desc = "migrate-to"
	case StatusNormal:
		desc = "normal"
	default:
		return "", fmt.Errorf("no support status %d", status)
	}
	return desc, nil
}

type HashSlotDB struct {
	kv             clientv3.KV
	status         *atomic.Int32
	provider       *atomic.Pointer[hashslot.IEdgeProvider]
	updatedAt      int64
	migratingSlots []string
	migratingHost  string
	KeyPrefix      string
}

func NewHashSlotDB(keyPrefix string, kv clientv3.KV) *HashSlotDB {
	at := &atomic.Int32{}
	at.Store(StatusNormal)
	return &HashSlotDB{
		KeyPrefix: keyPrefix,
		kv:        kv,
		provider:  &atomic.Pointer[hashslot.IEdgeProvider]{},
		status:    at,
	}
}

func (h *HashSlotDB) IsNormal() bool {
	return h.status.Load() == StatusNormal
}

func (h *HashSlotDB) GetMigrateTo() (bool, string, []string) {
	return h.status.Load() == StatusMigrateTo, h.migratingHost, h.migratingSlots
}

func (h *HashSlotDB) GetMigrateFrom() (bool, string, []string) {
	return h.status.Load() == StatusMigrateFrom, h.migratingHost, h.migratingSlots
}

func (h *HashSlotDB) IsExpired() bool {
	return time.Now().Unix()-h.updatedAt > MaxExpireUnix
}

// GetEdgeProvider identify is http-location of server
func (h *HashSlotDB) GetEdgeProvider(reload bool) (hashslot.IEdgeProvider, error) {
	item := h.provider.Load()
	if h.IsExpired() || reload {
		if err := h.reloadProvider(item); err != nil {
			return nil, err
		}
	}
	return *h.provider.Load(), nil
}

func (h *HashSlotDB) reloadProvider(old *hashslot.IEdgeProvider) error {
	slotsMap := make(map[string][]string)
	// get slots data from etcd (only master saves into to etcd)
	res, err := h.kv.Get(context.Background(), h.KeyPrefix, clientv3.WithPrefix())
	if err != nil {
		return fmt.Errorf("reloadProvider: %w", err)
	}
	// wrap slot
	for _, kv := range res.Kvs {
		var info hashslot.SlotInfo
		if err := util.DecodeMsgp(&info, kv.Value); err != nil {
			return fmt.Errorf("reloadProvider: %w", err)
		}
		slotsMap[info.ServerID] = info.Slots
	}
	data, err := hashslot.WrapSlots(slotsMap)
	if err != nil {
		return fmt.Errorf("reloadProvider: %w", err)
	}
	if h.provider.CompareAndSwap(old, &data) {
		h.updatedAt = time.Now().Unix()
		hsLog.Infof("update hash-slots success at %d", h.updatedAt)
	}
	return nil
}

func (h *HashSlotDB) ReadyMigrateFrom(loc string, slots []string) error {
	if h.status.CompareAndSwap(StatusNormal, StatusMigrateFrom) {
		h.migratingHost = loc
		h.migratingSlots = slots
		hsLog.Debugf("switch normal to migrate-from")
		return nil
	}
	return errors.New("status is not in normal")
}

func (h *HashSlotDB) ReadyMigrateTo(loc string, slots []string) error {
	if h.status.CompareAndSwap(StatusNormal, StatusMigrateTo) {
		h.migratingHost = loc
		h.migratingSlots = slots
		hsLog.Debugf("switch normal to migrate-to")
		return nil
	}
	return errors.New("status is not in normal")
}

func (h *HashSlotDB) FinishMigrateTo() error {
	if h.status.CompareAndSwap(StatusMigrateTo, StatusNormal) {
		h.migratingHost = ""
		h.migratingSlots = nil
		hsLog.Debugf("switch migrate-to to normal")
		return nil
	}
	return fmt.Errorf("status is not in migrate-to")
}

func (h *HashSlotDB) FinishMigrateFrom() error {
	if h.status.CompareAndSwap(StatusMigrateFrom, StatusNormal) {
		h.migratingHost = ""
		h.migratingSlots = nil
		hsLog.Debugf("switch migrate-from to normal")
		return nil
	}
	return fmt.Errorf("status is not in migrate-from")
}

// Get The 'id' is the store id which defined in configuration
func (h *HashSlotDB) Get(id string) (*hashslot.SlotInfo, bool, error) {
	key := fmt.Sprint(h.KeyPrefix, id)
	resp, err := h.kv.Get(context.Background(), key)
	if err != nil {
		return nil, false, err
	}
	if len(resp.Kvs) == 0 {
		return nil, false, nil
	}
	var info hashslot.SlotInfo
	if err = util.DecodeMsgp(&info, resp.Kvs[0].Value); err != nil {
		return nil, false, err
	}
	return &info, true, nil
}

func (h *HashSlotDB) Save(id string, info *hashslot.SlotInfo) (err error) {
	if h.status.Load() != StatusNormal {
		return errors.New("status not in normal")
	}
	key := fmt.Sprint(h.KeyPrefix, id)
	bt, err := util.EncodeMsgp(info)
	if err != nil {
		return err
	}
	info.GroupID = id
	// saving
	_, err = h.kv.Put(context.Background(), key, string(bt))
	return err
}

func (h *HashSlotDB) Remove(id string) error {
	if h.status.Load() != StatusNormal {
		return errors.New("status not in normal")
	}
	key := fmt.Sprint(h.KeyPrefix, id)
	_, err := h.kv.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	return nil
}

func (h *HashSlotDB) GetKeyIdentify(key string) (string, error) {
	slots, err := h.GetEdgeProvider(false)
	if err != nil {
		return "", err
	}
	// get slot's location of this key
	location, err := hashslot.GetStringIdentify(key, slots)
	if err != nil {
		return "", err
	}
	return location, nil
}

func (h *HashSlotDB) Close(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	for !h.status.CompareAndSwap(StatusNormal, StatusClosed) {
		select {
		case <-ctx.Done():
			desc, _ := StatusDesc(h.status.Load())
			return fmt.Errorf("close fail: db stay in %s timeout", desc)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
	return nil
}
