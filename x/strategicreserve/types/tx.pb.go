// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/strategicreserve/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("fury/strategicreserve/tx.proto", fileDescriptor_bbf4aff262295ba9) }

var fileDescriptor_bbf4aff262295ba9 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x4e, 0x4f, 0xd5,
	0x2f, 0x2e, 0x29, 0x4a, 0x2c, 0x49, 0x4d, 0xcf, 0x4c, 0x2e, 0x4a, 0x2d, 0x4e, 0x2d, 0x2a, 0x4b,
	0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x2f, 0x4e, 0x4f, 0xcd, 0x4b,
	0x2d, 0x29, 0xcf, 0x2f, 0xca, 0xd6, 0x2b, 0x4e, 0x4f, 0xd5, 0x43, 0x57, 0x69, 0xc4, 0xca, 0xc5,
	0xec, 0x5b, 0x9c, 0xee, 0xe4, 0x7b, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51,
	0xc6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xc5, 0xe9, 0xa9, 0xba,
	0x50, 0xd3, 0x40, 0x6c, 0xfd, 0x0a, 0x2c, 0x36, 0x57, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x6d,
	0x37, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xe4, 0x38, 0x67, 0x9e, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "merlinnetwork.fury.strategicreserve.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "fury/strategicreserve/tx.proto",
}
