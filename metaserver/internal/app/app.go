package app

import (
	"common/logs"
	"common/registry"
	"common/util"
	"fmt"
	. "metaserver/config"
	"metaserver/internal/controller/http"
	"metaserver/internal/usecase/repo"
	"metaserver/internal/usecase/service"
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"metaserver/internal/controller/grpc"
)

var logger = logs.Std()

func Run(cfg *Config) {
	// init logger
	logs.SetLevel(cfg.LogLevel)
	// open db file
	boltdb, err := bolt.Open(cfg.DataDir, os.ModePerm, &bolt.Options{
		Timeout:    12 * time.Second,
		NoGrowSync: false,
	})
	if err != nil {
		logger.Errorf("open db err: %v", err)
		return
	}
	// init components
	etcdCli, err := clientv3.New(clientv3.Config{
		Endpoints: cfg.Etcd.Endpoint,
		Username:  cfg.Etcd.Username,
		Password:  cfg.Etcd.Password,
	})
	if err != nil {
		logger.Errorf("create etcd client err: %v", err)
		return
	}
	netAddr := fmt.Sprint(util.GetHost(), ":", cfg.Port)
	metaRepo := repo.NewMetadataRepo(boltdb)
	metaService := service.NewMetadataService(metaRepo)
	grpcServer := grpc.NewRpcRaftServer(cfg.Cluster, metaRepo.DB)
	metaRepo.Raft = grpcServer.Raft
	httpServer := http.NewHttpServer(netAddr, grpcServer.Server, metaService)
	defer registry.NewEtcdRegistry(etcdCli, cfg.Registry, netAddr).MustRegister().Unregister()

	httpServer.ListenAndServe()
}
