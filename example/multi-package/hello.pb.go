// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello.proto

package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bar "github.com/quintans/gripmock/example/multi-package/bar"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GripmockClient is the client API for Gripmock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GripmockClient interface {
	Greet(ctx context.Context, in *bar.Bar, opts ...grpc.CallOption) (*Response, error)
}

type gripmockClient struct {
	cc *grpc.ClientConn
}

func NewGripmockClient(cc *grpc.ClientConn) GripmockClient {
	return &gripmockClient{cc}
}

func (c *gripmockClient) Greet(ctx context.Context, in *bar.Bar, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/hello.Gripmock/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GripmockServer is the server API for Gripmock service.
type GripmockServer interface {
	Greet(context.Context, *bar.Bar) (*Response, error)
}

func RegisterGripmockServer(s *grpc.Server, srv GripmockServer) {
	s.RegisterService(&_Gripmock_serviceDesc, srv)
}

func _Gripmock_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bar.Bar)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GripmockServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Gripmock/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GripmockServer).Greet(ctx, req.(*bar.Bar))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gripmock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Gripmock",
	HandlerType: (*GripmockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _Gripmock_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor_hello_86daf9ea719884b9) }

var fileDescriptor_hello_86daf9ea719884b9 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0xa4, 0x78, 0x93, 0x12, 0x8b,
	0xf4, 0x93, 0x12, 0x8b, 0x20, 0xa2, 0x52, 0x9c, 0x69, 0xf9, 0x50, 0x05, 0x46, 0x7a, 0x5c, 0x1c,
	0xee, 0x45, 0x99, 0x05, 0xb9, 0xf9, 0xc9, 0xd9, 0x42, 0x4a, 0x5c, 0xac, 0xee, 0x45, 0xa9, 0xa9,
	0x25, 0x42, 0x1c, 0x7a, 0x20, 0xb5, 0x4e, 0x89, 0x45, 0x52, 0xfc, 0x7a, 0x10, 0xd3, 0x82, 0x52,
	0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x93, 0xd8, 0xc0, 0xda, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xa4, 0x51, 0x1d, 0xb3, 0x66, 0x00, 0x00, 0x00,
}
