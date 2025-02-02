// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld/hello.proto

package helloworld

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

type GetHelloRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHelloRequest) Reset()         { *m = GetHelloRequest{} }
func (m *GetHelloRequest) String() string { return proto.CompactTextString(m) }
func (*GetHelloRequest) ProtoMessage()    {}
func (*GetHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6199620e86bcc922, []int{0}
}

func (m *GetHelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHelloRequest.Unmarshal(m, b)
}
func (m *GetHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHelloRequest.Marshal(b, m, deterministic)
}
func (m *GetHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHelloRequest.Merge(m, src)
}
func (m *GetHelloRequest) XXX_Size() int {
	return xxx_messageInfo_GetHelloRequest.Size(m)
}
func (m *GetHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHelloRequest proto.InternalMessageInfo

func (m *GetHelloRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GetHelloResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHelloResponse) Reset()         { *m = GetHelloResponse{} }
func (m *GetHelloResponse) String() string { return proto.CompactTextString(m) }
func (*GetHelloResponse) ProtoMessage()    {}
func (*GetHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6199620e86bcc922, []int{1}
}

func (m *GetHelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHelloResponse.Unmarshal(m, b)
}
func (m *GetHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHelloResponse.Marshal(b, m, deterministic)
}
func (m *GetHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHelloResponse.Merge(m, src)
}
func (m *GetHelloResponse) XXX_Size() int {
	return xxx_messageInfo_GetHelloResponse.Size(m)
}
func (m *GetHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHelloResponse proto.InternalMessageInfo

func (m *GetHelloResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*GetHelloRequest)(nil), "hello.GetHelloRequest")
	proto.RegisterType((*GetHelloResponse)(nil), "hello.GetHelloResponse")
}

func init() { proto.RegisterFile("helloworld/hello.proto", fileDescriptor_6199620e86bcc922) }

var fileDescriptor_6199620e86bcc922 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x07, 0x33, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85,
	0x58, 0xc1, 0x1c, 0x25, 0x65, 0x2e, 0x7e, 0xf7, 0xd4, 0x12, 0x0f, 0x10, 0x3b, 0x28, 0xb5, 0xb0,
	0x34, 0xb5, 0xb8, 0x44, 0x48, 0x80, 0x8b, 0x39, 0xb7, 0x38, 0x5d, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc4, 0x54, 0x52, 0xe1, 0x12, 0x40, 0x28, 0x2a, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xc5,
	0x54, 0x65, 0x14, 0xc4, 0x25, 0x08, 0x56, 0x12, 0x0e, 0xb2, 0x2b, 0x38, 0xb5, 0xa8, 0x2c, 0x33,
	0x39, 0x55, 0xc8, 0x96, 0x8b, 0x03, 0xa6, 0x55, 0x48, 0x4c, 0x0f, 0xe2, 0x00, 0x34, 0x0b, 0xa5,
	0xc4, 0x31, 0xc4, 0x21, 0x76, 0x28, 0x31, 0x38, 0x35, 0x30, 0x72, 0x71, 0x26, 0xe7, 0xe7, 0x42,
	0x14, 0x38, 0x71, 0x81, 0xa5, 0x03, 0x40, 0xee, 0x0f, 0x60, 0x8c, 0x32, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x48, 0x2c, 0xc8, 0x4d, 0xcc, 0x4b, 0xd6,
	0x4f, 0xcc, 0x29, 0xc8, 0x4b, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0xd5, 0x07, 0x7b, 0x54, 0x3f,
	0x3d, 0x5f, 0x1f, 0x11, 0x04, 0x8b, 0x98, 0x98, 0x3d, 0x22, 0x22, 0x56, 0x31, 0xb1, 0x82, 0x4d,
	0x3a, 0x05, 0xa5, 0x1f, 0x31, 0x41, 0x5c, 0x1e, 0xe3, 0x1e, 0xe0, 0xe4, 0x9b, 0x5a, 0x92, 0x98,
	0x92, 0x58, 0x92, 0xf8, 0x0a, 0x2a, 0x97, 0xc4, 0x06, 0x36, 0xc6, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xa9, 0x1d, 0x73, 0x1b, 0x49, 0x01, 0x00, 0x00,
}
