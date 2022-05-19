package objectstream

import (
	"encoding/base64"
	"fmt"
	"goodfs/apiserver/global"
	"goodfs/lib/util"
	"log"
)

type resumeToken struct {
	Name    string   `json:"name"`
	Hash    string   `json:"hash"`
	Size    int64    `json:"size"`
	Servers []string `json:"servers"`
	Ids     []string `json:"ids"`
}

//RSResumablePutStream 断点续传
type RSResumablePutStream struct {
	*RSPutStream
	*resumeToken
}

//NewRSResumablePutStreamFromToken TODO 解密
func NewRSResumablePutStreamFromToken(token string) (*RSResumablePutStream, error) {
	bt, e := base64.StdEncoding.DecodeString(token)
	if e != nil {
		return nil, e
	}
	var tk resumeToken
	if ok := util.GobDecodeGen2(bt, &tk); ok {
		return continueRSResumablePut(tk)
	}
	return nil, fmt.Errorf("invalid token")
}

func continueRSResumablePut(token resumeToken) (*RSResumablePutStream, error) {
	putStream, e := newRSPutStreamWithoutPost(token.Servers, token.Ids, token.Hash)
	if e != nil {
		return nil, e
	}
	return &RSResumablePutStream{putStream, &token}, nil
}

func NewRSResumablePutStream(ips []string, name, hash string, size int64) (*RSResumablePutStream, error) {
	putStream, e := NewRSPutStream(ips, hash, size)
	if e != nil {
		return nil, e
	}
	ids := make([]string, global.Config.Rs.AllShards())
	for i := range ids {
		ids[i] = putStream.writers[i].(*PutStream).tmpId
	}
	token := &resumeToken{
		Name:    name,
		Hash:    hash,
		Servers: ips,
		Size:    size,
		Ids:     ids,
	}
	return &RSResumablePutStream{putStream, token}, nil
}

//CurrentSize IO: 请求数据服务器获取分片大小
func (p *RSResumablePutStream) CurrentSize() int64 {
	size, e := HeadTmpObject(p.Servers[0], p.Ids[0])
	if e != nil {
		log.Println(e)
		return -1
	}
	size *= int64(global.Config.Rs.DataShards)
	if size > p.Size {
		return p.Size
	}
	return size
}

//Token TODO 加密
func (p *RSResumablePutStream) Token() string {
	tk := resumeToken{
		Name:    p.Name,
		Hash:    p.Hash,
		Size:    p.Size,
		Servers: p.Servers,
		Ids:     p.Ids,
	}
	return base64.StdEncoding.EncodeToString(util.GobEncode(tk))
}
