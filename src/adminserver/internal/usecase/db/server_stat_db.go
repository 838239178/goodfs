package db

import (
	"bytes"
	"common/cst"
	"common/graceful"
	"common/logs"
	"common/system"
	"common/util"
	"common/util/slices"
	"context"
	"math"
	"sync"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	statLog     = logs.New("stat-db")
	maxTimeStat = 59
)

type TimeStat struct {
	Time    time.Time `json:"time"`
	Percent float64   `json:"percent"`
}

type statTimeline struct {
	CpuTimeline []*TimeStat
	MemTimeline []*TimeStat
}

func newStatTimeline() *statTimeline {
	return &statTimeline{
		CpuTimeline: make([]*TimeStat, 0, 60),
		MemTimeline: make([]*TimeStat, 0, 60),
	}
}

func (st *statTimeline) Append(cpu *TimeStat, mem *TimeStat) {
	// limit size
	if len(st.CpuTimeline) == maxTimeStat {
		slices.RemoveFirst(&st.CpuTimeline)
	}
	if len(st.MemTimeline) == maxTimeStat {
		slices.RemoveFirst(&st.MemTimeline)
	}
	// add new stat
	st.CpuTimeline = append(st.CpuTimeline, cpu)
	st.MemTimeline = append(st.MemTimeline, mem)
}

type ServerStatCli struct {
	clientv3.Watcher
	clientv3.KV
}

type ServerStatDB struct {
	Cli       ServerStatCli
	GroupName string
	Services  []string
	closeFn   func()
	timeline  map[string]map[string]*statTimeline
}

func NewServerStatDB(cli ServerStatCli, groupName string, services []string) *ServerStatDB {
	o := &ServerStatDB{
		Cli:       cli,
		GroupName: groupName,
		Services:  services,
		timeline:  map[string]map[string]*statTimeline{},
	}
	o.init()
	return o
}

func (sdb *ServerStatDB) GetTimeline(servName string) map[string]*statTimeline {
	if tls, ok := sdb.timeline[servName]; ok {
		return tls
	}
	return map[string]*statTimeline{}
}

func (sdb *ServerStatDB) init() {
	ctx, cancel := context.WithCancel(context.Background())
	sdb.closeFn = cancel
	var mux sync.Mutex
	for _, name := range sdb.Services {
		go func(v string) {
			defer graceful.Recover()
			prefix := cst.EtcdPrefix.FmtSystemInfo(sdb.GroupName, v, "")
			// init value
			res, err := sdb.Cli.Get(ctx, prefix, clientv3.WithPrefix())
			if err != nil {
				statLog.Errorf("init stat of %s fail, %s", v, err)
			}
			for _, kv := range res.Kvs {
				mux.Lock()
				if err := sdb.addStat(v, kv.Key, kv.Value); err != nil {
					statLog.Error(err)
				}
				mux.Unlock()
			}
			// watch channel
			ch := sdb.Cli.Watch(ctx, prefix, clientv3.WithPrefix())
			sdb.watching(v, ch)
		}(name)
	}
}

func (sdb *ServerStatDB) addStat(serv string, key, value []byte) error {
	ts := time.Now()
	mp, ok := sdb.timeline[serv]
	if !ok {
		mp = map[string]*statTimeline{}
		sdb.timeline[serv] = mp
	}
	idx := bytes.LastIndex(key, cst.EtcdPrefix.Sep)
	id := string(key[idx+1:])
	var sysInfo system.Info
	if err := util.DecodeMsgp(&sysInfo, value); err != nil {
		return err
	}
	tl, ok := mp[id]
	if !ok {
		mp[id] = newStatTimeline()
		tl = mp[id]
	}
	tl.Append(&TimeStat{
		Time:    ts,
		Percent: sysInfo.CpuStatus.UsedPercent,
	}, &TimeStat{
		Time:    ts,
		Percent: math.Ceil(float64(sysInfo.MemStatus.Used)*100/float64(sysInfo.MemStatus.All)) / 100,
	})
	return nil
}

func (sdb *ServerStatDB) watching(serv string, ch clientv3.WatchChan) {
	for v := range ch {
		for _, event := range v.Events {
			if event.Type != mvccpb.PUT {
				continue
			}
			if err := sdb.addStat(serv, event.Kv.Key, event.Kv.Value); err != nil {
				statLog.Error(err)
			}
		}
	}
}

func (sdb *ServerStatDB) Close() error {
	sdb.closeFn()
	return nil
}
