package wechat

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

type VideoConfig struct {
}

func (m *VideoConfig) Reset()                    { *m = VideoConfig{} }
func (m *VideoConfig) String() string            { return proto.CompactTextString(m) }
func (*VideoConfig) ProtoMessage()               {}
func (*VideoConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*VideoConfig)(nil), "v2ray.core.transport.internet.headers.wechat.VideoConfig")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/headers/wechat/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x2d, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x29, 0x4a, 0xcc, 0x2b,
	0x2e, 0xc8, 0x2f, 0x2a, 0xd1, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0x2d, 0xd1, 0xcf, 0x48,
	0x4d, 0x4c, 0x49, 0x2d, 0x2a, 0xd6, 0x2f, 0x4f, 0x4d, 0xce, 0x48, 0x2c, 0xd1, 0x4f, 0xce, 0xcf,
	0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x81, 0x69, 0x2f, 0x4a, 0xd5,
	0x83, 0x6b, 0xd5, 0x83, 0x69, 0xd5, 0x83, 0x6a, 0xd5, 0x83, 0x68, 0x55, 0xe2, 0xe5, 0xe2, 0x0e,
	0xcb, 0x4c, 0x49, 0xcd, 0x77, 0x06, 0x1b, 0xe1, 0x54, 0xc6, 0x65, 0x90, 0x9c, 0x9f, 0xab, 0x47,
	0x8a, 0x11, 0x4e, 0xdc, 0x10, 0xbd, 0x01, 0x20, 0xdb, 0xa3, 0xd8, 0x20, 0x82, 0xab, 0x98, 0x74,
	0xc2, 0x8c, 0x82, 0x12, 0x2b, 0xf5, 0x9c, 0x41, 0x66, 0x84, 0xc0, 0xcd, 0xf0, 0x84, 0x99, 0xe1,
	0x01, 0x35, 0x23, 0x1c, 0xac, 0x3c, 0x89, 0x0d, 0xec, 0x76, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0x6a, 0x0f, 0x19, 0xfc, 0x00, 0x00, 0x00,
}