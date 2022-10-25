package usecase

import (
	"common/pb"
	"common/response"
	"io"
	"metaserver/internal/entity"
	"time"

	"github.com/hashicorp/raft"
	bolt "go.etcd.io/bbolt"
)

type (
	MetadataRpcService interface {
		ForeachVersionBytes(string, func([]byte) bool)
		GetMetadataBytes(string) ([]byte, error)
		FilterKeys(fn func(string) bool) []string
		FindByHash(hash string) (res []*pb.Version, err error)
		UpdateLocates(name string, seq int, locates []string) error
	}

	IMetadataService interface {
		MetadataRpcService
		AddMetadata(*entity.Metadata) error
		AddVersion(string, *entity.Version) (int, error)
		UpdateMetadata(string, *entity.Metadata) error
		UpdateVersion(string, int, *entity.Version) error
		RemoveMetadata(string) error
		RemoveVersion(string, int) error
		GetMetadata(string, int) (*entity.Metadata, *entity.Version, error)
		GetVersion(string, int) (*entity.Version, error)
		ListVersions(string, int, int) ([]*entity.Version, error)
		ListMetadata(prefix string, size int) (lst []*entity.Metadata, err error)
	}

	WritableRepo interface {
		AddMetadata(*entity.Metadata) error
		AddVersion(string, *entity.Version) error
		UpdateMetadata(string, *entity.Metadata) error
		UpdateVersion(string, *entity.Version) error
		RemoveMetadata(string) error
		RemoveVersion(string, uint64) error
	}

	ReadableRepo interface {
		GetMetadata(string) (*entity.Metadata, error)
		GetVersion(string, uint64) (*entity.Version, error)
		ListVersions(string, int, int) ([]*entity.Version, error)
		ListMetadata(prefix string, size int) (lst []*entity.Metadata, err error)
	}

	IHashIndexRepo interface {
		Remove(hash, key string) error
		FindAll(hash string) ([]string, error)
		Sync() error
	}

	IBatchMetaRepo interface {
		WritableRepo
		ForeachKeys(func(string) bool)
		Sync() error
	}

	IMetadataRepo interface {
		WritableRepo
		ReadableRepo
		RemoveAllVersion(string) error
		ApplyRaft(*entity.RaftData) (bool, *response.RaftFsmResp)
		GetLastVersionNumber(name string) uint64
		ReadDB() (io.ReadCloser, error)
		ReplaceDB(io.Reader) error
		ForeachVersionBytes(string, func([]byte) bool)
		GetMetadataBytes(string) ([]byte, error)
	}

	TxFunc func(*bolt.Tx) error

	ITransaction interface {
		Update(func(*bolt.Tx) error) error
		Batch(func(*bolt.Tx) error) error
		View(func(*bolt.Tx) error) error
	}

	IRaft interface {
		Apply(cmd []byte, timeout time.Duration) raft.ApplyFuture
		ApplyLog(log raft.Log, timeout time.Duration) raft.ApplyFuture
	}

	IRaftLeaderChanged interface {
		OnLeaderChanged(bool)
	}

	IHashSlotService interface {
		AutoMigrate(toLoc *pb.LocationInfo, slots []string) error
		PrepareMigrationFrom(loc *pb.LocationInfo, slots []string) error
		PrepareMigrationTo(loc *pb.LocationInfo, slots []string) error
		ReceiveItem(*pb.MigrationItem) error
		FinishReceiveItem(bool) error
		GetCurrentSlots(bool) (map[string][]string, error)
	}

	IMetaCache interface {
		ReadableRepo
		WritableRepo
	}
)