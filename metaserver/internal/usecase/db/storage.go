package db

import (
	"common/util"
	"time"

	bolt "go.etcd.io/bbolt"
)

type Storage struct {
	*bolt.DB
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Stop() error {
	// FIXME close directly may cause panic
	return s.DB.Close()
}

func (s *Storage) Open(path string) (err error) {
	// FIXME close directly may cause panic
	s.DB, err = bolt.Open(path, util.OS_ModeUser, &bolt.Options{
		Timeout:    12 * time.Second,
		NoGrowSync: false,
	})
	return
}
