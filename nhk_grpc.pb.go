// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: nhk.proto

package nhk_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NhkServiceClient is the client API for NhkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NhkServiceClient interface {
	GetNews(ctx context.Context, in *NewsRequest, opts ...grpc.CallOption) (*NewsReply, error)
	GetWords(ctx context.Context, in *WordRequest, opts ...grpc.CallOption) (*WordReply, error)
}

type nhkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNhkServiceClient(cc grpc.ClientConnInterface) NhkServiceClient {
	return &nhkServiceClient{cc}
}

func (c *nhkServiceClient) GetNews(ctx context.Context, in *NewsRequest, opts ...grpc.CallOption) (*NewsReply, error) {
	out := new(NewsReply)
	err := c.cc.Invoke(ctx, "/nhk.NhkService/GetNews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nhkServiceClient) GetWords(ctx context.Context, in *WordRequest, opts ...grpc.CallOption) (*WordReply, error) {
	out := new(WordReply)
	err := c.cc.Invoke(ctx, "/nhk.NhkService/GetWords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NhkServiceServer is the server API for NhkService service.
// All implementations must embed UnimplementedNhkServiceServer
// for forward compatibility
type NhkServiceServer interface {
	GetNews(context.Context, *NewsRequest) (*NewsReply, error)
	GetWords(context.Context, *WordRequest) (*WordReply, error)
	mustEmbedUnimplementedNhkServiceServer()
}

// UnimplementedNhkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNhkServiceServer struct {
}

func (UnimplementedNhkServiceServer) GetNews(context.Context, *NewsRequest) (*NewsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNews not implemented")
}
func (UnimplementedNhkServiceServer) GetWords(context.Context, *WordRequest) (*WordReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWords not implemented")
}
func (UnimplementedNhkServiceServer) mustEmbedUnimplementedNhkServiceServer() {}

// UnsafeNhkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NhkServiceServer will
// result in compilation errors.
type UnsafeNhkServiceServer interface {
	mustEmbedUnimplementedNhkServiceServer()
}

func RegisterNhkServiceServer(s grpc.ServiceRegistrar, srv NhkServiceServer) {
	s.RegisterService(&NhkService_ServiceDesc, srv)
}

func _NhkService_GetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NhkServiceServer).GetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nhk.NhkService/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NhkServiceServer).GetNews(ctx, req.(*NewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NhkService_GetWords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NhkServiceServer).GetWords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nhk.NhkService/GetWords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NhkServiceServer).GetWords(ctx, req.(*WordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NhkService_ServiceDesc is the grpc.ServiceDesc for NhkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NhkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nhk.NhkService",
	HandlerType: (*NhkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNews",
			Handler:    _NhkService_GetNews_Handler,
		},
		{
			MethodName: "GetWords",
			Handler:    _NhkService_GetWords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nhk.proto",
}
