package cache

import (
	"common/graceful"
	"common/util"
	"log"

	"github.com/allegro/bigcache"
)

type Cache struct {
	cache         *bigcache.BigCache
	notifyEvicted []chan Entry
}

type Entry struct {
	Key    string
	Value  []byte
	Reason bigcache.RemoveReason
}

func NewCache(config bigcache.Config) *Cache {
	res := &Cache{notifyEvicted: make([]chan Entry, 0, 16)}
	config.OnRemoveWithReason = res.onRemove
	b, e := bigcache.NewBigCache(config)
	if e != nil {
		panic(e)
	}
	res.cache = b
	return res
}

func GetGob[T interface{}](c ICache, k string) (*T, bool) {
	if bt := c.Get(k); bt != nil {
		return util.GobDecodeGen[T](bt)
	}
	return nil, false
}

func GetGob2[T interface{}](c ICache, k string, v *T) bool {
	if bt := c.Get(k); bt != nil {
		if r, ok := util.GobDecodeGen[T](bt); ok {
			*v = *r
			return true
		}
		return false
	}
	return false
}

func (c *Cache) onRemove(k string, v []byte, r bigcache.RemoveReason) {
	go func() {
		defer graceful.Recover()
		for _, ch := range c.notifyEvicted {
			ch <- Entry{k, v, r}
		}
	}()
}

func (c *Cache) NotifyEvicted() <-chan Entry {
	ch := make(chan Entry, 5)
	c.notifyEvicted = append(c.notifyEvicted, ch)
	return ch
}

func (c *Cache) Get(k string) []byte {
	if v, e := c.cache.Get(k); e == nil {
		return v
	}
	return nil
}

func (c *Cache) HasGet(k string) ([]byte, bool) {
	r := c.Get(k)
	return r, r != nil
}

func (c *Cache) Has(k string) bool {
	_, ok := c.HasGet(k)
	return ok
}

func (c *Cache) Set(k string, v []byte) bool {
	return c.cache.Set(k, v) != nil
}

func (c *Cache) Delete(k string) {
	c.cache.Delete(k)
}

func (c *Cache) Close() error {
	for _, ch := range c.notifyEvicted {
		close(ch)
	}
	return c.cache.Close()
}

func (c *Cache) SetGob(k string, v interface{}) bool {
	bt := util.GobEncode(v)
	if bt == nil {
		return false
	}
	if e := c.cache.Set(k, bt); e != nil {
		log.Println(e)
		return false
	}
	return true
}

func (c *Cache) Refresh(k string) {
	if bt, ok := c.HasGet(k); ok {
		c.Delete(k)
		c.Set(k, bt)
	}
}