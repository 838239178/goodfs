package app

import (
	. "apiserver/config"
	"apiserver/internal/controller/http"
	"apiserver/internal/usecase/pool"
	"apiserver/internal/usecase/repo"
	"apiserver/internal/usecase/service"
	"common/graceful"
	"common/logs"
	"common/registry"
	"common/util"
)

func Run(cfg *Config) {
	pool.InitPool(cfg)
	defer pool.Close()
	// init log
	logs.SetLevel(cfg.LogLevel)
	//init services
	versionRepo := repo.NewVersionRepo(pool.Etcd)
	metaRepo := repo.NewMetadataRepo(pool.Etcd, versionRepo)
	metaService := service.NewMetaService(metaRepo, versionRepo)
	objService := service.NewObjectService(metaService, pool.Etcd)
	// register
	cfg.Registry.HttpAddr = util.GetHostPort(cfg.Port)
	defer registry.NewEtcdRegistry(pool.Etcd, cfg.Registry).MustRegister().Unregister()
	//start api server
	graceful.ListenAndServe(
		http.NewHttpServer(cfg.Registry.HttpAddr, objService, metaService),
	)
}