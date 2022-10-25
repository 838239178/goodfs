package logic

import (
	"common/constrant"
	"common/graceful"
	"common/logs"
	"common/system"
	"common/util"
	"context"
	"metaserver/internal/usecase/pool"
	"time"
)

type SystemStatLogic struct {
}

func NewDiskLogic() *SystemStatLogic {
	return &SystemStatLogic{}
}

func (d SystemStatLogic) StartAutoSave() func() {
	ctx, cancel := context.WithCancel(context.Background())
	tk := time.NewTicker(time.Minute)
	go func() {
		defer graceful.Recover()
		for {
			select {
			case <-ctx.Done():
				logs.Std().Info("stop auto save sys-info")
				return
			case <-tk.C:
				util.LogErrWithPre("auto save sys-info", d.Save())
			}
		}
	}()
	return cancel
}

func (d SystemStatLogic) Save() error {
	info, err := system.SystemInfo(`/`)
	if err != nil {
		return err
	}
	bt, err := util.EncodeMsgp(info)
	if err != nil {
		return err
	}
	keyDisk := constrant.EtcdPrefix.FmtSystemInfo(pool.Config.Registry.Group, pool.Config.Registry.Name, pool.Config.Registry.ServerID)
	_, err = pool.Etcd.Put(context.Background(), keyDisk, string(bt))
	return err
}