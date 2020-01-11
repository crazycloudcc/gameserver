// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package protos

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

////////////////////////////////////////////////////////////////
//枚举
type TestState int32

const (
	TestState_State_0 TestState = 0
	TestState_State_1 TestState = 1
)

var TestState_name = map[int32]string{
	0: "State_0",
	1: "State_1",
}

var TestState_value = map[string]int32{
	"State_0": 0,
	"State_1": 1,
}

func (x TestState) String() string {
	return proto.EnumName(TestState_name, int32(x))
}

func (TestState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type TestInfo struct {
	Test1                string   `protobuf:"bytes,1,opt,name=test1,proto3" json:"test1,omitempty"`
	Test2                int32    `protobuf:"varint,2,opt,name=test2,proto3" json:"test2,omitempty"`
	Test3                int32    `protobuf:"varint,3,opt,name=test3,proto3" json:"test3,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestInfo) Reset()         { *m = TestInfo{} }
func (m *TestInfo) String() string { return proto.CompactTextString(m) }
func (*TestInfo) ProtoMessage()    {}
func (*TestInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *TestInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestInfo.Unmarshal(m, b)
}
func (m *TestInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestInfo.Marshal(b, m, deterministic)
}
func (m *TestInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestInfo.Merge(m, src)
}
func (m *TestInfo) XXX_Size() int {
	return xxx_messageInfo_TestInfo.Size(m)
}
func (m *TestInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TestInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TestInfo proto.InternalMessageInfo

func (m *TestInfo) GetTest1() string {
	if m != nil {
		return m.Test1
	}
	return ""
}

func (m *TestInfo) GetTest2() int32 {
	if m != nil {
		return m.Test2
	}
	return 0
}

func (m *TestInfo) GetTest3() int32 {
	if m != nil {
		return m.Test3
	}
	return 0
}

func init() {
	proto.RegisterEnum("protos.TestState", TestState_name, TestState_value)
	proto.RegisterType((*TestInfo)(nil), "protos.TestInfo")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x3e, 0x5c,
	0x1c, 0x21, 0xa9, 0xc5, 0x25, 0x9e, 0x79, 0x69, 0xf9, 0x42, 0x22, 0x5c, 0xac, 0x25, 0xa9, 0xc5,
	0x25, 0x86, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x0e, 0x4c, 0xd4, 0x48, 0x82, 0x49,
	0x81, 0x51, 0x83, 0x15, 0x22, 0x6a, 0x04, 0x13, 0x35, 0x96, 0x60, 0x46, 0x88, 0x1a, 0x6b, 0xa9,
	0x72, 0x71, 0x82, 0x4c, 0x0b, 0x2e, 0x49, 0x2c, 0x49, 0x15, 0xe2, 0xe6, 0x62, 0x07, 0x33, 0xe2,
	0x0d, 0x04, 0x18, 0x10, 0x1c, 0x43, 0x01, 0xc6, 0x24, 0x88, 0xe5, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x95, 0xb0, 0xc3, 0x5f, 0x93, 0x00, 0x00, 0x00,
}
