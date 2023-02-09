// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: metadata.proto

package pb

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

// MetadataApiClient is the client API for MetadataApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetadataApiClient interface {
	GetVersionsByHash(ctx context.Context, in *MetaReq, opts ...grpc.CallOption) (*Msgpack, error)
	GetBucket(ctx context.Context, in *MetaReq, opts ...grpc.CallOption) (*Msgpack, error)
	GetMetadata(ctx context.Context, in *MetaReq, opts ...grpc.CallOption) (*Msgpack, error)
	GetVersion(ctx context.Context, in *MetaReq, opts ...grpc.CallOption) (*Msgpack, error)
	GetPeers(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*StringsResp, error)
}

type metadataApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMetadataApiClient(cc grpc.ClientConnInterface) MetadataApiClient {
	return &metadataApiClient{cc}
}

func (c *metadataApiClient) GetVersionsByHash(ctx context.Context, in *MetaReq, opts ...grpc.CallOption) (*Msgpack, error) {
	out := new(Msgpack)
	err := c.cc.Invoke(ctx, "/proto.MetadataApi/getVersionsByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataApiClient) GetBucket(ctx context.Context, in *MetaReq, opts ...grpc.CallOption) (*Msgpack, error) {
	out := new(Msgpack)
	err := c.cc.Invoke(ctx, "/proto.MetadataApi/getBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataApiClient) GetMetadata(ctx context.Context, in *MetaReq, opts ...grpc.CallOption) (*Msgpack, error) {
	out := new(Msgpack)
	err := c.cc.Invoke(ctx, "/proto.MetadataApi/getMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataApiClient) GetVersion(ctx context.Context, in *MetaReq, opts ...grpc.CallOption) (*Msgpack, error) {
	out := new(Msgpack)
	err := c.cc.Invoke(ctx, "/proto.MetadataApi/getVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataApiClient) GetPeers(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*StringsResp, error) {
	out := new(StringsResp)
	err := c.cc.Invoke(ctx, "/proto.MetadataApi/getPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetadataApiServer is the server API for MetadataApi service.
// All implementations must embed UnimplementedMetadataApiServer
// for forward compatibility
type MetadataApiServer interface {
	GetVersionsByHash(context.Context, *MetaReq) (*Msgpack, error)
	GetBucket(context.Context, *MetaReq) (*Msgpack, error)
	GetMetadata(context.Context, *MetaReq) (*Msgpack, error)
	GetVersion(context.Context, *MetaReq) (*Msgpack, error)
	GetPeers(context.Context, *EmptyReq) (*StringsResp, error)
	mustEmbedUnimplementedMetadataApiServer()
}

// UnimplementedMetadataApiServer must be embedded to have forward compatible implementations.
type UnimplementedMetadataApiServer struct {
}

func (UnimplementedMetadataApiServer) GetVersionsByHash(context.Context, *MetaReq) (*Msgpack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersionsByHash not implemented")
}
func (UnimplementedMetadataApiServer) GetBucket(context.Context, *MetaReq) (*Msgpack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBucket not implemented")
}
func (UnimplementedMetadataApiServer) GetMetadata(context.Context, *MetaReq) (*Msgpack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetadata not implemented")
}
func (UnimplementedMetadataApiServer) GetVersion(context.Context, *MetaReq) (*Msgpack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedMetadataApiServer) GetPeers(context.Context, *EmptyReq) (*StringsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeers not implemented")
}
func (UnimplementedMetadataApiServer) mustEmbedUnimplementedMetadataApiServer() {}

// UnsafeMetadataApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetadataApiServer will
// result in compilation errors.
type UnsafeMetadataApiServer interface {
	mustEmbedUnimplementedMetadataApiServer()
}

func RegisterMetadataApiServer(s grpc.ServiceRegistrar, srv MetadataApiServer) {
	s.RegisterService(&MetadataApi_ServiceDesc, srv)
}

func _MetadataApi_GetVersionsByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataApiServer).GetVersionsByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetadataApi/getVersionsByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataApiServer).GetVersionsByHash(ctx, req.(*MetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataApi_GetBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataApiServer).GetBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetadataApi/getBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataApiServer).GetBucket(ctx, req.(*MetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataApi_GetMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataApiServer).GetMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetadataApi/getMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataApiServer).GetMetadata(ctx, req.(*MetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataApi_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetaReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataApiServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetadataApi/getVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataApiServer).GetVersion(ctx, req.(*MetaReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetadataApi_GetPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataApiServer).GetPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MetadataApi/getPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataApiServer).GetPeers(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MetadataApi_ServiceDesc is the grpc.ServiceDesc for MetadataApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetadataApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MetadataApi",
	HandlerType: (*MetadataApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getVersionsByHash",
			Handler:    _MetadataApi_GetVersionsByHash_Handler,
		},
		{
			MethodName: "getBucket",
			Handler:    _MetadataApi_GetBucket_Handler,
		},
		{
			MethodName: "getMetadata",
			Handler:    _MetadataApi_GetMetadata_Handler,
		},
		{
			MethodName: "getVersion",
			Handler:    _MetadataApi_GetVersion_Handler,
		},
		{
			MethodName: "getPeers",
			Handler:    _MetadataApi_GetPeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metadata.proto",
}
