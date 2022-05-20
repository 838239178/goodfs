package objectstream

import (
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
)

//PutStream 需要确保调用了Close或者Commit
//Close() Commit() 可重复调用
type PutStream struct {
	Locate    string
	name      string
	tmpId     string
	writer    *io.PipeWriter
	errorChan chan error
	closed    atomic.Value
}

type GetStream struct {
	io.ReadCloser
	Locate string
}

//NewPutStream IO: 发送Post请求到数据服务器
func NewPutStream(ip, name string, size int64) (*PutStream, error) {
	c := make(chan error, 1)
	id, e := PostTmpObject(ip, name, size)
	if e != nil {
		return nil, e
	}
	flag := atomic.Value{}
	flag.Store(false)
	res := &PutStream{errorChan: c, Locate: ip, name: name, tmpId: id, closed: flag}
	return res, nil
}

//newExistedPutStream 不发送Post请求
func newExistedPutStream(ip, name, id string) *PutStream {
	c := make(chan error, 1)
	flag := atomic.Value{}
	flag.Store(false)
	res := &PutStream{errorChan: c, Locate: ip, name: name, tmpId: id, closed: flag}
	return res
}

//NewGetStream IO: Get object
func NewGetStream(ip, name string) (*GetStream, error) {
	resp, err := GetObject(ip, name)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dataServer return http code %v", resp.StatusCode)
	}
	return &GetStream{resp.Body, ip}, nil
}

func (p *PutStream) Close() error {
	if p.closed.CompareAndSwap(false, true) {
		defer close(p.errorChan)
		//如果没有发生写入
		if p.writer != nil {
			if err := p.writer.Close(); err != nil {
				return err
			}
		} else {
			p.errorChan <- nil
		}
		return <-p.errorChan
	}
	return fmt.Errorf("already closed")
}

func (p *PutStream) Write(b []byte) (n int, err error) {
	if p.writer == nil {
		p.startWrite()
	}
	return p.writer.Write(b)
}

func (p *PutStream) startWrite() {
	reader, writer := io.Pipe()
	p.writer = writer
	go func() {
		p.errorChan <- PatchTmpObject(p.Locate, p.tmpId, reader)
	}()
}

//Commit IO: send commit message and close stream
func (p *PutStream) Commit(ok bool) error {
	if e := p.Close(); e != nil {
		return e
	}

	if !ok {
		go DeleteTmpObject(p.Locate, p.tmpId)
		return nil
	}

	return PutTmpObject(p.Locate, p.tmpId, p.name)
}
