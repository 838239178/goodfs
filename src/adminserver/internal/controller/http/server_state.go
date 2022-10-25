package http

import (
	"adminserver/internal/entity"
	"adminserver/internal/usecase/logic"
	"adminserver/internal/usecase/pool"
	"common/response"
	"common/util"

	"github.com/gin-gonic/gin"
)

type ServerStateController struct {
}

func NewServerStateController() *ServerStateController {
	return &ServerStateController{}
}

func (ss *ServerStateController) Register(r gin.IRouter) {
	r.Group("server").GET("/stat", ss.Stat)
}

func (ss *ServerStateController) Stat(c *gin.Context) {
	monitor := logic.NewServerMonitor()
	dg := util.NewDoneGroup()
	dg.Todo()
	var info [2]map[string]*entity.ServerInfo
	go func () {
		defer dg.Done()
		metaInfo, err := monitor.ServerStat(pool.Config.Discovery.MetaServName)
		if err != nil {
			dg.Error(err)
			return
		}
		info[0] = metaInfo
	}()
	dg.Todo()
	go func () {
		defer dg.Done()
		dataInfo, err := monitor.ServerStat(pool.Config.Discovery.DataServName)
		if err != nil {
			dg.Error(err)
			return
		}
		info[1] = dataInfo
	}()
	if err := dg.WaitUntilError(); err != nil {
		response.FailErr(err, c)
		return
	}
	response.OkJson(gin.H{
		"metaServer": info[0],
		"dataServer": info[1],
	}, c)
}