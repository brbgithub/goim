// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message_ack.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	message_ack.proto

It has these top-level messages:
	MessageACK
*/
package pb

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

// 投递消息回执
type MessageACK struct {
	MessageId    int64 `protobuf:"varint,1,opt,name=message_id,json=messageId" json:"message_id,omitempty"`
	SyncSequence int64 `protobuf:"varint,2,opt,name=sync_sequence,json=syncSequence" json:"sync_sequence,omitempty"`
	ReceiveTime  int64 `protobuf:"varint,3,opt,name=receive_time,json=receiveTime" json:"receive_time,omitempty"`
}

func (m *MessageACK) Reset()                    { *m = MessageACK{} }
func (m *MessageACK) String() string            { return proto.CompactTextString(m) }
func (*MessageACK) ProtoMessage()               {}
func (*MessageACK) Descriptor() ([]byte, []int) { return messageACKFileDescriptor, []int{0} }

func (m *MessageACK) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *MessageACK) GetSyncSequence() int64 {
	if m != nil {
		return m.SyncSequence
	}
	return 0
}

func (m *MessageACK) GetReceiveTime() int64 {
	if m != nil {
		return m.ReceiveTime
	}
	return 0
}

func init() {
	proto.RegisterType((*MessageACK)(nil), "pb.MessageACK")
}

func init() { proto.RegisterFile("message_ack.proto", messageACKFileDescriptor) }

var messageACKFileDescriptor = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x8d, 0x4f, 0x4c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0x48, 0x52, 0x2a, 0xe6, 0xe2, 0xf2, 0x85, 0x48, 0x38, 0x3a, 0x7b, 0x0b, 0xc9, 0x72, 0x71, 0xc1,
	0x94, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x71, 0x42, 0x45, 0x3c, 0x53, 0x84,
	0x94, 0xb9, 0x78, 0x8b, 0x2b, 0xf3, 0x92, 0xe3, 0x8b, 0x53, 0x0b, 0x4b, 0x53, 0xf3, 0x92, 0x53,
	0x25, 0x98, 0xc0, 0x2a, 0x78, 0x40, 0x82, 0xc1, 0x50, 0x31, 0x21, 0x45, 0x2e, 0x9e, 0xa2, 0xd4,
	0xe4, 0xd4, 0xcc, 0xb2, 0xd4, 0xf8, 0x92, 0xcc, 0xdc, 0x54, 0x09, 0x66, 0xb0, 0x1a, 0x6e, 0xa8,
	0x58, 0x48, 0x66, 0x6e, 0x6a, 0x12, 0x1b, 0xd8, 0x7e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x30, 0x3c, 0x49, 0xf5, 0x94, 0x00, 0x00, 0x00,
}
