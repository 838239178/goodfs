package logic

import (
	"apiserver/internal/entity"
	"apiserver/internal/usecase/pool"
	"apiserver/internal/usecase/selector"
	"common/collection/set"
	"common/constrant"
	"common/util"
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type ipCache struct {
	ips       []string
	updatedAt int64
}

var (
	groupIPCache    = map[string]*ipCache{}
	expiredDuration = int64(time.Minute.Seconds())
)

type Discovery struct{}

func NewDiscovery() Discovery { return Discovery{} }

func (Discovery) GetDataServers() []string {
	return pool.Discovery.GetServices(pool.Config.Discovery.DataServName)
}

func (Discovery) GetMetaServers(master bool) []string {
	return pool.Discovery.GetServicesWith(pool.Config.Discovery.MetaServName, master)
}

func (Discovery) SelectMetaByGroupID(gid string, defLoc string) string {
	if cache, ok := groupIPCache[gid]; ok && time.Now().Unix()-cache.updatedAt < expiredDuration {
		return selector.NewIPSelector(pool.Balancer, cache.ips).Select()
	}
	resp, err := pool.Etcd.Get(context.Background(), constrant.EtcdPrefix.FmtPeersInfo(gid, ""), clientv3.WithPrefix())
	if err != nil || len(resp.Kvs) == 0 {
		delete(groupIPCache, gid)
		return defLoc
	}
	alive := set.OfString(pool.Discovery.GetServices(pool.Config.Discovery.MetaServName))
	ips := make([]string, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		var info entity.PeerInfo
		if err = util.DecodeMsgp(&info, kv.Value); err == nil {
			ip := fmt.Sprint(info.Location, ":", info.HttpPort)
			if alive.Contains(ip) {
				ips = append(ips, ip)
			}
		}
	}
	if len(ips) == 0 {
		delete(groupIPCache, gid)
		return defLoc
	}
	groupIPCache[gid] = &ipCache{ips, time.Now().Unix()}
	return selector.NewIPSelector(pool.Balancer, ips).Select()
}

func (d Discovery) SelectDataServer(sel selector.Selector, size int) []string {
	ds := d.GetDataServers()
	if len(ds) == 0 {
		return []string{}
	}
	serv := make([]string, size)
	lb := selector.IPSelector{Selector: sel, IPs: ds}
	for i := 0; i < size; i++ {
		serv[i] = lb.Select()
	}
	return serv
}
