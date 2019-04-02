// Code generated by protoc-gen-go. DO NOT EDIT.
// source: norm.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	norm.proto

It has these top-level messages:
	Norm
	NormAction
	NormPut
	NormGetKey
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Norm struct {
	NormId     []byte `protobuf:"bytes,1,opt,name=normId,proto3" json:"normId,omitempty"`
	CreateTime int64  `protobuf:"varint,2,opt,name=createTime" json:"createTime,omitempty"`
	Key        []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value      []byte `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Norm) Reset()                    { *m = Norm{} }
func (m *Norm) String() string            { return proto.CompactTextString(m) }
func (*Norm) ProtoMessage()               {}
func (*Norm) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Norm) GetNormId() []byte {
	if m != nil {
		return m.NormId
	}
	return nil
}

func (m *Norm) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Norm) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Norm) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type NormAction struct {
	// Types that are valid to be assigned to Value:
	//	*NormAction_Nput
	Value isNormAction_Value `protobuf_oneof:"value"`
	Ty    int32              `protobuf:"varint,5,opt,name=ty" json:"ty,omitempty"`
}

func (m *NormAction) Reset()                    { *m = NormAction{} }
func (m *NormAction) String() string            { return proto.CompactTextString(m) }
func (*NormAction) ProtoMessage()               {}
func (*NormAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isNormAction_Value interface {
	isNormAction_Value()
}

type NormAction_Nput struct {
	Nput *NormPut `protobuf:"bytes,1,opt,name=nput,oneof"`
}

func (*NormAction_Nput) isNormAction_Value() {}

func (m *NormAction) GetValue() isNormAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *NormAction) GetNput() *NormPut {
	if x, ok := m.GetValue().(*NormAction_Nput); ok {
		return x.Nput
	}
	return nil
}

func (m *NormAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NormAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NormAction_OneofMarshaler, _NormAction_OneofUnmarshaler, _NormAction_OneofSizer, []interface{}{
		(*NormAction_Nput)(nil),
	}
}

func _NormAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NormAction)
	// value
	switch x := m.Value.(type) {
	case *NormAction_Nput:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Nput); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NormAction.Value has unexpected type %T", x)
	}
	return nil
}

func _NormAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NormAction)
	switch tag {
	case 1: // value.nput
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NormPut)
		err := b.DecodeMessage(msg)
		m.Value = &NormAction_Nput{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NormAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NormAction)
	// value
	switch x := m.Value.(type) {
	case *NormAction_Nput:
		s := proto.Size(x.Nput)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NormPut struct {
	Key   []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *NormPut) Reset()                    { *m = NormPut{} }
func (m *NormPut) String() string            { return proto.CompactTextString(m) }
func (*NormPut) ProtoMessage()               {}
func (*NormPut) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NormPut) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *NormPut) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type NormGetKey struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (m *NormGetKey) Reset()                    { *m = NormGetKey{} }
func (m *NormGetKey) String() string            { return proto.CompactTextString(m) }
func (*NormGetKey) ProtoMessage()               {}
func (*NormGetKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *NormGetKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func init() {
	proto.RegisterType((*Norm)(nil), "types.Norm")
	proto.RegisterType((*NormAction)(nil), "types.NormAction")
	proto.RegisterType((*NormPut)(nil), "types.NormPut")
	proto.RegisterType((*NormGetKey)(nil), "types.NormGetKey")
}

func init() { proto.RegisterFile("norm.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcb, 0x2f, 0xca,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x4a, 0xe3,
	0x62, 0xf1, 0xcb, 0x2f, 0xca, 0x15, 0x12, 0xe3, 0x62, 0x03, 0x49, 0x7a, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0xf0, 0x04, 0x41, 0x79, 0x42, 0x72, 0x5c, 0x5c, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xa9,
	0x21, 0x99, 0xb9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x48, 0x22, 0x42, 0x02, 0x5c,
	0xcc, 0xd9, 0xa9, 0x95, 0x12, 0xcc, 0x60, 0x4d, 0x20, 0xa6, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62,
	0x4e, 0x69, 0xaa, 0x04, 0x0b, 0x58, 0x0c, 0xc2, 0x51, 0xf2, 0xe6, 0xe2, 0x02, 0xd9, 0xe3, 0x98,
	0x5c, 0x92, 0x99, 0x9f, 0x27, 0xa4, 0xc2, 0xc5, 0x92, 0x57, 0x50, 0x5a, 0x02, 0xb6, 0x8b, 0xdb,
	0x88, 0x4f, 0x0f, 0xec, 0x16, 0x3d, 0x90, 0x82, 0x80, 0xd2, 0x12, 0x0f, 0x86, 0x20, 0xb0, 0xac,
	0x10, 0x1f, 0x17, 0x53, 0x49, 0xa5, 0x04, 0xab, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x53, 0x49, 0xa5,
	0x13, 0x3b, 0xd4, 0x64, 0x25, 0x43, 0x2e, 0x76, 0xa8, 0x5a, 0x98, 0xfd, 0x8c, 0x58, 0xec, 0x67,
	0x42, 0xb6, 0x5f, 0x0e, 0x62, 0xbf, 0x7b, 0x6a, 0x89, 0x77, 0x6a, 0x25, 0xa6, 0xae, 0x24, 0x36,
	0x70, 0xa8, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x32, 0x24, 0xd0, 0x7e, 0x23, 0x01, 0x00,
	0x00,
}
