package entity

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Metadata) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "create_time":
			z.CreateTime, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "CreateTime")
				return
			}
		case "update_time":
			z.UpdateTime, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "UpdateTime")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Metadata) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "name"
	err = en.Append(0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "create_time"
	err = en.Append(0xab, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.CreateTime)
	if err != nil {
		err = msgp.WrapError(err, "CreateTime")
		return
	}
	// write "update_time"
	err = en.Append(0xab, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.UpdateTime)
	if err != nil {
		err = msgp.WrapError(err, "UpdateTime")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Metadata) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "name"
	o = append(o, 0x83, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "create_time"
	o = append(o, 0xab, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt64(o, z.CreateTime)
	// string "update_time"
	o = append(o, 0xab, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt64(o, z.UpdateTime)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Metadata) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "create_time":
			z.CreateTime, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CreateTime")
				return
			}
		case "update_time":
			z.UpdateTime, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "UpdateTime")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Metadata) Msgsize() (s int) {
	s = 1 + 5 + msgp.StringPrefixSize + len(z.Name) + 12 + msgp.Int64Size + 12 + msgp.Int64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Version) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "sequence":
			z.Sequence, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Sequence")
				return
			}
		case "hash":
			z.Hash, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Hash")
				return
			}
		case "size":
			z.Size, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Size")
				return
			}
		case "ts":
			z.Ts, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "Ts")
				return
			}
		case "ec_algo":
			z.EcAlgo, err = dc.ReadInt8()
			if err != nil {
				err = msgp.WrapError(err, "EcAlgo")
				return
			}
		case "data_shards":
			z.DataShards, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "DataShards")
				return
			}
		case "parity_shards":
			z.ParityShards, err = dc.ReadInt32()
			if err != nil {
				err = msgp.WrapError(err, "ParityShards")
				return
			}
		case "shard_size":
			z.ShardSize, err = dc.ReadInt64()
			if err != nil {
				err = msgp.WrapError(err, "ShardSize")
				return
			}
		case "locate":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				err = msgp.WrapError(err, "Locate")
				return
			}
			if cap(z.Locate) >= int(zb0002) {
				z.Locate = (z.Locate)[:zb0002]
			} else {
				z.Locate = make([]string, zb0002)
			}
			for za0001 := range z.Locate {
				z.Locate[za0001], err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Locate", za0001)
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Version) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "sequence"
	err = en.Append(0x89, 0xa8, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Sequence)
	if err != nil {
		err = msgp.WrapError(err, "Sequence")
		return
	}
	// write "hash"
	err = en.Append(0xa4, 0x68, 0x61, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteString(z.Hash)
	if err != nil {
		err = msgp.WrapError(err, "Hash")
		return
	}
	// write "size"
	err = en.Append(0xa4, 0x73, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Size)
	if err != nil {
		err = msgp.WrapError(err, "Size")
		return
	}
	// write "ts"
	err = en.Append(0xa2, 0x74, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Ts)
	if err != nil {
		err = msgp.WrapError(err, "Ts")
		return
	}
	// write "ec_algo"
	err = en.Append(0xa7, 0x65, 0x63, 0x5f, 0x61, 0x6c, 0x67, 0x6f)
	if err != nil {
		return
	}
	err = en.WriteInt8(z.EcAlgo)
	if err != nil {
		err = msgp.WrapError(err, "EcAlgo")
		return
	}
	// write "data_shards"
	err = en.Append(0xab, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.DataShards)
	if err != nil {
		err = msgp.WrapError(err, "DataShards")
		return
	}
	// write "parity_shards"
	err = en.Append(0xad, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.ParityShards)
	if err != nil {
		err = msgp.WrapError(err, "ParityShards")
		return
	}
	// write "shard_size"
	err = en.Append(0xaa, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.ShardSize)
	if err != nil {
		err = msgp.WrapError(err, "ShardSize")
		return
	}
	// write "locate"
	err = en.Append(0xa6, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Locate)))
	if err != nil {
		err = msgp.WrapError(err, "Locate")
		return
	}
	for za0001 := range z.Locate {
		err = en.WriteString(z.Locate[za0001])
		if err != nil {
			err = msgp.WrapError(err, "Locate", za0001)
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Version) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "sequence"
	o = append(o, 0x89, 0xa8, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint64(o, z.Sequence)
	// string "hash"
	o = append(o, 0xa4, 0x68, 0x61, 0x73, 0x68)
	o = msgp.AppendString(o, z.Hash)
	// string "size"
	o = append(o, 0xa4, 0x73, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.Size)
	// string "ts"
	o = append(o, 0xa2, 0x74, 0x73)
	o = msgp.AppendInt64(o, z.Ts)
	// string "ec_algo"
	o = append(o, 0xa7, 0x65, 0x63, 0x5f, 0x61, 0x6c, 0x67, 0x6f)
	o = msgp.AppendInt8(o, z.EcAlgo)
	// string "data_shards"
	o = append(o, 0xab, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73)
	o = msgp.AppendInt32(o, z.DataShards)
	// string "parity_shards"
	o = append(o, 0xad, 0x70, 0x61, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x73)
	o = msgp.AppendInt32(o, z.ParityShards)
	// string "shard_size"
	o = append(o, 0xaa, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x69, 0x7a, 0x65)
	o = msgp.AppendInt64(o, z.ShardSize)
	// string "locate"
	o = append(o, 0xa6, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Locate)))
	for za0001 := range z.Locate {
		o = msgp.AppendString(o, z.Locate[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Version) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "sequence":
			z.Sequence, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Sequence")
				return
			}
		case "hash":
			z.Hash, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Hash")
				return
			}
		case "size":
			z.Size, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Size")
				return
			}
		case "ts":
			z.Ts, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ts")
				return
			}
		case "ec_algo":
			z.EcAlgo, bts, err = msgp.ReadInt8Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "EcAlgo")
				return
			}
		case "data_shards":
			z.DataShards, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "DataShards")
				return
			}
		case "parity_shards":
			z.ParityShards, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ParityShards")
				return
			}
		case "shard_size":
			z.ShardSize, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ShardSize")
				return
			}
		case "locate":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Locate")
				return
			}
			if cap(z.Locate) >= int(zb0002) {
				z.Locate = (z.Locate)[:zb0002]
			} else {
				z.Locate = make([]string, zb0002)
			}
			for za0001 := range z.Locate {
				z.Locate[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Locate", za0001)
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Version) Msgsize() (s int) {
	s = 1 + 9 + msgp.Uint64Size + 5 + msgp.StringPrefixSize + len(z.Hash) + 5 + msgp.Int64Size + 3 + msgp.Int64Size + 8 + msgp.Int8Size + 12 + msgp.Int32Size + 14 + msgp.Int32Size + 11 + msgp.Int64Size + 7 + msgp.ArrayHeaderSize
	for za0001 := range z.Locate {
		s += msgp.StringPrefixSize + len(z.Locate[za0001])
	}
	return
}
