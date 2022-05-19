package temp

import (
	"goodfs/lib/util"
	"goodfs/lib/util/cache"
	"goodfs/objectserver/global"
	"goodfs/objectserver/model"
	"goodfs/objectserver/service"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Patch(g *gin.Context) {
	id := g.Param("name")
	if e := service.PutFile(global.Config.TempPath, id, g.Request.Body); e != nil {
		util.AbortInternalError(g, e)
		return
	}
	//刷新缓存 防止失效
	global.Cache.Refresh(id)
	g.Status(http.StatusOK)
}

func Delete(g *gin.Context) {
	id := g.Param("name")
	defer global.Cache.Delete(id)
	g.Status(http.StatusOK)
}

func Post(g *gin.Context) {
	var req model.TempPostReq
	if e := req.Bind(g); e != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		return
	}
	tmpInfo := &model.TempInfo{Name: req.Name, Size: req.Size}
	tmpInfo.Id = model.TempKeyPrefix + uuid.NewString()
	if !global.Cache.SetGob(tmpInfo.Id, tmpInfo) {
		g.AbortWithStatus(http.StatusServiceUnavailable)
	}
	g.Status(http.StatusOK)
	_, _ = g.Writer.Write([]byte(tmpInfo.Id))
}

func Put(g *gin.Context) {
	id := g.Param("name")
	var ti *model.TempInfo
	var ok bool
	if ti, ok = cache.GetGob[model.TempInfo](global.Cache, id); ok {
		if e := service.MvTmpToStorage(id, ti.Name); e != nil {
			_ = g.AbortWithError(http.StatusServiceUnavailable, e)
			return
		}
		service.MarkExist(ti.Name)
	} else {
		g.JSON(http.StatusNotFound, gin.H{"msg": "Temp file has been removed"})
		return
	}
	g.Status(http.StatusOK)
}

//Head 获取分片临时对象的大小
func Head(g *gin.Context) {
	id := g.Param("name")
	s, e := os.Stat(global.Config.TempPath + id)
	if e != nil {
		g.Status(http.StatusNotFound)
	} else {
		g.Header("Size", util.NumToString(s.Size()))
	}
}

//Get 获取临时对象分片
func Get(g *gin.Context) {
	if e := service.GetTemp(g.Param("name"), g.Writer); e != nil {
		log.Println(e)
		g.Status(http.StatusNotFound)
	}
}

func HandleTempRemove(ch <-chan cache.CacheEntry) {
	log.Println("Start handle temp file removal..")
	for entry := range ch {
		if strings.HasPrefix(entry.Key, model.TempKeyPrefix) {
			var ti model.TempInfo
			if ok := util.GobDecodeGen2(entry.Value, &ti); ok {
				if e := service.DeleteFile(global.Config.TempPath, ti.Id); e != nil {
					log.Printf("Remove temp %v(name=%v) error, %v", ti.Id, ti.Name, e)
				}
			} else {
				log.Printf("Handle evicted key=%v error, value cannot cast to TempInfo", entry.Key)
			}
		}
	}
}
