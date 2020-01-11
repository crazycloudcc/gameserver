// Code generated by protoc-gen-go. DO NOT EDIT.
// source: errorcode.proto

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

type ErrorCode int32

const (
	ErrorCode_OK     ErrorCode = 0
	ErrorCode_UNKNOW ErrorCode = 1
)

var ErrorCode_name = map[int32]string{
	0: "OK",
	1: "UNKNOW",
}

var ErrorCode_value = map[string]int32{
	"OK":     0,
	"UNKNOW": 1,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94c848a11f627b0e, []int{0}
}

func init() {
	proto.RegisterEnum("protos.ErrorCode", ErrorCode_name, ErrorCode_value)
}

func init() { proto.RegisterFile("errorcode.proto", fileDescriptor_94c848a11f627b0e) }

var fileDescriptor_94c848a11f627b0e = []byte{
	// 81 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2d, 0x2a, 0xca,
	0x2f, 0x4a, 0xce, 0x4f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5,
	0x5a, 0xf2, 0x5c, 0x9c, 0xae, 0x20, 0x29, 0xe7, 0xfc, 0x94, 0x54, 0x21, 0x36, 0x2e, 0x26, 0x7f,
	0x6f, 0x01, 0x06, 0x21, 0x2e, 0x2e, 0xb6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x01, 0xc6, 0x24,
	0x88, 0x42, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xed, 0x53, 0xdc, 0x42, 0x00, 0x00,
	0x00,
}