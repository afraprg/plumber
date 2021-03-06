// Code generated by protoc-gen-go. DO NOT EDIT.
// source: outbound.proto

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

// Emitted by the reader to HSB which is then consumed by the replayer
type Outbound struct {
	ReplayId             string   `protobuf:"bytes,1,opt,name=replay_id,json=replayId,proto3" json:"replay_id,omitempty"`
	Blob                 []byte   `protobuf:"bytes,2,opt,name=blob,proto3" json:"blob,omitempty"`
	Last                 bool     `protobuf:"varint,3,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Outbound) Reset()         { *m = Outbound{} }
func (m *Outbound) String() string { return proto.CompactTextString(m) }
func (*Outbound) ProtoMessage()    {}
func (*Outbound) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dbaa15aa01abbc0, []int{0}
}

func (m *Outbound) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Outbound.Unmarshal(m, b)
}
func (m *Outbound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Outbound.Marshal(b, m, deterministic)
}
func (m *Outbound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Outbound.Merge(m, src)
}
func (m *Outbound) XXX_Size() int {
	return xxx_messageInfo_Outbound.Size(m)
}
func (m *Outbound) XXX_DiscardUnknown() {
	xxx_messageInfo_Outbound.DiscardUnknown(m)
}

var xxx_messageInfo_Outbound proto.InternalMessageInfo

func (m *Outbound) GetReplayId() string {
	if m != nil {
		return m.ReplayId
	}
	return ""
}

func (m *Outbound) GetBlob() []byte {
	if m != nil {
		return m.Blob
	}
	return nil
}

func (m *Outbound) GetLast() bool {
	if m != nil {
		return m.Last
	}
	return false
}

func init() {
	proto.RegisterType((*Outbound)(nil), "events.Outbound")
}

func init() { proto.RegisterFile("outbound.proto", fileDescriptor_5dbaa15aa01abbc0) }

var fileDescriptor_5dbaa15aa01abbc0 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0xb1, 0x0a, 0xc2, 0x30,
	0x14, 0x00, 0x89, 0x4a, 0x69, 0x83, 0x38, 0x64, 0x2a, 0xb8, 0x14, 0xa7, 0x0e, 0xd2, 0x0c, 0xfe,
	0x81, 0x9b, 0x53, 0xa1, 0xa3, 0x8b, 0xe4, 0x25, 0xa1, 0x2d, 0xa4, 0x7d, 0x21, 0x79, 0x11, 0xfc,
	0x7b, 0x31, 0x75, 0x3b, 0xee, 0x96, 0xe3, 0x27, 0x4c, 0x04, 0x98, 0x56, 0xd3, 0xf9, 0x80, 0x84,
	0xa2, 0xb0, 0x6f, 0xbb, 0x52, 0xbc, 0xf4, 0xbc, 0xec, 0xff, 0x45, 0x9c, 0x79, 0x15, 0xac, 0x77,
	0xea, 0xf3, 0x9a, 0x4d, 0xcd, 0x1a, 0xd6, 0x56, 0x43, 0xb9, 0x89, 0x87, 0x11, 0x82, 0x1f, 0xc0,
	0x21, 0xd4, 0xbb, 0x86, 0xb5, 0xc7, 0x21, 0xf3, 0xcf, 0x39, 0x15, 0xa9, 0xde, 0x37, 0xac, 0x2d,
	0x87, 0xcc, 0xf7, 0xee, 0x79, 0x1d, 0x67, 0x9a, 0x12, 0x74, 0x1a, 0x17, 0x09, 0x8a, 0xf4, 0xa4,
	0x31, 0x78, 0x19, 0xf5, 0x64, 0x17, 0x15, 0x25, 0xa4, 0xd9, 0x19, 0x39, 0xa2, 0xdc, 0x06, 0xa0,
	0xc8, 0x3f, 0xb7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xa7, 0x9d, 0x20, 0xa1, 0x00, 0x00,
	0x00,
}
