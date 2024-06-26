// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: compiler.proto

package compiler

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

// CompilerServiceClient is the client API for CompilerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompilerServiceClient interface {
	Execute(ctx context.Context, in *SourceCode, opts ...grpc.CallOption) (*Output, error)
}

type compilerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompilerServiceClient(cc grpc.ClientConnInterface) CompilerServiceClient {
	return &compilerServiceClient{cc}
}

func (c *compilerServiceClient) Execute(ctx context.Context, in *SourceCode, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := c.cc.Invoke(ctx, "/compiler.CompilerService/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompilerServiceServer is the server API for CompilerService service.
// All implementations must embed UnimplementedCompilerServiceServer
// for forward compatibility
type CompilerServiceServer interface {
	Execute(context.Context, *SourceCode) (*Output, error)
	mustEmbedUnimplementedCompilerServiceServer()
}

// UnimplementedCompilerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompilerServiceServer struct {
}

func (UnimplementedCompilerServiceServer) Execute(context.Context, *SourceCode) (*Output, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedCompilerServiceServer) mustEmbedUnimplementedCompilerServiceServer() {}

// UnsafeCompilerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompilerServiceServer will
// result in compilation errors.
type UnsafeCompilerServiceServer interface {
	mustEmbedUnimplementedCompilerServiceServer()
}

func RegisterCompilerServiceServer(s grpc.ServiceRegistrar, srv CompilerServiceServer) {
	s.RegisterService(&CompilerService_ServiceDesc, srv)
}

func _CompilerService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SourceCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompilerServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/compiler.CompilerService/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompilerServiceServer).Execute(ctx, req.(*SourceCode))
	}
	return interceptor(ctx, in, info, handler)
}

// CompilerService_ServiceDesc is the grpc.ServiceDesc for CompilerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompilerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "compiler.CompilerService",
	HandlerType: (*CompilerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _CompilerService_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "compiler.proto",
}
