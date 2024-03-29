package http

import (
	"apiserver/internal/entity"
	"apiserver/internal/usecase"
	"common/datasize"
	"common/logs"
	"common/response"
	"common/util"
	"github.com/gin-gonic/gin"
	"io"
)

type ObjectsController struct {
	objectService usecase.IObjectService
	metaService   usecase.IMetaService
}

func NewObjectsController(obj usecase.IObjectService, meta usecase.IMetaService) *ObjectsController {
	return &ObjectsController{obj, meta}
}

func (oc *ObjectsController) Register(r gin.IRoutes) {
	r.PUT("/objects/:name", oc.ValidatePut, oc.Put)
	r.GET("/objects/:name", oc.Get)
}

func (oc *ObjectsController) Put(c *gin.Context) {
	req := c.Value("PutReq").(*entity.PutReq)
	req.Body = c.Request.Body
	if c.Request.ContentLength <= 0 {
		response.BadRequestMsg("content-length invalid", c)
		return
	}
	verNum, err := oc.objectService.StoreObject(req, &entity.Metadata{
		Name:   req.Name,
		Bucket: req.Bucket,
		Versions: []*entity.Version{{
			Size:          c.Request.ContentLength,
			Hash:          req.Hash,
			StoreStrategy: req.Store,
			Compress:      req.Compress,
		}},
	})

	if err != nil {
		response.FailErr(err, c)
		return
	}

	response.CreatedJson(&entity.PutResp{
		Name:    req.Name,
		Bucket:  req.Bucket,
		Version: verNum,
	}, c)
}

func (oc *ObjectsController) Get(c *gin.Context) {
	var req entity.GetReq
	if e := req.Bind(c); e != nil {
		response.BadRequestErr(e, c)
		return
	}
	// get metadata
	metaData, err := oc.metaService.GetMetadata(req.Name, req.Bucket, req.Version, false)
	if err != nil {
		response.FailErr(err, c).Abort()
		return
	}
	// get object stream
	stream, err := oc.objectService.GetObject(metaData, metaData.Versions[0])
	if err != nil {
		response.FailErr(err, c).Abort()
		return
	}
	defer util.CloseAndLog(stream)
	// try seek
	if tp, ok := req.Range.GetFirstBytes(); ok {
		if _, err = stream.Seek(tp.First, io.SeekCurrent); err != nil {
			response.FailErr(err, c)
			return
		}
	}
	// copy to response
	n, err := io.Copy(c.Writer, stream)
	if err != nil {
		logs.Std().Error(err)
		return
	}
	response.OkHeader(gin.H{
		"Accept-Ranges":  "bytes",
		"Content-Length": n,
	}, c)
}

func (oc *ObjectsController) ValidatePut(g *gin.Context) {
	var req entity.PutReq
	if err := req.Bind(g); err != nil {
		response.BadRequestErr(err, g).Abort()
		return
	}
	if g.Request.ContentLength <= 0 {
		response.BadRequestMsg("empty request", g).Abort()
		return
	}
	if req.Store == 0 {
		if g.Request.ContentLength > int64(datasize.KB*64) {
			req.Store = entity.ECReedSolomon
		} else {
			req.Store = entity.MultiReplication
		}
	}
	if ext, ok := util.GetFileExt(req.Name, false); ok {
		req.Ext = ext
	} else {
		req.Ext = "bytes"
	}
	g.Set("PutReq", &req)
}
