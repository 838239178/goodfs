package resource

import (
	"common/logs"
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-contrib/static"
)

//go:embed web
var embedFs embed.FS

type embedFileSystem struct {
	http.FileSystem
	indexes bool
}

func (e embedFileSystem) Open(name string) (http.File, error) {
	file, err := e.FileSystem.Open(name)
	if strings.HasSuffix(name, "config.js") {
		file = newConfigJS(file)
	}
	return file, err
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	if strings.HasPrefix(path, "/api") {
		return false
	}
	if path == "/" {
		path = "/index.html"
	}
	f, err := e.Open(path)
	if err != nil {
		return false
	}

	// check if indexing is allowed
	s, err := f.Stat()
	if err != nil {
		logs.Std().Error(err)
		return false
	}
	if s.IsDir() && !e.indexes {
		return false
	}

	return true
}

func FileSystem() static.ServeFileSystem {
	fsys, err := fs.Sub(embedFs, "web")
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
		indexes:    false,
	}
}
