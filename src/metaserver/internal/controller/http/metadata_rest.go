package http

import (
	"common/proto/msg"
	"common/response"
	"common/util"
	"metaserver/internal/usecase"
	. "metaserver/internal/usecase"
	"sort"

	"github.com/gin-gonic/gin"
)

type MetadataController struct {
	service IMetadataService
}

func NewMetadataController(service IMetadataService) *MetadataController {
	return &MetadataController{service}
}

func (m *MetadataController) RegisterRoute(engine gin.IRouter) {
	engine.PUT("/metadata/:name", m.Put)
	engine.POST("/metadata/:name", m.Post)
	engine.GET("/metadata/:name", m.Get)
	engine.DELETE("/metadata/:name", m.Delete)
	engine.GET("/metadata/list", m.List)
}

func (m *MetadataController) Post(g *gin.Context) {
	var data msg.Metadata
	if err := g.ShouldBindJSON(&data); err != nil {
		response.FailErr(err, g)
		return
	}
	if err := m.service.AddMetadata(g.Param("name"), &data); err != nil {
		response.FailErr(err, g)
		return
	}
	response.Created(g)
}

func (m *MetadataController) Put(g *gin.Context) {
	var data msg.Metadata
	_ = g.ShouldBindJSON(&data)
	if err := m.service.UpdateMetadata(g.Param("name"), &data); err != nil {
		response.FailErr(err, g)
		return
	}
	response.Ok(g)
}

func (m *MetadataController) Get(g *gin.Context) {
	qry := struct {
		Version   int  `form:"version"`
		WithExtra bool `form:"with_extra"`
	}{}
	if err := g.ShouldBindQuery(&qry); err != nil {
		response.FailErr(err, g)
		return
	}
	meta, vers, err := m.service.GetMetadata(g.Param("name"), qry.Version, qry.WithExtra)
	if err != nil {
		response.FailErr(err, g)
		return
	}
	var versionList []*msg.Version
	if vers != nil {
		versionList = append(versionList, vers)
	}
	// metadata and version format
	response.OkJson(struct {
		*msg.Metadata
		Versions []*msg.Version `json:"versions,omitempty"`
	}{meta, versionList}, g)
}

func (m *MetadataController) Delete(g *gin.Context) {
	err := m.service.RemoveMetadata(g.Param("name"))
	if err != nil {
		response.FailErr(err, g)
		return
	}
	response.NoContent(g)
}

func (m *MetadataController) List(c *gin.Context) {
	req := struct {
		Prefix   string `form:"prefix"`
		PageSize int    `form:"page_size" binding:"required,lte=10000"`
		OrderBy  string `form:"order_by"`
		Desc     bool   `form:"desc"`
	}{}
	if err := c.ShouldBindQuery(&req); err != nil {
		response.FailErr(err, c)
		return
	}
	res, total, err := m.service.ListMetadata(req.Prefix, req.PageSize)
	if usecase.IsNotFound(err) {
		response.OkJson([]struct{}{}, c)
		return
	}
	if err != nil {
		response.FailErr(err, c)
		return
	}
	sort.Slice(res, func(i, j int) bool {
		var b bool
		switch req.OrderBy {
		default:
			fallthrough
		case "create_time":
			b = res[i].CreateTime < res[j].CreateTime
		case "update_time":
			b = res[i].UpdateTime < res[j].UpdateTime
		case "name":
			b = res[i].Name < res[j].Name
		}
		return util.IfElse(req.Desc, !b, b)
	})
	response.Exec(c).
		Header(gin.H{"X-Total-Count": total}).
		JSON(res)
}
