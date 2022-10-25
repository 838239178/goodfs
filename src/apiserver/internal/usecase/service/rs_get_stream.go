package service

import (
	"apiserver/internal/usecase"
	"apiserver/internal/usecase/componet/selector"
	"apiserver/internal/usecase/logic"
	global "apiserver/internal/usecase/pool"
	"bytes"
	"common/graceful"
	"common/logs"
	"common/util"
	"fmt"
	"io"
	"sync"
)

type RSGetStream struct {
	*rsDecoder
}

type provideStream struct {
	stream io.Reader
	index  int
	err    error
}

func provideGetStream(hash string, locates []string) <-chan *provideStream {
	respChan := make(chan *provideStream, 1)
	go func() {
		defer graceful.Recover()
		defer close(respChan)
		var wg sync.WaitGroup
		for i, ip := range locates {
			wg.Add(1)
			go func(idx int, ip string) {
				defer graceful.Recover()
				defer wg.Done()
				if len(ip) > 0 {
					reader, e := NewGetStream(ip, fmt.Sprintf("%s.%d", hash, idx))
					respChan <- &provideStream{reader, idx, e}
				} else {
					respChan <- &provideStream{nil, idx, fmt.Errorf("shard %s.%d lost", hash, idx)}
				}
			}(i, ip)
		}
		wg.Wait()
	}()
	return respChan
}

func NewRSGetStream(size int64, hash string, locates []string) (*RSGetStream, error) {
	readers := make([]io.Reader, global.Config.Rs.AllShards())
	writers := make([]io.Writer, global.Config.Rs.AllShards())
	dsNum := int64(global.Config.Rs.DataShards)
	perSize := (size + dsNum - 1) / dsNum
	lb := selector.IPSelector{Selector: global.Balancer, IPs: logic.NewDiscovery().GetDataServers()}
	var e error
	for r := range provideGetStream(hash, locates) {
		if r.err != nil {
			logs.Std().Error(r.err)
			ip := lb.Select()
			writers[r.index], r.err = NewPutStream(ip, fmt.Sprintf("%s.%d", hash, r.index), perSize)
			if r.err != nil {
				return nil, r.err
			}
			//需更新元数据
			locates[r.index] = ip
			e = usecase.ErrNeedUpdateMeta
		} else {
			readers[r.index] = r.stream
		}
	}
	dec := NewDecoder(readers, writers, size)
	return &RSGetStream{dec}, e
}

func (g *RSGetStream) Seek(offset int64, whence int) (int64, error) {
	if whence != io.SeekCurrent {
		panic("only support io.SeekCurrent")
	}
	if offset < 0 {
		return 0, fmt.Errorf("only support forward seek offest")
	}

	//读取offset长度的数据，丢弃于内存
	length := int64(global.Config.Rs.BlockSize())
	buf := bytes.NewBuffer(make([]byte, length))
	for offset > 0 {
		if length > offset {
			//当剩余未读取内容少于BlockSize时 减少读取量
			length = offset
		}
		if _, e := io.CopyN(buf, g, length); e != nil {
			return offset, e
		}
		offset -= length
	}
	buf = nil
	return offset, nil
}

func (g *RSGetStream) Close() error {
	wg := util.NewDoneGroup()
	defer wg.Close()
	for _, w := range g.writers {
		if util.InstanceOf[Commiter](w) {
			wg.Todo()
			go func(cm Commiter) {
				defer graceful.Recover()
				defer wg.Done()
				if e := cm.Commit(true); e != nil {
					wg.Error(e)
				}
			}(w.(Commiter))
		}
	}
	return wg.WaitUntilError()
}