package usecase

import (
	"common/response"
	"errors"
)

var (
	ErrNotFound       = response.NewError(404, "not found")
	ErrUnknownRaftLog = response.NewError(500, "unknown raft log type")
	ErrReadOnly       = response.NewError(502, "server is readonly")
	ErrOldData        = response.NewError(400, "data expired")
	ErrExists         = response.NewError(400, "data exists")
	ErrNilData        = response.NewError(400, "null value")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}
