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
	GetNews(ctx context.Context, in *NewsRequest, opts ...grpc.CallOption) (NhkService_GetNewsClient, error)
}

type nhkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNhkServiceClient(cc grpc.ClientConnInterface) NhkServiceClient {
	return &nhkServiceClient{cc}
}

func (c *nhkServiceClient) GetNews(ctx context.Context, in *NewsRequest, opts ...grpc.CallOption) (NhkService_GetNewsClient, error) {
	stream, err := c.cc.NewStream(ctx, &NhkService_ServiceDesc.Streams[0], "/nhk.NhkService/GetNews", opts...)
	if err != nil {
		return nil, err
	}
	x := &nhkServiceGetNewsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NhkService_GetNewsClient interface {
	Recv() (*News, error)
	grpc.ClientStream
}

type nhkServiceGetNewsClient struct {
	grpc.ClientStream
}

func (x *nhkServiceGetNewsClient) Recv() (*News, error) {
	m := new(News)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NhkServiceServer is the server API for NhkService service.
// All implementations must embed UnimplementedNhkServiceServer
// for forward compatibility
type NhkServiceServer interface {
	GetNews(*NewsRequest, NhkService_GetNewsServer) error
	mustEmbedUnimplementedNhkServiceServer()
}

// UnimplementedNhkServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNhkServiceServer struct {
}

func (UnimplementedNhkServiceServer) GetNews(*NewsRequest, NhkService_GetNewsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNews not implemented")
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

func _NhkService_GetNews_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NewsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NhkServiceServer).GetNews(m, &nhkServiceGetNewsServer{stream})
}

type NhkService_GetNewsServer interface {
	Send(*News) error
	grpc.ServerStream
}

type nhkServiceGetNewsServer struct {
	grpc.ServerStream
}

func (x *nhkServiceGetNewsServer) Send(m *News) error {
	return x.ServerStream.SendMsg(m)
}

// NhkService_ServiceDesc is the grpc.ServiceDesc for NhkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NhkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nhk.NhkService",
	HandlerType: (*NhkServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNews",
			Handler:       _NhkService_GetNews_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "nhk.proto",
}
