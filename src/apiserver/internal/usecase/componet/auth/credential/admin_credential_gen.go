package credential

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *AdminCredential) DecodeMsg(dc *msgp.Reader) (err error) {
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
		case "username":
			z.Username, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Username")
				return
			}
		case "password":
			z.Password, err = dc.ReadString()
			if err != nil {
				err = msgp.WrapError(err, "Password")
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
func (z AdminCredential) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "username"
	err = en.Append(0x82, 0xa8, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.Username)
	if err != nil {
		err = msgp.WrapError(err, "Username")
		return
	}
	// write "password"
	err = en.Append(0xa8, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	if err != nil {
		return
	}
	err = en.WriteString(z.Password)
	if err != nil {
		err = msgp.WrapError(err, "Password")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z AdminCredential) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "username"
	o = append(o, 0x82, 0xa8, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Username)
	// string "password"
	o = append(o, 0xa8, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64)
	o = msgp.AppendString(o, z.Password)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *AdminCredential) UnmarshalMsg(bts []byte) (o []byte, err error) {
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
		case "username":
			z.Username, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Username")
				return
			}
		case "password":
			z.Password, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Password")
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
func (z AdminCredential) Msgsize() (s int) {
	s = 1 + 9 + msgp.StringPrefixSize + len(z.Username) + 9 + msgp.StringPrefixSize + len(z.Password)
	return
}
