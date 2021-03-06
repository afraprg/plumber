// Code generated by protoc-gen-go. DO NOT EDIT.
// source: collect.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// This type is used to convey info about a collection request; it includes
// auth details that enables the collector to associate intake events with a
// particular team's collection request.
type Collect struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Schema               *Schema  `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Collect) Reset()         { *m = Collect{} }
func (m *Collect) String() string { return proto.CompactTextString(m) }
func (*Collect) ProtoMessage()    {}
func (*Collect) Descriptor() ([]byte, []int) {
	return fileDescriptor_37ca5d531a1e5e54, []int{0}
}

func (m *Collect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collect.Unmarshal(m, b)
}
func (m *Collect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collect.Marshal(b, m, deterministic)
}
func (m *Collect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collect.Merge(m, src)
}
func (m *Collect) XXX_Size() int {
	return xxx_messageInfo_Collect.Size(m)
}
func (m *Collect) XXX_DiscardUnknown() {
	xxx_messageInfo_Collect.DiscardUnknown(m)
}

var xxx_messageInfo_Collect proto.InternalMessageInfo

func (m *Collect) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Collect) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Collect) GetSchema() *Schema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func init() {
	proto.RegisterType((*Collect)(nil), "events.Collect")
}

func init() { proto.RegisterFile("collect.proto", fileDescriptor_37ca5d531a1e5e54) }

var fileDescriptor_37ca5d531a1e5e54 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0xc9,
	0x49, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2d, 0x4b, 0xcd, 0x2b,
	0x29, 0x96, 0xe2, 0x29, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x84, 0x88, 0x2a, 0x85, 0x73, 0xb1, 0x3b,
	0x43, 0x94, 0x09, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31,
	0x65, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0x96, 0xe4, 0x67, 0xa7, 0xe6, 0x49, 0x30, 0x81, 0x85, 0x20,
	0x1c, 0x21, 0x35, 0x2e, 0x36, 0x88, 0x01, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x7c, 0x7a,
	0x10, 0x73, 0xf5, 0x82, 0xc1, 0xa2, 0x41, 0x50, 0x59, 0x27, 0xbd, 0x28, 0x9d, 0xf4, 0xcc, 0x92,
	0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0xa4, 0xc4, 0x92, 0xe4, 0x8c, 0xe4, 0xfc, 0xa2,
	0x02, 0x7d, 0x88, 0x74, 0xb1, 0x7e, 0x52, 0x69, 0x66, 0x4e, 0x8a, 0x7e, 0x7a, 0xbe, 0x3e, 0x44,
	0x7b, 0x12, 0x1b, 0xd8, 0x3d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0x42, 0x95, 0xbd,
	0xb6, 0x00, 0x00, 0x00,
}
