package app

import (
	"common/cst"
	"common/graceful"
	"common/system"
	"common/util"
	"fmt"
	"metaserver/config"
	"metaserver/internal/controller/grpc"
	"metaserver/internal/controller/http"
	"metaserver/internal/usecase/logic"
	"metaserver/internal/usecase/pool"
	"metaserver/internal/usecase/raftimpl"
	"metaserver/internal/usecase/repo"
	"metaserver/internal/usecase/service"
)

func Run(cfg *config.Config) {
	// init components
	pool.InitPool(cfg)
	defer pool.Close()
	// init services
	metaRepo := repo.NewMetadataRepo(pool.Storage, repo.NewMetadataCacheRepo(pool.Cache))
	bucketRepo := repo.NewBucketRepo(pool.Storage, repo.NewBucketCacheRepo(pool.Cache))
	fsm := raftimpl.NewFSM(metaRepo, repo.NewBatchRepo(pool.Storage), bucketRepo, repo.NewBatchBucketRepo(pool.Storage), metaRepo)
	raftWrapper := raftimpl.NewRaft(util.ServerAddress(cfg.Port), cfg.Cluster, fsm)
	bucketServ := service.NewBucketService(bucketRepo, raftWrapper)
	metaService := service.NewMetadataService(
		metaRepo,
		repo.NewBatchRepo(pool.Storage),
		repo.NewHashIndexRepo(pool.Storage),
		raftWrapper,
	)
	hsService := service.NewHashSlotService(pool.HashSlot, metaService, bucketServ, &cfg.HashSlot)
	grpcServer := grpc.NewRpcServer(cfg.MaxConcurrentStreams, raftWrapper, metaService, hsService, bucketServ)
	httpServer := http.NewHttpServer(cfg.Port, grpcServer, metaService, bucketServ)
	// register on leader change
	raftWrapper.RegisterLeaderChangedEvent(hsService)
	raftWrapper.RegisterLeaderChangedEvent(logic.NewRegistry())
	raftWrapper.Init()
	pool.RaftWrapper = raftWrapper
	// unregister service
	defer pool.Registry.Unregister()
	// remove slots info from etcd if shutdown as a leader
	defer func() {
		if raftWrapper.IsLeader() || !raftWrapper.Enabled {
			util.LogErr(pool.HashSlot.Remove(cfg.HashSlot.StoreID))
		}
	}()
	// auto sync sys-info
	syncer := system.Syncer(pool.Etcd, fmt.Sprint(cst.EtcdPrefix.SystemInfo, "/", pool.Config.Registry.RegisterKey()), <-pool.Registry.LifecycleLease())
	defer syncer.StartAutoSave()()
	// flush config
	defer cfg.Persist()

	graceful.ListenAndServe(nil, httpServer)
}
