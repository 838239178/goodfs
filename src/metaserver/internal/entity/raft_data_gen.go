package entity

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Dest) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int8
		zb0001, err = dc.ReadInt8()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = Dest(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Dest) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt8(int8(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Dest) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt8(o, int8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Dest) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int8
		zb0001, bts, err = msgp.ReadInt8Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = Dest(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Dest) Msgsize() (s int) {
	s = msgp.Int8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *LogType) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 int8
		zb0001, err = dc.ReadInt8()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = LogType(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z LogType) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteInt8(int8(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z LogType) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendInt8(o, int8(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *LogType) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 int8
		zb0001, bts, err = msgp.ReadInt8Bytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = LogType(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z LogType) Msgsize() (s int) {
	s = msgp.Int8Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *RaftData) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "type":
			{
				var zb0002 int8
				zb0002, err = dc.ReadInt8()
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = LogType(zb0002)
			}
		case "dest":
			{
				var zb0003 int8
				zb0003, err = dc.ReadInt8()
				if err != nil {
					err = msgp.WrapError(err, "Dest")
					return
				}
				z.Dest = Dest(zb0003)
			}
		case "name":
			z.Name, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "sequence":
			z.Sequence, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Sequence")
				return
			}
		case "version":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Version")
					return
				}
				z.Version = nil
			} else {
				if z.Version == nil {
					z.Version = new(Version)
				}
				err = z.Version.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Version")
					return
				}
			}
		case "metadata":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Metadata")
					return
				}
				z.Metadata = nil
			} else {
				if z.Metadata == nil {
					z.Metadata = new(Metadata)
				}
				err = z.Metadata.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Metadata")
					return
				}
			}
		case "bucket":
			if dc.IsNil() {
				err = dc.ReadNil()
				if err != nil {
					err = msgp.WrapError(err, "Bucket")
					return
				}
				z.Bucket = nil
			} else {
				if z.Bucket == nil {
					z.Bucket = new(Bucket)
				}
				err = z.Bucket.DecodeMsg(dc)
				if err != nil {
					err = msgp.WrapError(err, "Bucket")
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
func (z *RaftData) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "type"
	err = en.Append(0x87, 0xa4, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt8(int8(z.Type))
	if err != nil {
		err = msgp.WrapError(err, "Type")
		return
	}
	// write "dest"
	err = en.Append(0xa4, 0x64, 0x65, 0x73, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt8(int8(z.Dest))
	if err != nil {
		err = msgp.WrapError(err, "Dest")
		return
	}
	// write "name"
	err = en.Append(0xa4, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Name)
	if err != nil {
		err = msgp.WrapError(err, "Name")
		return
	}
	// write "sequence"
	err = en.Append(0xa8, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Sequence)
	if err != nil {
		err = msgp.WrapError(err, "Sequence")
		return
	}
	// write "version"
	err = en.Append(0xa7, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if err != nil {
		return
	}
	if z.Version == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Version.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Version")
			return
		}
	}
	// write "metadata"
	err = en.Append(0xa8, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61)
	if err != nil {
		return
	}
	if z.Metadata == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Metadata.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Metadata")
			return
		}
	}
	// write "bucket"
	err = en.Append(0xa6, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74)
	if err != nil {
		return
	}
	if z.Bucket == nil {
		err = en.WriteNil()
		if err != nil {
			return
		}
	} else {
		err = z.Bucket.EncodeMsg(en)
		if err != nil {
			err = msgp.WrapError(err, "Bucket")
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *RaftData) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "type"
	o = append(o, 0x87, 0xa4, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendInt8(o, int8(z.Type))
	// string "dest"
	o = append(o, 0xa4, 0x64, 0x65, 0x73, 0x74)
	o = msgp.AppendInt8(o, int8(z.Dest))
	// string "name"
	o = append(o, 0xa4, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "sequence"
	o = append(o, 0xa8, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65)
	o = msgp.AppendUint64(o, z.Sequence)
	// string "version"
	o = append(o, 0xa7, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e)
	if z.Version == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Version.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Version")
			return
		}
	}
	// string "metadata"
	o = append(o, 0xa8, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61)
	if z.Metadata == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Metadata.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Metadata")
			return
		}
	}
	// string "bucket"
	o = append(o, 0xa6, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74)
	if z.Bucket == nil {
		o = msgp.AppendNil(o)
	} else {
		o, err = z.Bucket.MarshalMsg(o)
		if err != nil {
			err = msgp.WrapError(err, "Bucket")
			return
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *RaftData) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "type":
			{
				var zb0002 int8
				zb0002, bts, err = msgp.ReadInt8Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Type")
					return
				}
				z.Type = LogType(zb0002)
			}
		case "dest":
			{
				var zb0003 int8
				zb0003, bts, err = msgp.ReadInt8Bytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Dest")
					return
				}
				z.Dest = Dest(zb0003)
			}
		case "name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Name")
				return
			}
		case "sequence":
			z.Sequence, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Sequence")
				return
			}
		case "version":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Version = nil
			} else {
				if z.Version == nil {
					z.Version = new(Version)
				}
				bts, err = z.Version.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Version")
					return
				}
			}
		case "metadata":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Metadata = nil
			} else {
				if z.Metadata == nil {
					z.Metadata = new(Metadata)
				}
				bts, err = z.Metadata.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Metadata")
					return
				}
			}
		case "bucket":
			if msgp.IsNil(bts) {
				bts, err = msgp.ReadNilBytes(bts)
				if err != nil {
					return
				}
				z.Bucket = nil
			} else {
				if z.Bucket == nil {
					z.Bucket = new(Bucket)
				}
				bts, err = z.Bucket.UnmarshalMsg(bts)
				if err != nil {
					err = msgp.WrapError(err, "Bucket")
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
func (z *RaftData) Msgsize() (s int) {
	s = 1 + 5 + msgp.Int8Size + 5 + msgp.Int8Size + 5 + msgp.StringPrefixSize + len(z.Name) + 9 + msgp.Uint64Size + 8
	if z.Version == nil {
		s += msgp.NilSize
	} else {
		s += z.Version.Msgsize()
	}
	s += 9
	if z.Metadata == nil {
		s += msgp.NilSize
	} else {
		s += z.Metadata.Msgsize()
	}
	s += 7
	if z.Bucket == nil {
		s += msgp.NilSize
	} else {
		s += z.Bucket.Msgsize()
	}
	return
}
