package http

import (
	"apiserver/internal/entity"
	"apiserver/internal/usecase"
	"apiserver/internal/usecase/logic"
	"apiserver/internal/usecase/webapi"
	"common/response"
	"github.com/gin-gonic/gin"
)

type MetadataController struct {
	Service usecase.IMetaService
}

func NewMetadataController(service usecase.IMetaService) *MetadataController {
	return &MetadataController{Service: service}
}

func (mc *MetadataController) Register(route gin.IRouter) {
	route.Group("metadata").
		GET("/:name/versions", mc.Versions).
		GET("/:name", mc.Get)
}

func (mc *MetadataController) Get(c *gin.Context) {
	body := struct {
		entity.Binder
		Name    string `uri:"name"`
		Version int32  `form:"json"`
	}{}
	if err := body.Bind(c, false); err != nil {
		response.FailErr(err, c)
		return
	}
	data, err := mc.Service.GetMetadata(body.Name, body.Version)
	if err != nil {
		response.FailErr(err, c)
		return
	}
	response.OkJson(data, c)
}

func (mc *MetadataController) Versions(c *gin.Context) {
	body := struct {
		entity.Binder
		Page     int    `form:"page" binding:"required"`
		PageSize int    `form:"page_size" binding:"required"`
		Name     string `uri:"name"`
	}{}
	if err := body.Bind(c, false); err != nil {
		response.FailErr(err, c)
		return
	}
	loc, gid, err := logic.NewHashSlot().FindMetaLocOfName(body.Name)
	if err != nil {
		response.FailErr(err, c)
		return
	}
	loc = logic.NewDiscovery().SelectMetaByGroupID(gid, loc)
	version, err := webapi.ListVersion(loc, body.Name, body.Page, body.PageSize)
	if _, err := c.Writer.Write(version); err != nil {
		response.FailErr(err, c)
		return
	}
}