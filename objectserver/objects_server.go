package main

import (
	"fmt"
	"goodfs/objectserver/config"
	"goodfs/objectserver/controller"
	"goodfs/objectserver/controller/heartbeat"
	"goodfs/objectserver/controller/locate"
	"goodfs/objectserver/controller/temp"
	"goodfs/objectserver/global"
	"goodfs/util/cache"
	"goodfs/util/datasize"
	"os"
	"strconv"

	"github.com/838239178/goodmq"
	"github.com/allegro/bigcache"
	"github.com/gin-gonic/gin"
)

func initialize() {
	//init amqp
	{
		hn, e := os.Hostname()
		if e != nil {
			panic(e)
		}
		config.LocalAddr = fmt.Sprintf("%v:%v", hn, config.Port)
		global.AmqpConnection = goodmq.NewAmqpConnection(config.AmqpAddress)
	}

	//init cache
	{
		cacheConf := bigcache.DefaultConfig(config.CacheTTL)
		cacheConf.CleanWindow = config.CacheCleanInterval
		cacheConf.HardMaxCacheSize = config.CacheMaxSizeMB
		cacheConf.Shards = ((config.CacheMaxSizeMB * datasize.MB) / config.CacheItemMaxSize).IntValue()
		global.Cache = cache.NewCache(cacheConf)
	}
}

func close() {
	global.AmqpConnection.Close()
	global.Cache.Close()
}

func main() {
	initialize()
	defer close()

	locate.SyncExistingFilter()

	go temp.HandleTempRemove(global.Cache.NotifyEvicted())
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()

	router := gin.Default()

	//init router
	controller.Router(router)

	router.Run(":" + strconv.Itoa(config.Port))
}
