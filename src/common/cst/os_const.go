package cst

import (
	"io/fs"
	"os"
)

var OS = struct {
	ModeUser   fs.FileMode
	WriteFlag  int
	PageSize   int
	NetPkgSize int
}{
	ModeUser:   0700,
	WriteFlag:  os.O_WRONLY | os.O_CREATE | os.O_TRUNC,
	PageSize:   4096,
	NetPkgSize: 64 << 10,
}
