// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: content-fetcher.proto

package cf

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

// ContentFetcherServiceClient is the client API for ContentFetcherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContentFetcherServiceClient interface {
	GetContent(ctx context.Context, in *GetContentRequest, opts ...grpc.CallOption) (*GetContentResponse, error)
}

type contentFetcherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentFetcherServiceClient(cc grpc.ClientConnInterface) ContentFetcherServiceClient {
	return &contentFetcherServiceClient{cc}
}

func (c *contentFetcherServiceClient) GetContent(ctx context.Context, in *GetContentRequest, opts ...grpc.CallOption) (*GetContentResponse, error) {
	out := new(GetContentResponse)
	err := c.cc.Invoke(ctx, "/cf.ContentFetcherService/GetContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentFetcherServiceServer is the server API for ContentFetcherService service.
// All implementations must embed UnimplementedContentFetcherServiceServer
// for forward compatibility
type ContentFetcherServiceServer interface {
	GetContent(context.Context, *GetContentRequest) (*GetContentResponse, error)
	mustEmbedUnimplementedContentFetcherServiceServer()
}

// UnimplementedContentFetcherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContentFetcherServiceServer struct {
}

func (UnimplementedContentFetcherServiceServer) GetContent(context.Context, *GetContentRequest) (*GetContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContent not implemented")
}
func (UnimplementedContentFetcherServiceServer) mustEmbedUnimplementedContentFetcherServiceServer() {}

// UnsafeContentFetcherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContentFetcherServiceServer will
// result in compilation errors.
type UnsafeContentFetcherServiceServer interface {
	mustEmbedUnimplementedContentFetcherServiceServer()
}

func RegisterContentFetcherServiceServer(s grpc.ServiceRegistrar, srv ContentFetcherServiceServer) {
	s.RegisterService(&ContentFetcherService_ServiceDesc, srv)
}

func _ContentFetcherService_GetContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentFetcherServiceServer).GetContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cf.ContentFetcherService/GetContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentFetcherServiceServer).GetContent(ctx, req.(*GetContentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ContentFetcherService_ServiceDesc is the grpc.ServiceDesc for ContentFetcherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContentFetcherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cf.ContentFetcherService",
	HandlerType: (*ContentFetcherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetContent",
			Handler:    _ContentFetcherService_GetContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content-fetcher.proto",
}
