package grpc

import (
	"common/logs"
	"common/pb"
	"common/util"
	"context"
	"errors"
	"metaserver/config"
	"metaserver/internal/usecase"
	"metaserver/internal/usecase/raftimpl"
	"net"

	raftServer "github.com/Jille/raft-grpc-transport"
	"github.com/hashicorp/raft"
	netGrpc "google.golang.org/grpc"
)

var log = logs.New("grpc-server")

type Server struct {
	*netGrpc.Server
	Port string
}

// NewRpcServer init a grpc raft server. if no available nodes return empty object
func NewRpcServer(cfg config.ClusterConfig, repo usecase.IMetadataRepo, serv1 usecase.IMetadataService, serv2 usecase.IHashSlotService) (*Server, *raftimpl.RaftWrapper) {
	server := netGrpc.NewServer(netGrpc.ChainUnaryInterceptor(
		CheckLocalUnary,
		CheckWritableUnary,
		CheckRaftEnabledUnary,
		CheckRaftLeaderUnary,
		CheckRaftNonLeaderUnary,
		AllowValidMetaServerUnary,
	), netGrpc.ChainStreamInterceptor(
		CheckWritableStreaming,
		AllowValidMetaServerStreaming,
	))
	// init raft service
	var raftWrapper *raftimpl.RaftWrapper
	if cfg.Enable {
		raftGrpcServ := raftServer.New(raft.ServerAddress(util.GetHostPort(cfg.Port)), []netGrpc.DialOption{netGrpc.WithInsecure()})
		raftWrapper = raftimpl.NewRaft(cfg, raftimpl.NewFSM(repo), raftGrpcServ.Transport())
		raftGrpcServ.Register(server)
		pb.RegisterRaftCmdServer(server, NewRaftCmdServer(raftWrapper))
	} else {
		raftWrapper = raftimpl.NewDisabledRaft()
	}
	// register hash-slot services
	pb.RegisterHashSlotServer(server, NewHashSlotServer(serv2))
	pb.RegisterMetadataApiServer(server, NewMetadataApiServer(serv1))
	return &Server{server, cfg.Port}, raftWrapper
}

func (r *Server) Shutdown(ctx context.Context) error {
	if r.Server == nil {
		return nil
	}
	finish := make(chan struct{})
	go func() {
		defer close(finish)
		r.Server.GracefulStop()
	}()
	select {
	case <-ctx.Done():
		r.Server.Stop()
		return errors.New("graceful stop grpc server timeout")
	case <-finish:
		return nil
	}
}

func (r *Server) ListenAndServe() error {
	if r.Server == nil {
		return nil
	}
	sock, err := net.Listen("tcp", util.GetHostPort(r.Port))
	if err != nil {
		panic(err)
	}
	log.Infof("server listening on %s", sock.Addr().String())
	return r.Serve(sock)
}
