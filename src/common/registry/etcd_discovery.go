package registry

import (
	"common/constrant"
	"common/graceful"
	"common/util"
	"context"
	"strings"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdDiscovery struct {
	cli         *clientv3.Client
	group       string
	httpService map[string]*serviceList
	rpcService  map[string]*serviceList
}

func NewEtcdDiscovery(cli *clientv3.Client, cfg *Config) *EtcdDiscovery {
	hs := make(map[string]*serviceList)
	rs := make(map[string]*serviceList)
	d := &EtcdDiscovery{cli, cfg.Group, hs, rs}
	for _, s := range cfg.Services {
		d.initService(s)
	}
	return d
}

func (e *EtcdDiscovery) initService(serv string) {
	e.httpService[serv] = newServiceList()
	e.rpcService[serv] = newServiceList()
	go func() {
		defer graceful.Recover()
		// fetch kvs
		prefix := constrant.EtcdPrefix.FmtRegistry(e.group, serv)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		res, err := e.cli.Get(ctx, prefix, clientv3.WithPrefix())
		if err != nil {
			log.Warnf("discovery init service %s error: %s", prefix, err)
			return
		}
		// wrap kvs
		https := make(map[string]string)
		rpcs := make(map[string]string)
		for _, kv := range res.Kvs {
			value := RegisterValue(kv.Value)
			https[value.HttpAddr()] = string(kv.Key)
			rpcs[value.RpcAddr()] = string(kv.Key)
		}
		// init serv
		e.httpService[serv].replace(https)
		e.rpcService[serv].replace(rpcs)
		// start watch change
		ch := e.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
		e.asyncWatch(serv, ch)
	}()
}

func (e *EtcdDiscovery) asyncWatch(serv string, ch clientv3.WatchChan) {
	go func() {
		defer graceful.Recover()
		for res := range ch {
			for _, event := range res.Events {
				//Key will be like ${group}/${serv}/${id}_${slave/master}
				key := string(event.Kv.Key)
				addr := RegisterValue(event.Kv.Value)
				switch event.Type {
				case mvccpb.PUT:
					e.addService(serv, addr, key)
				case mvccpb.DELETE:
					e.removeService(serv, addr)
				}
			}
		}
	}()
}

func (e *EtcdDiscovery) GetServiceMapping(name string, rpc bool) map[string]string {
	res := make(map[string]string)
	service := util.IfElse(rpc, e.rpcService, e.httpService)
	if sl, ok := service[name]; ok {
		for k, v := range sl.copy() {
			sp := strings.Split(v, "/")
			sp = strings.Split(sp[len(sp)-1], "_")
			res[sp[0]] = k
		}
	}
	return res
}

func (e *EtcdDiscovery) GetServices(name string, rpc bool) []string {
	service := util.IfElse(rpc, e.rpcService, e.httpService)
	if sl, ok := service[name]; ok {
		return sl.list()
	}
	return []string{}
}

func (e *EtcdDiscovery) GetServiceByAddr(name, addr string, rpc bool) (id string, httpAddr string, rpcAddr string) {
	if rpc {
		id = e.rpcService[name].data[addr]
		rpcAddr = addr
		httpAddr = e.GetServiceMapping(name, false)[id]
	} else {
		id = e.httpService[name].data[addr]
		httpAddr = addr
		rpcAddr = e.GetServiceMapping(name, true)[id]
	}
	return
}

func (e *EtcdDiscovery) GetServiceMappingWith(name string, rpc bool, master bool) map[string]string {
	s := util.IfElse(master, "master", "slave")
	c := util.IfElse(master, "slave", "master")
	service := util.IfElse(rpc, e.rpcService, e.httpService)
	if sl, ok := service[name]; ok {
		res := make(map[string]string, len(sl.data))
		for k, v := range sl.copy() {
			sp := strings.Split(v, "/")
			sp = strings.Split(sp[len(sp)-1], "_")
			// if it doesn't have any suffix means standalone
			if strings.HasSuffix(v, s) || !strings.HasSuffix(v, c) {
				res[sp[0]] = k
			}
		}
		return res
	}
	return map[string]string{}
}

func (e *EtcdDiscovery) GetServicesWith(name string, rpc bool, master bool) []string {
	mp := e.GetServiceMappingWith(name, rpc, master)
	arr := make([]string, len(mp))
	for _, v := range mp {
		arr = append(arr, v)
	}
	return arr
}

func (e *EtcdDiscovery) addService(name string, value RegisterValue, key string) {
	h, r := value.Addr()
	e.httpService[name].add(h, key)
	e.rpcService[name].add(r, key)
}

func (e *EtcdDiscovery) removeService(name string, value RegisterValue) {
	h, r := value.Addr()
	e.httpService[name].remove(h)
	e.rpcService[name].remove(r)
}
