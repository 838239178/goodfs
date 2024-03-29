package repo

import (
	"bytes"
	"common/cst"
	"common/graceful"
	"common/logs"
	"common/proto/msg"
	"common/util"
	"errors"
	"fmt"
	"io"
	"metaserver/internal/usecase"
	"metaserver/internal/usecase/db"
	"metaserver/internal/usecase/logic"
	"os"
	"strings"
	"time"

	bolt "go.etcd.io/bbolt"
)

type MetadataRepo struct {
	MainDB *db.Storage
	Cache  usecase.IMetaCache
}

func NewMetadataRepo(db *db.Storage, c usecase.IMetaCache) *MetadataRepo {
	return &MetadataRepo{MainDB: db, Cache: c}
}

func (m *MetadataRepo) AddMetadata(id string, data *msg.Metadata) error {
	if data == nil {
		return usecase.ErrNilData
	}
	if err := m.MainDB.Update(logic.AddMeta(id, data)); err != nil {
		return err
	}
	go func() {
		defer graceful.Recover()
		err := m.Cache.AddMetadata(id, data)
		util.LogErrWithPre("metadata cache", err)
	}()
	return nil
}

func (m *MetadataRepo) UpdateMetadata(name string, data *msg.Metadata) error {
	if data == nil {
		return usecase.ErrNilData
	}
	if err := m.MainDB.Update(logic.UpdateMeta(name, data)); err != nil {
		return err
	}
	go func() {
		defer graceful.Recover()
		err := m.Cache.UpdateMetadata(name, data)
		util.LogErrWithPre("metadata cache", err)
	}()
	return nil
}

func (m *MetadataRepo) RemoveMetadata(name string) error {
	lastVer := m.GetLastVersionNumber(name)
	if err := m.MainDB.Update(logic.RemoveMeta(name)); err != nil {
		return err
	}
	go func() {
		defer graceful.Recover()
		util.LogErrWithPre("metadata cache", m.Cache.RemoveMetadata(name))
		for i := uint64(1); i <= lastVer; i++ {
			util.LogErrWithPre("metadata cache", m.Cache.RemoveVersion(name, i))
		}
	}()
	return nil
}

func (m *MetadataRepo) GetMetadata(id string) (*msg.Metadata, error) {
	if data, err := m.Cache.GetMetadata(id); err == nil {
		return data, nil
	}
	data := &msg.Metadata{}
	if err := m.MainDB.View(logic.GetMeta(id, data)); err != nil {
		return nil, err
	}
	go func() {
		defer graceful.Recover()
		util.LogErrWithPre("add metadata cache", m.Cache.AddMetadata(id, data))
	}()
	return data, nil
}

func (m *MetadataRepo) AddVersion(id string, data *msg.Version) error {
	if data == nil {
		return usecase.ErrNilData
	}
	if data.Hash == "" {
		return errors.New("version doesn't contains Hash value")
	}
	if data.UniqueId == "" {
		return errors.New("version doesn't contains UniqueId value")
	}
	if err := m.MainDB.Update(logic.AddVer(id, data)); err != nil {
		return err
	}
	go func() {
		defer graceful.Recover()
		err := m.Cache.AddVersion(id, data)
		util.LogErrWithPre("metadata cache", err)
	}()
	return nil
}

func (m *MetadataRepo) AddVersionFromRaft(id string, data *msg.Version) error {
	if data == nil {
		return usecase.ErrNilData
	}
	if data.Hash == "" {
		return errors.New("version doesn't contains Hash value")
	}
	if data.UniqueId == "" {
		return errors.New("version doesn't contains UniqueId value")
	}
	if err := m.MainDB.Update(logic.AddVerWithSequence(id, data)); err != nil {
		return err
	}
	go func() {
		defer graceful.Recover()
		err := m.Cache.AddVersionFromRaft(id, data)
		util.LogErrWithPre("metadata cache", err)
	}()
	return nil
}

func (m *MetadataRepo) UpdateVersion(id string, data *msg.Version) error {
	if data == nil {
		return usecase.ErrNilData
	}
	if err := m.MainDB.Update(logic.UpdateVer(id, data)); err != nil {
		return err
	}
	go func() {
		defer graceful.Recover()
		err := m.Cache.UpdateVersion(id, data)
		util.LogErrWithPre("metadata cache", err)
	}()
	return nil
}

func (m *MetadataRepo) RemoveVersion(name string, ver uint64) error {
	if err := m.MainDB.Update(logic.RemoveVer(name, ver)); err != nil {
		return err
	}
	go func() {
		defer graceful.Recover()
		err := m.Cache.RemoveVersion(name, ver)
		util.LogErrWithPre("metadata cache", err)
	}()
	return nil
}

func (m *MetadataRepo) RemoveAllVersion(name string) error {
	last := m.GetLastVersionNumber(name)
	if err := m.MainDB.Update(func(tx *bolt.Tx) error {
		// delete bucket
		if err := logic.RemoveVersionBucket(tx, name); err != nil {
			return err
		}
		// create an empty bucket
		return logic.CreateVersionBucket(tx, name)
	}); err != nil {
		return err
	}
	go func() {
		defer graceful.Recover()
		for i := uint64(0); i <= last; i++ {
			err := m.Cache.RemoveVersion(name, i)
			util.LogErrWithPre("remove version cache", err)
		}
	}()
	return nil
}

func (m *MetadataRepo) GetFirstVersionNumber(name string) uint64 {
	var fst uint64 = 1
	if err := m.MainDB.View(func(tx *bolt.Tx) error {
		if buk := logic.GetVersionBucket(tx, name); buk != nil {
			k, v := buk.Cursor().First()
			if k == nil || v == nil {
				return usecase.ErrNotFound
			}
			idx := bytes.LastIndexByte(k, logic.Sep[0])
			if idx < 0 {
				return usecase.ErrNotFound
			}
			fst = util.ToUint64(util.BytesToStr(k[idx+1:]))
		}
		return nil
	}); err != nil {
		return 0
	}
	return fst
}

func (m *MetadataRepo) GetLastVersionNumber(name string) uint64 {
	var max uint64 = 1
	if err := m.MainDB.View(func(tx *bolt.Tx) error {
		if buk := logic.GetVersionBucket(tx, name); buk != nil {
			k, v := buk.Cursor().Last()
			if k == nil || v == nil {
				return usecase.ErrNotFound
			}
			idx := bytes.LastIndexByte(k, logic.Sep[0])
			if idx < 0 {
				return usecase.ErrNotFound
			}
			max = util.ToUint64(util.BytesToStr(k[idx+1:]))
		}
		return nil
	}); err != nil {
		return 0
	}
	return max
}

func (m *MetadataRepo) GetVersion(name string, ver uint64) (*msg.Version, error) {
	if data, err := m.Cache.GetVersion(name, ver); err == nil {
		return data, nil
	}
	data := &msg.Version{}
	if err := m.MainDB.View(logic.GetVer(name, ver, data)); err != nil {
		return nil, err
	}
	go func() {
		defer graceful.Recover()
		util.LogErrWithPre("add metadata cache", m.Cache.AddVersion(name, data))
	}()
	return data, nil
}

func (m *MetadataRepo) ListVersions(name string, start int, end int) (lst []*msg.Version, total int, err error) {
	size := end - start + 1
	lst, _, err = m.Cache.ListVersions(name, start, end)
	if util.InstanceOf[PartlyMatchedErr](err) {
		start = err.(PartlyMatchedErr).Last()
	} else if err != nil {
		return
	}
	err = m.MainDB.View(func(tx *bolt.Tx) error {
		buk := logic.GetVersionBucket(tx, name)
		if buk == nil {
			return usecase.ErrNotFound
		}
		c := buk.Cursor()

		min := util.StrToBytes(fmt.Sprint(name, logic.Sep, start))

		for k, v := c.Seek(min); k != nil && len(lst) < size; k, v = c.Next() {
			data := &msg.Version{}
			if err := util.DecodeMsgp(data, v); err != nil {
				return err
			}
			lst = append(lst, data)
		}
		// record total
		total = buk.Stats().KeyN
		return nil
	})
	return
}

func (m *MetadataRepo) ListMetadata(prefix string, size int) (lst []*msg.Metadata, total int, err error) {
	err = m.MainDB.View(func(tx *bolt.Tx) error {
		root := logic.GetMetadataBucket(tx)
		if root == nil {
			return usecase.ErrNotFound
		}
		cur := root.Cursor()
		var k, v []byte
		if prefix != "" {
			k, v = cur.Seek(util.StrToBytes(prefix))
			defer func() { total = len(lst) }()
		} else {
			k, v = cur.First()
			total = root.Stats().KeyN
		}
		for k != nil && len(lst) < size {
			if prefix != "" && !strings.HasPrefix(util.BytesToStr(k), prefix) {
				break
			}
			if len(v) > 0 {
				var data msg.Metadata
				if err := util.DecodeMsgp(&data, v); err != nil {
					return err
				}
				lst = append(lst, &data)
			}
			k, v = cur.Next()
		}
		return nil
	})
	return
}

func (m *MetadataRepo) Snapshot() (usecase.SnapshotTx, error) {
	return m.MainDB.DB().Begin(false)
}

func (m *MetadataRepo) Restore(r io.Reader) (err error) {
	dbPath := m.MainDB.DB().Path() + "_replace"
	// open new db file
	newFile, err := os.OpenFile(dbPath, os.O_WRONLY|os.O_CREATE, cst.OS.ModeUser)
	if err != nil {
		logs.Std().Errorf("restore fail on open new file: %v", err)
		return err
	}
	// save new db data
	n, err := io.Copy(newFile, r)
	if err != nil {
		logs.Std().Errorf("restore fail on copy data to new file: %v, written %d", err, n)
		return err
	}
	if err := newFile.Close(); err != nil {
		logs.Std().Errorf("close new db file err: %s", err)
		return err
	}
	// reopen db
	return m.MainDB.Replace(dbPath)
}

func (m *MetadataRepo) ForeachVersionBytes(name string, fn func([]byte) bool) {
	_ = m.MainDB.View(func(tx *bolt.Tx) error {
		_ = logic.GetVersionBucket(tx, name).ForEach(func(k, v []byte) error {
			if !fn(v) {
				return usecase.ErrNotFound
			}
			return nil
		})
		return nil
	})
}

func (m *MetadataRepo) GetMetadataBytes(key string) ([]byte, error) {
	var res []byte
	err := m.MainDB.View(func(tx *bolt.Tx) error {
		v := logic.GetMetadataBucket(tx).Get(util.StrToBytes(key))
		if v == nil {
			return usecase.ErrNotFound
		}
		// copy values from the mmap pointer
		res, _ = io.ReadAll(bytes.NewBuffer(v))
		return nil
	})
	return res, err
}

func (m *MetadataRepo) GetExtra(id string) (*msg.Extra, error) {
	var i msg.Extra
	err := m.MainDB.View(logic.GetExtra(id, &i))
	return &i, err
}

const lastAppliedIndexKey = "go.dfs.metadata.special.lastAppliedIndexKey"

func (m *MetadataRepo) LastAppliedIndex() (uint64, error) {
	var r uint64
	err := m.MainDB.View(func(tx *bolt.Tx) error {
		specBuc := tx.Bucket(util.StrToBytes(lastAppliedIndexKey))
		if specBuc == nil {
			return nil
		}
		r = specBuc.Sequence()
		return nil
	})
	return r, err
}

func (m *MetadataRepo) ApplyIndex(i uint64) error {
	return m.MainDB.Update(func(tx *bolt.Tx) error {
		specBuc, err := tx.CreateBucketIfNotExists(util.StrToBytes(lastAppliedIndexKey))
		if err != nil {
			return err
		}
		if i > specBuc.Sequence() {
			return specBuc.SetSequence(i)
		}
		return nil
	})
}

func (m *MetadataRepo) UpdateLocateByHash(hash string, index int, value string) error {
	return m.MainDB.Update(func(tx *bolt.Tx) error {
		var keys []string
		if err := logic.NewHashIndexLogic().GetIndex(hash, &keys)(tx); err != nil {
			return err
		}
		var v msg.Version
		for _, key := range keys {
			idx := strings.LastIndexByte(key, '.')
			metaId, sequence := key[:idx], util.ToUint64(key[idx+1:])
			if err := logic.GetVer(metaId, sequence, &v)(tx); err != nil {
				return err
			}
			v.Ts = time.Now().UnixMilli()
			v.Locate[index] = value
			if err := logic.UpdateVer(metaId, &v)(tx); err != nil {
				return err
			}
		}
		return nil
	})
}

func (m *MetadataRepo) ExistsByUniqueId(uniqueId string) (bool, error) {
	var res []string
	if err := m.MainDB.View(logic.NewUniqueIdIndex().GetIndex(uniqueId, &res)); err != nil {
		return false, err
	}
	return len(res) > 0, nil
}
