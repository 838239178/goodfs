package grpc

import (
	"common/collection/set"
	"common/util"
	"context"
	"errors"
	"metaserver/internal/usecase/pool"

	netGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

var checkRaftEnabledMethods = set.OfString([]string{
	"/proto.RaftCmd/Bootstrap",
	"/proto.RaftCmd/AddVoter",
})

var checkRaftLeaderMethods = set.OfString([]string{
	"/proto.RaftCmd/AddVoter",
})

var checkRaftNonLeaderMethods = set.OfString([]string{
	"/proto.RaftCmd/JoinLeader",
})

var checkLocalMethods = set.OfString([]string{
	"/proto.RaftCmd/Bootstrap",
	"/proto.RaftCmd/AddVoter",
	"/proto.RaftCmd/JoinLeader",
	"/proto.RaftCmd/AddVoter",
	"/proto.HashSlot/StartMigration",
})

var checkWritableMethods = set.OfString([]string{
	"/proto.HashSlot/StartMigration",
	"/proto.HashSlot/PrepareMigration",
	"/proto.HashSlot/StreamingReceive",
})

var checkValidMetaServerMethods = set.OfString([]string{
	"/proto.HashSlot/PrepareMigration",
	"/proto.HashSlot/StreamingReceive",
})

func CheckRaftEnabledUnary(ctx context.Context, req interface{}, info *netGrpc.UnaryServerInfo, handler netGrpc.UnaryHandler) (any, error) {
	if checkRaftEnabledMethods.Contains(info.FullMethod) {
		if !pool.RaftWrapper.Enabled {
			return nil, errors.New("raft is not enabled")
		}
	}
	return handler(ctx, req)
}

func CheckRaftLeaderUnary(ctx context.Context, req interface{}, info *netGrpc.UnaryServerInfo, handler netGrpc.UnaryHandler) (any, error) {
	if checkRaftLeaderMethods.Contains(info.FullMethod) {
		if !pool.RaftWrapper.IsLeader() {
			return nil, errors.New("server is not a leader")
		}
	}
	return handler(ctx, req)
}

func CheckWritableUnary(ctx context.Context, req interface{}, info *netGrpc.UnaryServerInfo, handler netGrpc.UnaryHandler) (any, error) {
	if checkWritableMethods.Contains(info.FullMethod) {
		if pool.RaftWrapper.Enabled && !pool.RaftWrapper.IsLeader() {
			return nil, errors.New("server is not writable")
		}
	}
	return handler(ctx, req)
}

func CheckWritableStreaming(srv interface{}, ss netGrpc.ServerStream, info *netGrpc.StreamServerInfo, handler netGrpc.StreamHandler) error {
	if checkWritableMethods.Contains(info.FullMethod) {
		if pool.RaftWrapper.Enabled && !pool.RaftWrapper.IsLeader() {
			return errors.New("server is not writable")
		}
	}
	return handler(srv, ss)
}

func CheckRaftNonLeaderUnary(ctx context.Context, req interface{}, info *netGrpc.UnaryServerInfo, handler netGrpc.UnaryHandler) (any, error) {
	if checkRaftNonLeaderMethods.Contains(info.FullMethod) {
		if pool.RaftWrapper.IsLeader() {
			return nil, errors.New("server is a leader")
		}
	}
	return handler(ctx, req)
}

func CheckLocalUnary(ctx context.Context, req interface{}, info *netGrpc.UnaryServerInfo, handler netGrpc.UnaryHandler) (any, error) {
	if checkLocalMethods.Contains(info.FullMethod) {
		if pr, ok := peer.FromContext(ctx); ok {
			if util.GetHostFromAddr(pr.Addr.String()) != util.GetHost() {
				return nil, errors.New("deny")
			}
		} else {
			return nil, errors.New("get client ip fail")
		}
	}
	return handler(ctx, req)
}

func AllowValidMetaServerUnary(ctx context.Context, req interface{}, info *netGrpc.UnaryServerInfo, handler netGrpc.UnaryHandler) (any, error) {
	if checkValidMetaServerMethods.Contains(info.FullMethod) {
		//TODO
	}
	return handler(ctx, req)
}

func AllowValidMetaServerStreaming(srv interface{}, ss netGrpc.ServerStream, info *netGrpc.StreamServerInfo, handler netGrpc.StreamHandler) error {
	if checkValidMetaServerMethods.Contains(info.FullMethod) {
		//TODO
	}
	return handler(srv, ss)
}