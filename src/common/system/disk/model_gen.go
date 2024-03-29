package disk

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DevID) DecodeMsg(dc *msgp.Reader) (err error) {
	{
		var zb0001 string
		zb0001, err = dc.ReadString()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = DevID(zb0001)
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z DevID) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteString(string(z))
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z DevID) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendString(o, string(z))
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DevID) UnmarshalMsg(bts []byte) (o []byte, err error) {
	{
		var zb0001 string
		zb0001, bts, err = msgp.ReadStringBytes(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		(*z) = DevID(zb0001)
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z DevID) Msgsize() (s int) {
	s = msgp.StringPrefixSize + len(string(z))
	return
}

// DecodeMsg implements msgp.Decodable
func (z *IOStats) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "read_bytes":
			err = z.ReadBytes.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "ReadBytes")
				return
			}
		case "write_bytes":
			err = z.WriteBytes.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "WriteBytes")
				return
			}
		case "read_count":
			z.ReadCount, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "ReadCount")
				return
			}
		case "write_count":
			z.WriteCount, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "WriteCount")
				return
			}
		case "read_time":
			z.ReadTime, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "ReadTime")
				return
			}
		case "write_time":
			z.WriteTime, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "WriteTime")
				return
			}
		case "current_ios":
			z.CurrentIOs, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "CurrentIOs")
				return
			}
		case "io_time":
			z.IoTime, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "IoTime")
				return
			}
		case "weighted_io":
			z.WeightedIO, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "WeightedIO")
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
func (z *IOStats) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "read_bytes"
	err = en.Append(0x89, 0xaa, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73)
	if err != nil {
		return
	}
	err = z.ReadBytes.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "ReadBytes")
		return
	}
	// write "write_bytes"
	err = en.Append(0xab, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73)
	if err != nil {
		return
	}
	err = z.WriteBytes.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "WriteBytes")
		return
	}
	// write "read_count"
	err = en.Append(0xaa, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.ReadCount)
	if err != nil {
		err = msgp.WrapError(err, "ReadCount")
		return
	}
	// write "write_count"
	err = en.Append(0xab, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.WriteCount)
	if err != nil {
		err = msgp.WrapError(err, "WriteCount")
		return
	}
	// write "read_time"
	err = en.Append(0xa9, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.ReadTime)
	if err != nil {
		err = msgp.WrapError(err, "ReadTime")
		return
	}
	// write "write_time"
	err = en.Append(0xaa, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.WriteTime)
	if err != nil {
		err = msgp.WrapError(err, "WriteTime")
		return
	}
	// write "current_ios"
	err = en.Append(0xab, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6f, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.CurrentIOs)
	if err != nil {
		err = msgp.WrapError(err, "CurrentIOs")
		return
	}
	// write "io_time"
	err = en.Append(0xa7, 0x69, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.IoTime)
	if err != nil {
		err = msgp.WrapError(err, "IoTime")
		return
	}
	// write "weighted_io"
	err = en.Append(0xab, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6f)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.WeightedIO)
	if err != nil {
		err = msgp.WrapError(err, "WeightedIO")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *IOStats) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "read_bytes"
	o = append(o, 0x89, 0xaa, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73)
	o, err = z.ReadBytes.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "ReadBytes")
		return
	}
	// string "write_bytes"
	o = append(o, 0xab, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73)
	o, err = z.WriteBytes.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "WriteBytes")
		return
	}
	// string "read_count"
	o = append(o, 0xaa, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendUint64(o, z.ReadCount)
	// string "write_count"
	o = append(o, 0xab, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendUint64(o, z.WriteCount)
	// string "read_time"
	o = append(o, 0xa9, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendUint64(o, z.ReadTime)
	// string "write_time"
	o = append(o, 0xaa, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendUint64(o, z.WriteTime)
	// string "current_ios"
	o = append(o, 0xab, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6f, 0x73)
	o = msgp.AppendUint64(o, z.CurrentIOs)
	// string "io_time"
	o = append(o, 0xa7, 0x69, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendUint64(o, z.IoTime)
	// string "weighted_io"
	o = append(o, 0xab, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6f)
	o = msgp.AppendUint64(o, z.WeightedIO)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *IOStats) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "read_bytes":
			bts, err = z.ReadBytes.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReadBytes")
				return
			}
		case "write_bytes":
			bts, err = z.WriteBytes.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "WriteBytes")
				return
			}
		case "read_count":
			z.ReadCount, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReadCount")
				return
			}
		case "write_count":
			z.WriteCount, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "WriteCount")
				return
			}
		case "read_time":
			z.ReadTime, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "ReadTime")
				return
			}
		case "write_time":
			z.WriteTime, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "WriteTime")
				return
			}
		case "current_ios":
			z.CurrentIOs, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "CurrentIOs")
				return
			}
		case "io_time":
			z.IoTime, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "IoTime")
				return
			}
		case "weighted_io":
			z.WeightedIO, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "WeightedIO")
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
func (z *IOStats) Msgsize() (s int) {
	s = 1 + 11 + z.ReadBytes.Msgsize() + 12 + z.WriteBytes.Msgsize() + 11 + msgp.Uint64Size + 12 + msgp.Uint64Size + 10 + msgp.Uint64Size + 11 + msgp.Uint64Size + 12 + msgp.Uint64Size + 8 + msgp.Uint64Size + 12 + msgp.Uint64Size
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Info) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "total":
			err = z.Total.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Total")
				return
			}
		case "free":
			err = z.Free.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Free")
				return
			}
		case "used":
			err = z.Used.DecodeMsg(dc)
			if err != nil {
				err = msgp.WrapError(err, "Used")
				return
			}
		case "files":
			z.Files, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Files")
				return
			}
		case "f_free":
			z.Ffree, err = dc.ReadUint64()
			if err != nil {
				err = msgp.WrapError(err, "Ffree")
				return
			}
		case "major":
			z.Major, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "Major")
				return
			}
		case "minor":
			z.Minor, err = dc.ReadUint32()
			if err != nil {
				err = msgp.WrapError(err, "Minor")
				return
			}
		case "fs_type":
			z.FSType, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "FSType")
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
func (z *Info) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 8
	// write "total"
	err = en.Append(0x88, 0xa5, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	if err != nil {
		return
	}
	err = z.Total.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Total")
		return
	}
	// write "free"
	err = en.Append(0xa4, 0x66, 0x72, 0x65, 0x65)
	if err != nil {
		return
	}
	err = z.Free.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Free")
		return
	}
	// write "used"
	err = en.Append(0xa4, 0x75, 0x73, 0x65, 0x64)
	if err != nil {
		return
	}
	err = z.Used.EncodeMsg(en)
	if err != nil {
		err = msgp.WrapError(err, "Used")
		return
	}
	// write "files"
	err = en.Append(0xa5, 0x66, 0x69, 0x6c, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Files)
	if err != nil {
		err = msgp.WrapError(err, "Files")
		return
	}
	// write "f_free"
	err = en.Append(0xa6, 0x66, 0x5f, 0x66, 0x72, 0x65, 0x65)
	if err != nil {
		return
	}
	err = en.WriteUint64(z.Ffree)
	if err != nil {
		err = msgp.WrapError(err, "Ffree")
		return
	}
	// write "major"
	err = en.Append(0xa5, 0x6d, 0x61, 0x6a, 0x6f, 0x72)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Major)
	if err != nil {
		err = msgp.WrapError(err, "Major")
		return
	}
	// write "minor"
	err = en.Append(0xa5, 0x6d, 0x69, 0x6e, 0x6f, 0x72)
	if err != nil {
		return
	}
	err = en.WriteUint32(z.Minor)
	if err != nil {
		err = msgp.WrapError(err, "Minor")
		return
	}
	// write "fs_type"
	err = en.Append(0xa7, 0x66, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.FSType)
	if err != nil {
		err = msgp.WrapError(err, "FSType")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Info) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 8
	// string "total"
	o = append(o, 0x88, 0xa5, 0x74, 0x6f, 0x74, 0x61, 0x6c)
	o, err = z.Total.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Total")
		return
	}
	// string "free"
	o = append(o, 0xa4, 0x66, 0x72, 0x65, 0x65)
	o, err = z.Free.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Free")
		return
	}
	// string "used"
	o = append(o, 0xa4, 0x75, 0x73, 0x65, 0x64)
	o, err = z.Used.MarshalMsg(o)
	if err != nil {
		err = msgp.WrapError(err, "Used")
		return
	}
	// string "files"
	o = append(o, 0xa5, 0x66, 0x69, 0x6c, 0x65, 0x73)
	o = msgp.AppendUint64(o, z.Files)
	// string "f_free"
	o = append(o, 0xa6, 0x66, 0x5f, 0x66, 0x72, 0x65, 0x65)
	o = msgp.AppendUint64(o, z.Ffree)
	// string "major"
	o = append(o, 0xa5, 0x6d, 0x61, 0x6a, 0x6f, 0x72)
	o = msgp.AppendUint32(o, z.Major)
	// string "minor"
	o = append(o, 0xa5, 0x6d, 0x69, 0x6e, 0x6f, 0x72)
	o = msgp.AppendUint32(o, z.Minor)
	// string "fs_type"
	o = append(o, 0xa7, 0x66, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65)
	o = msgp.AppendString(o, z.FSType)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Info) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "total":
			bts, err = z.Total.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Total")
				return
			}
		case "free":
			bts, err = z.Free.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Free")
				return
			}
		case "used":
			bts, err = z.Used.UnmarshalMsg(bts)
			if err != nil {
				err = msgp.WrapError(err, "Used")
				return
			}
		case "files":
			z.Files, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Files")
				return
			}
		case "f_free":
			z.Ffree, bts, err = msgp.ReadUint64Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Ffree")
				return
			}
		case "major":
			z.Major, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Major")
				return
			}
		case "minor":
			z.Minor, bts, err = msgp.ReadUint32Bytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Minor")
				return
			}
		case "fs_type":
			z.FSType, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "FSType")
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
func (z *Info) Msgsize() (s int) {
	s = 1 + 6 + z.Total.Msgsize() + 5 + z.Free.Msgsize() + 5 + z.Used.Msgsize() + 6 + msgp.Uint64Size + 7 + msgp.Uint64Size + 6 + msgp.Uint32Size + 6 + msgp.Uint32Size + 8 + msgp.StringPrefixSize + len(z.FSType)
	return
}
