package repo

import (
	"common/cache"
	"common/proto/msg"
	"common/util"
	"fmt"
	"metaserver/internal/usecase"
	"metaserver/internal/usecase/logic"
)

type PartlyMatchedErr interface {
	error
	Last() int
}

type partlyErr struct {
	lastIndex int
}

func newPartlyErr(last int) *partlyErr {
	return &partlyErr{lastIndex: last}
}

func (pe *partlyErr) Last() int {
	return pe.lastIndex
}

func (pe *partlyErr) Error() string {
	return "partly matched"
}

const (
	MetaCachePrefix = "metadata_"
)

type MetadataCacheRepo struct {
	cache cache.ICache
}

func NewMetadataCacheRepo(c cache.ICache) *MetadataCacheRepo {
	return &MetadataCacheRepo{c}
}

func (m *MetadataCacheRepo) ListMetadata(string, int) ([]*msg.Metadata, int, error) {
	panic("not impl ListMetadata")
}

func (m *MetadataCacheRepo) GetMetadata(s string) (*msg.Metadata, error) {
	if bt, ok := m.cache.HasGet(fmt.Sprint(MetaCachePrefix, s)); ok {
		var en msg.Metadata
		if err := util.DecodeMsgp(&en, bt); err != nil {
			return nil, err
		}
		return &en, nil
	}
	return nil, usecase.ErrNotFound
}

func (m *MetadataCacheRepo) GetVersion(s string, u uint64) (*msg.Version, error) {
	key := fmt.Sprint(MetaCachePrefix, s, logic.Sep, u)
	if bt, ok := m.cache.HasGet(key); ok {
		var en msg.Version
		if err := util.DecodeMsgp(&en, bt); err != nil {
			return nil, err
		}
		return &en, nil
	}
	return nil, usecase.ErrNotFound
}

// ListVersions return successfully matched cache until failure.
// returning PartlyMatchedErr if not fully matched all
func (m *MetadataCacheRepo) ListVersions(s string, start int, end int) ([]*msg.Version, int, error) {
	size := end - start + 1
	res := make([]*msg.Version, 0, size)
	for i := start; i <= end; i++ {
		v, err := m.GetVersion(s, uint64(i))
		if err != nil {
			return res, 0, newPartlyErr(i)
		}
		res = append(res, v)
	}
	return res, 0, nil
}

func (m *MetadataCacheRepo) AddMetadata(id string, metadata *msg.Metadata) error {
	bt, err := util.EncodeMsgp(metadata)
	if err != nil {
		return err
	}
	m.cache.Set(fmt.Sprint(MetaCachePrefix, id), bt)
	return nil
}

func (m *MetadataCacheRepo) AddVersion(s string, version *msg.Version) error {
	key := fmt.Sprint(MetaCachePrefix, s, logic.Sep, version.Sequence)
	bt, err := util.EncodeMsgp(version)
	if err != nil {
		return err
	}
	m.cache.Set(key, bt)
	return nil
}

func (m *MetadataCacheRepo) UpdateMetadata(id string, metadata *msg.Metadata) error {
	return m.AddMetadata(id, metadata)
}

func (m *MetadataCacheRepo) UpdateVersion(s string, version *msg.Version) error {
	return m.AddVersion(s, version)
}

func (m *MetadataCacheRepo) RemoveMetadata(s string) error {
	m.cache.Delete(fmt.Sprint(MetaCachePrefix, s))
	return nil
}

func (m *MetadataCacheRepo) RemoveVersion(s string, u uint64) error {
	m.cache.Delete(fmt.Sprint(MetaCachePrefix, s, logic.Sep, u))
	return nil
}

func (m *MetadataCacheRepo) RemoveAllVersion(string) error {
	panic("not impl RemoveAllVersion")
}

func (m *MetadataCacheRepo) AddVersionFromRaft(s string, version *msg.Version) error {
	return m.AddVersion(s, version)
}
