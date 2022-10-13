package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"io"
)

// SHA256IO sha256算法对二进制流进行计算
func SHA256IO(reader io.Reader) string {
	crypto := sha256.New()
	if _, e := io.CopyBuffer(crypto, reader, make([]byte, 2048)); e == nil {
		b := crypto.Sum(make([]byte, 0, crypto.Size()))
		return base64.StdEncoding.EncodeToString(b)
	}
	return ""
}

func SHA256(bt []byte) string {
	return Hash(bt, sha256.New())
}

func MD5(bt []byte) string {
	return Hash(bt, md5.New())
}

func Hash(bt []byte, exec hash.Hash) string {
	_, _ = exec.Write(bt)
	res := exec.Sum(make([]byte, 0, exec.Size()))
	return base64.StdEncoding.EncodeToString(res)
}
