package entity

//go:generate msgp -tests=false

type LogType int8
type Dest int8

const (
	LogInsert LogType = 1 << iota
	LogRemove
	LogUpdate
	LogMigrate
)

const (
	DestVersion Dest = 1 << iota
	DestVersionAll
	DestMetadata
	DestBucket
)

type RaftData struct {
	Type     LogType   `msg:"type" json:"type"`
	Dest     Dest      `msg:"dest" json:"dest"`
	Name     string    `msg:"name" json:"name"`
	Sequence uint64    `msg:"sequence" json:"sequence,omitempty"`
	Version  *Version  `msg:"version" json:"version,omitempty"`
	Metadata *Metadata `msg:"metadata" json:"metadata,omitempty"`
	Bucket   *Bucket   `msg:"bucket" json:"bucket,omitempty"`
	Batch    bool      `msg:"-" json:"-"`
}
