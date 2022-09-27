package repo

import (
	"metaserver/internal/entity"
	"metaserver/internal/usecase"
	"metaserver/internal/usecase/db"
	"metaserver/internal/usecase/logic"
)

type BatchMetaRepo struct {
	Storage *db.Storage
}

func NewBatchRepo(stroe *db.Storage) *BatchMetaRepo {
	return &BatchMetaRepo{Storage: stroe}
}

func (br *BatchMetaRepo) Sync() error {
	return br.Storage.DB().Sync()
}

func (br *BatchMetaRepo) ForeachKeys(fn func(string) bool) {
	_ = br.Storage.View(logic.ForeachKeys(fn))
}

func (br *BatchMetaRepo) AddVersion(name string, data *entity.Version) error {
	if data == nil {
		return usecase.ErrNilData
	}
	return br.Storage.DB().Batch(logic.AddVer(name, data))
}

func (br *BatchMetaRepo) UpdateVersion(name string, data *entity.Version) error {
	if data == nil {
		return usecase.ErrNilData
	}
	return br.Storage.DB().Batch(logic.UpdateVer(name, data))
}

func (br *BatchMetaRepo) RemoveVersion(name string, ver uint64) error {
	return br.Storage.DB().Batch(logic.RemoveVer(name, ver))
}

func (br *BatchMetaRepo) AddMetadata(data *entity.Metadata) error {
	if data == nil {
		return usecase.ErrNilData
	}
	return br.Storage.DB().Batch(logic.AddMeta(data))
}

func (br *BatchMetaRepo) UpdateMetadata(name string, data *entity.Metadata) error {
	return br.Storage.DB().Batch(logic.UpdateMeta(name, data))
}

func (br *BatchMetaRepo) RemoveMetadata(name string) error {
	return br.Storage.DB().Batch(logic.RemoveMeta(name))
}