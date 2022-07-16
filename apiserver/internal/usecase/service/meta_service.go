package service

import (
	"apiserver/internal/entity"
	"apiserver/internal/usecase/repo"
)

type MetaService struct {
	repo        repo.IMetadataRepo
	versionRepo repo.IVersionRepo
}

func NewMetaService(repo repo.IMetadataRepo, versionRepo repo.IVersionRepo) *MetaService {
	return &MetaService{repo: repo, versionRepo: versionRepo}
}

func (m *MetaService) SaveMetadata(md *entity.Metadata) (int32, error) {
	ver := md.Versions[0]
	metaD := m.repo.FindByNameAndVerMode(md.Name, entity.VerModeNot)
	var verNum int32
	if metaD != nil {
		verNum = m.versionRepo.Add(nil, metaD.Name, ver)
	} else {
		verNum = 0
		var e error
		if metaD, e = m.repo.Insert(md); e != nil {
			verNum = repo.ErrVersion
		}
	}

	if verNum == repo.ErrVersion {
		return -1, ErrInternalServer
	} else {
		return verNum, nil
	}
}

func (m *MetaService) UpdateVersion(version *entity.Version) {
	m.versionRepo.Update(nil, version)
}

func (m *MetaService) GetVersion(name string, version int32) (*entity.Version, bool) {
	res := m.versionRepo.Find(name, version)
	if res == nil {
		return nil, false
	}
	return res, true
}

func (m *MetaService) GetMetadata(name string, ver int32) (*entity.Metadata, bool) {
	res := m.repo.FindByNameAndVerMode(name, entity.VerMode(ver))
	if res == nil {
		return nil, false
	}
	return res, true
}
