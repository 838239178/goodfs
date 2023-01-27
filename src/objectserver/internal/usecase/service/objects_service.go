package service

import (
	"bytes"
	"common/cst"
	"common/datasize"
	"common/graceful"
	"common/logs"
	"common/response"
	"common/system/disk"
	"common/util"
	"fmt"
	"io"
	global "objectserver/internal/usecase/pool"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/s2"
)

const (
	LocateKeyPrefix = "LocateCache#"
)

// Exist check if the object exists. pass to ExistPath
func Exist(name string) bool {
	if global.Cache.Has(LocateKeyPrefix+name) || global.Cache.Has(name) {
		return true
	}
	realPath, ok := FindRealStoragePath(name)
	if !ok {
		return false
	}
	if ExistPath(realPath) {
		MarkExist(name)
		return true
	} else {
		return false
	}
}

// FindRealStoragePath find storage path with real mount point of this file
func FindRealStoragePath(fileName string) (string, bool) {
	path, err := global.PathDB.GetLast(fileName)
	if err != nil {
		path, err = global.DriverManager.FindMountPath(filepath.Join(global.Config.StoragePath, fileName))
		if err != nil {
			return path, false
		}
	}
	return path, true
}

// ExistPath check if the given path exists.
func ExistPath(fullPath string) bool {
	_, err := os.Stat(fullPath)
	return !os.IsNotExist(err)
}

// MarkExist save object mark into cache
func MarkExist(name string) {
	global.Cache.Set(LocateKeyPrefix+name, []byte{})
}

// UnMarkExist remove object mark from cache
func UnMarkExist(name string) {
	global.Cache.Delete(LocateKeyPrefix + name)
}

// Put save object to storage path
func Put(fileName string, fileStream io.Reader, compress bool) (err error) {
	if Exist(fileName) {
		return
	}

	mp := global.Config.BaseMountPoint
	if dm, err := global.DriverManager.SelectDriver(); err == nil {
		mp = dm.MountPoint
	}
	fullPath := filepath.Join(mp, global.Config.StoragePath, fileName)

	var size int64
	if compress {
		size, err = WriteFileCompress(fullPath, fileStream)
	} else {
		size, err = WriteFile(fullPath, fileStream)
	}
	if err != nil {
		return
	}
	go func() {
		defer graceful.Recover()
		global.ObjectCap.CurrentCap.Add(uint64(size))
		util.LogErr(global.PathDB.Put(fileName, mp))
		MarkExist(fileName)
	}()
	return
}

// Get read object to writer with provided size. pass to GetFile
func Get(name string, offset, size int64, compress bool, writer io.Writer) (err error) {
	fullPath, ok := FindRealStoragePath(name)
	if !ok {
		return response.NewError(404, "object not found")
	}
	if compress {
		err = GetFileCompress(fullPath, offset, size, writer)
	} else {
		err = GetFile(fullPath, offset, size, writer)
	}
	if err != nil {
		return err
	}
	MarkExist(name)
	return nil
}

// GetFile read file with provided size. offset corresponds to io.SeekStart.
// if offset is not multiple of 4KB, direct-io will be disabled.
func GetFile(fullPath string, offset, size int64, writer io.Writer) error {
	file, err := disk.OpenFileDirectIO(fullPath, os.O_RDONLY, cst.OS.ModeUser)
	if os.IsNotExist(err) {
		return response.NewError(404, "object not found")
	}
	if err != nil {
		return err
	}
	defer file.Close()
	if offset > 0 {
		if int(offset)%cst.OS.PageSize > 0 {
			logs.Std().Warn("offset must be power of 4KB, direct-io will be disabled")
			if err = disk.DisableDirectIO(file); err != nil {
				return fmt.Errorf("diable direct-io: %w", err)
			}
		}
		if _, err = file.Seek(offset, io.SeekStart); err != nil {
			return err
		}
	}
	_, err = io.CopyBuffer(writer, disk.LimitReader(file, size), disk.AlignedBlock(8*cst.OS.PageSize))
	return err
}

// Delete remove the object under the storage path
func Delete(name string) error {
	fullPath, ok := FindRealStoragePath(name)
	if !ok {
		return response.NewError(404, "object not found")
	}
	size, err := DeleteFile(fullPath, name)
	if err != nil {
		return err
	}
	go func() {
		defer graceful.Recover()
		global.ObjectCap.CurrentCap.Sub(uint64(size))
		util.LogErr(global.PathDB.Remove(name, fullPath))
		UnMarkExist(name)
	}()
	return nil
}

// DeleteFile will remove the file under the path. if file not exist, it will do nothing.
func DeleteFile(path, name string) (int64, error) {
	pt := filepath.Join(path, name)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	if err = os.Remove(pt); err != nil && !os.IsNotExist(err) {
		return 0, err
	}
	return info.Size(), nil
}

// WriteFileWithSize will append data to file using provided curSize to remove padding of data
// and work with DIO keeping read data aligned to multiple of 4KB
func WriteFileWithSize(fullPath string, curSize int64, fileStream io.Reader) (int64, error) {
	file, err := disk.OpenFileDirectIO(fullPath, os.O_RDWR|os.O_CREATE, cst.OS.ModeUser)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		return 0, err
	}
	// paddingLen always gte 0 and lt 4096
	pageSize := int64(cst.OS.PageSize)
	paddingLen := fi.Size() - curSize
	if paddingLen >= pageSize {
		return 0, fmt.Errorf("err padding length %d", paddingLen)
	}
	if paddingLen > 0 {
		// read the last 4KB of data
		if _, err = file.Seek(-pageSize, io.SeekEnd); err != nil {
			return 0, err
		}
		bt := make([]byte, pageSize)
		_, err = io.ReadFull(file, bt)
		if err != nil {
			return 0, fmt.Errorf("read tail 4KB err: %w", err)
		}
		// remove padding
		bt = bt[:len(bt)-int(paddingLen)]
		// concatenation with fileStream
		fileStream = disk.MultiReader(bytes.NewBuffer(bt), fileStream)
		// seek back
		if _, err = file.Seek(-pageSize, io.SeekEnd); err != nil {
			return 0, err
		}
	}
	// write file and aligned to power of 4KB
	return io.CopyBuffer(file, disk.NewAlignedReader(fileStream), disk.AlignedBlock(8*cst.OS.PageSize))
}

// WriteFile will append data to file and work with DIO. make sure size of each write is a multiple of 4096 (except last)
func WriteFile(fullPath string, fileStream io.Reader) (int64, error) {
	file, err := disk.OpenFileDirectIO(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, cst.OS.ModeUser)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	// write file and aligned to power of 4KB
	return io.CopyBuffer(file, disk.NewAlignedReader(fileStream), disk.AlignedBlock(8*cst.OS.PageSize))
}

// GetFileCompress read s2-compressed file and work with COW
func GetFileCompress(fullPath string, offset, size int64, writer io.Writer) error {
	file, err := os.Open(fullPath)
	if os.IsNotExist(err) {
		return response.NewError(404, "object not found")
	}
	if err != nil {
		return err
	}
	defer file.Close()
	if offset > 0 {
		if _, err = file.Seek(offset, io.SeekStart); err != nil {
			return err
		}
	}
	_, err = io.CopyBuffer(writer, disk.LimitReader(s2.NewReader(file), size), disk.AlignedBlock(8*cst.OS.PageSize))
	return err
}

// WriteFileCompress append data to file compressing by s2 and work with COW
func WriteFileCompress(fullPath string, fileStream io.Reader) (int64, error) {
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, cst.OS.ModeUser)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	wt := s2.NewWriter(file, s2.WriterBetterCompression())
	n, err := io.CopyBuffer(wt, fileStream, make([]byte, 8*cst.OS.PageSize))
	if err != nil {
		return n, err
	}
	return n, wt.Close()
}

// CommitFile move the temp file to storage path with a new name
func CommitFile(mountPoint, tmpName, fileName string, compress bool) error {
	filePath := filepath.Join(mountPoint, global.Config.StoragePath, fileName)
	tempPath := filepath.Join(mountPoint, global.Config.TempPath, tmpName)
	if ExistPath(filePath) {
		return nil
	}
	if compress {
		tmp, err := os.Open(tempPath)
		if err != nil {
			if os.IsNotExist(err) {
				return response.NewError(404, "object not found")
			}
			return err
		}
		defer tmp.Close()
		target, err := os.OpenFile(filePath, cst.OS.WriteFlag, cst.OS.ModeUser)
		if err != nil {
			return err
		}
		defer target.Close()
		wt := s2.NewWriter(target, s2.WriterBlockSize(int(datasize.MB*2)))
		_, err = io.CopyBuffer(wt, tmp, make([]byte, datasize.MB*4))
		if err != nil {
			return err
		}
		if err = wt.Close(); err != nil {
			return err
		}
	} else {
		if err := os.Rename(tempPath, filePath); err != nil {
			if os.IsNotExist(err) {
				return response.NewError(404, "object not found")
			}
			return err
		}
	}

	go func() {
		defer graceful.Recover()
		if info, err := os.Stat(filePath); err == nil {
			global.ObjectCap.CurrentCap.Add(uint64(info.Size()))
		}
		MarkExist(fileName)
	}()
	return nil
}
