// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: raft_cmd.proto

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

// RaftCmdClient is the client API for RaftCmd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftCmdClient interface {
	Bootstrap(ctx context.Context, in *BootstrapReq, opts ...grpc.CallOption) (*Response, error)
	AddVoter(ctx context.Context, in *AddVoterReq, opts ...grpc.CallOption) (*Response, error)
	JoinLeader(ctx context.Context, in *JoinLeaderReq, opts ...grpc.CallOption) (*Response, error)
	RemoveFollower(ctx context.Context, in *RemoveFollowerReq, opts ...grpc.CallOption) (*Response, error)
	AppliedIndex(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*Response, error)
	Peers(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*Response, error)
	Config(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*Response, error)
	LeaveCluster(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*Response, error)
}

type raftCmdClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftCmdClient(cc grpc.ClientConnInterface) RaftCmdClient {
	return &raftCmdClient{cc}
}

func (c *raftCmdClient) Bootstrap(ctx context.Context, in *BootstrapReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.RaftCmd/Bootstrap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCmdClient) AddVoter(ctx context.Context, in *AddVoterReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.RaftCmd/AddVoter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCmdClient) JoinLeader(ctx context.Context, in *JoinLeaderReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.RaftCmd/JoinLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCmdClient) RemoveFollower(ctx context.Context, in *RemoveFollowerReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.RaftCmd/RemoveFollower", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCmdClient) AppliedIndex(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.RaftCmd/AppliedIndex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCmdClient) Peers(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.RaftCmd/Peers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCmdClient) Config(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.RaftCmd/Config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftCmdClient) LeaveCluster(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/proto.RaftCmd/LeaveCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftCmdServer is the server API for RaftCmd service.
// All implementations must embed UnimplementedRaftCmdServer
// for forward compatibility
type RaftCmdServer interface {
	Bootstrap(context.Context, *BootstrapReq) (*Response, error)
	AddVoter(context.Context, *AddVoterReq) (*Response, error)
	JoinLeader(context.Context, *JoinLeaderReq) (*Response, error)
	RemoveFollower(context.Context, *RemoveFollowerReq) (*Response, error)
	AppliedIndex(context.Context, *EmptyReq) (*Response, error)
	Peers(context.Context, *EmptyReq) (*Response, error)
	Config(context.Context, *EmptyReq) (*Response, error)
	LeaveCluster(context.Context, *EmptyReq) (*Response, error)
	mustEmbedUnimplementedRaftCmdServer()
}

// UnimplementedRaftCmdServer must be embedded to have forward compatible implementations.
type UnimplementedRaftCmdServer struct {
}

func (UnimplementedRaftCmdServer) Bootstrap(context.Context, *BootstrapReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bootstrap not implemented")
}
func (UnimplementedRaftCmdServer) AddVoter(context.Context, *AddVoterReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVoter not implemented")
}
func (UnimplementedRaftCmdServer) JoinLeader(context.Context, *JoinLeaderReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinLeader not implemented")
}
func (UnimplementedRaftCmdServer) RemoveFollower(context.Context, *RemoveFollowerReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFollower not implemented")
}
func (UnimplementedRaftCmdServer) AppliedIndex(context.Context, *EmptyReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppliedIndex not implemented")
}
func (UnimplementedRaftCmdServer) Peers(context.Context, *EmptyReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Peers not implemented")
}
func (UnimplementedRaftCmdServer) Config(context.Context, *EmptyReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Config not implemented")
}
func (UnimplementedRaftCmdServer) LeaveCluster(context.Context, *EmptyReq) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveCluster not implemented")
}
func (UnimplementedRaftCmdServer) mustEmbedUnimplementedRaftCmdServer() {}

// UnsafeRaftCmdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftCmdServer will
// result in compilation errors.
type UnsafeRaftCmdServer interface {
	mustEmbedUnimplementedRaftCmdServer()
}

func RegisterRaftCmdServer(s grpc.ServiceRegistrar, srv RaftCmdServer) {
	s.RegisterService(&RaftCmd_ServiceDesc, srv)
}

func _RaftCmd_Bootstrap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BootstrapReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCmdServer).Bootstrap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RaftCmd/Bootstrap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCmdServer).Bootstrap(ctx, req.(*BootstrapReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCmd_AddVoter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVoterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCmdServer).AddVoter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RaftCmd/AddVoter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCmdServer).AddVoter(ctx, req.(*AddVoterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCmd_JoinLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinLeaderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCmdServer).JoinLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RaftCmd/JoinLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCmdServer).JoinLeader(ctx, req.(*JoinLeaderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCmd_RemoveFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFollowerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCmdServer).RemoveFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RaftCmd/RemoveFollower",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCmdServer).RemoveFollower(ctx, req.(*RemoveFollowerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCmd_AppliedIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCmdServer).AppliedIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RaftCmd/AppliedIndex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCmdServer).AppliedIndex(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCmd_Peers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCmdServer).Peers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RaftCmd/Peers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCmdServer).Peers(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCmd_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCmdServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RaftCmd/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCmdServer).Config(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftCmd_LeaveCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftCmdServer).LeaveCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RaftCmd/LeaveCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftCmdServer).LeaveCluster(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RaftCmd_ServiceDesc is the grpc.ServiceDesc for RaftCmd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftCmd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RaftCmd",
	HandlerType: (*RaftCmdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bootstrap",
			Handler:    _RaftCmd_Bootstrap_Handler,
		},
		{
			MethodName: "AddVoter",
			Handler:    _RaftCmd_AddVoter_Handler,
		},
		{
			MethodName: "JoinLeader",
			Handler:    _RaftCmd_JoinLeader_Handler,
		},
		{
			MethodName: "RemoveFollower",
			Handler:    _RaftCmd_RemoveFollower_Handler,
		},
		{
			MethodName: "AppliedIndex",
			Handler:    _RaftCmd_AppliedIndex_Handler,
		},
		{
			MethodName: "Peers",
			Handler:    _RaftCmd_Peers_Handler,
		},
		{
			MethodName: "Config",
			Handler:    _RaftCmd_Config_Handler,
		},
		{
			MethodName: "LeaveCluster",
			Handler:    _RaftCmd_LeaveCluster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft_cmd.proto",
}
