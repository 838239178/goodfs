package raft

import (
	"common/cmd"
	"common/util"
	"context"
	"fmt"
	"metaserver/internal/usecase/pb"
	"strings"

	"google.golang.org/grpc"
)

var address string

func init() {
	cmd.Register("raft", func(args []string) {
		if len(args) < 2 {
			fmt.Println("should input command: raft port [add_voter/join_leader/applied_index] [..]")
			return
		}
		address = util.GetHostPort(args[0])
		switch args[1] {
		case "add_voter":
			if len(args) < 3 {
				fmt.Println("add_voter require 'id,host:port,index'")
				return
			}
			addVoter(args[2:])
		case "join_leader":
			if len(args) < 3 {
				fmt.Println("join_leader require leader's address")
				return
			}
			joinLeader(args[2])
		case "applied_index":
			getAppliedIndex()
		case "bootstrap":
			if len(args) < 3 {
				bootstrap(nil)
			} else {
				bootstrap(args[2:])
			}
		case "peers":
			peers()
		default:
			fmt.Printf("no such command %s\n", args[1])
		}
	})
}

func getClient() (pb.RaftCmdClient, error) {
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// if cc.GetState() != connectivity.Ready {
	// 	fmt.Println("connect to port " + port + " fail")
	// 	return nil, errors.New("connect fail")
	// }
	return pb.NewRaftCmdClient(cc), nil
}

func addVoter(voters []string) {
	cli, err := getClient()
	if err != nil {
		return
	}
	req := &pb.AddVoterReq{Voters: make([]*pb.Voter, len(voters))}
	for i, v := range voters {
		args := strings.Split(v, ",")
		if len(args) != 3 {
			fmt.Printf("voter (%s) format error, reuqire id,host:port,index\n", v)
			continue
		}
		req.Voters[i] = &pb.Voter{
			Id:        args[0],
			Address:   args[1],
			PrevIndex: util.ToUint64(args[2]),
		}
	}
	resp, err := cli.AddVoter(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("success: %v, message: %s\n", resp.Success, resp.Message)
	}
}

func joinLeader(leader string) {
	cli, err := getClient()
	if err != nil {
		return
	}
	resp, err := cli.JoinLeader(context.Background(), &pb.JoinLeaderReq{Address: leader})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("success: %v, message: %s\n", resp.Success, resp.Message)
	}
}

func getAppliedIndex() {
	cli, err := getClient()
	if err != nil {
		return
	}
	resp, err := cli.AppliedIndex(context.Background(), &pb.EmptyReq{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("success: %v, message: %s\n", resp.Success, resp.Message)
	}
}

func bootstrap(voters []string) {
	cli, err := getClient()
	if err != nil {
		return
	}
	req := &pb.BootstrapReq{Services: make([]*pb.RaftServerItem, len(voters))}
	for i, v := range voters {
		args := strings.Split(v, ",")
		if len(args) != 2 {
			fmt.Printf("voter (%s) format error, reuqire id,host:port\n", v)
			continue
		}
		req.Services[i] = &pb.RaftServerItem{
			Id:      args[0],
			Address: args[1],
		}
	}
	resp, err := cli.Bootstrap(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("success: %v, message: %s\n", resp.Success, resp.Message)
	}
}

func peers() {
	cli, err := getClient()
	if err != nil {
		return
	}
	resp, err := cli.Peers(context.Background(), &pb.EmptyReq{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("success: %v, message: %s\n", resp.Success, resp.Message)
	}
}
