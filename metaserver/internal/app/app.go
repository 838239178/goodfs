package app

import (
	"common/graceful"
	"common/logs"
	"common/registry"
	"common/util"
	. "metaserver/config"
	"metaserver/internal/controller/grpc"
	"metaserver/internal/controller/http"
	"metaserver/internal/usecase/repo"
	"metaserver/internal/usecase/service"
	"metaserver/internal/usecase/pool"
)

func Run(cfg *Config) {
	// init logger
	logs.SetLevel(cfg.LogLevel)
	// init components
	pool.InitPool(cfg)
	defer pool.Close()
	// init services
	var grpcServer *grpc.RpcRaftServer
	netAddr := util.GetHostPort(cfg.Port)
	metaRepo := repo.NewMetadataRepo(pool.Storage)
	metaService := service.NewMetadataService(metaRepo)
	grpcServer, pool.RaftWrapper = grpc.NewRpcRaftServer(cfg.Cluster, pool.Storage)
	httpServer := http.NewHttpServer(netAddr, metaService)
	defer registry.NewEtcdRegistry(pool.Etcd, cfg.Registry, netAddr).MustRegister().Unregister()
	graceful.ListenAndServe(httpServer, grpcServer)
}
