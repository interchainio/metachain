// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metachain/v1beta/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

func init() { proto.RegisterFile("metachain/v1beta/querier.proto", fileDescriptor_a4eb7742e5632b88) }

var fileDescriptor_a4eb7742e5632b88 = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0x4d, 0x2d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0xd4, 0x2f, 0x2c, 0x4d,
	0x2d, 0xca, 0x4c, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x86, 0xcb, 0xeb, 0x21,
	0x58, 0x10, 0x95, 0x86, 0x52, 0x5a, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xfa, 0x49, 0x89, 0xc5,
	0xa9, 0x60, 0x7d, 0x95, 0x50, 0x43, 0x0c, 0xf5, 0x0b, 0x12, 0xd3, 0x33, 0xf3, 0x12, 0x4b, 0x32,
	0xf3, 0xf3, 0x20, 0x06, 0x19, 0xb1, 0x73, 0xb1, 0x06, 0x82, 0x54, 0x38, 0x05, 0x9e, 0x78, 0x24,
	0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78,
	0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x79, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e,
	0x72, 0x7e, 0xae, 0x7e, 0x66, 0x5e, 0x49, 0x6a, 0x11, 0xd8, 0xb6, 0xa4, 0xd4, 0xa2, 0x9c, 0xcc,
	0x3c, 0x7d, 0x84, 0x3b, 0x2b, 0x90, 0xd8, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x2b,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x01, 0x30, 0xb4, 0xcd, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metachain.metachain.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "metachain/v1beta/querier.proto",
}