package main

import (
	"context"
	"fmt"
	"log"

	pb "git.ecobin.ir/ecomicro/protobuf/foo/grpc"
	"git.ecobin.ir/ecomicro/transport"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFooServer
}

func (s *server) Bar(ctx context.Context, in *pb.BarRequest) (*pb.BarResponse, error) {
	log.Printf("Received: %v", in.UserId)
	return &pb.BarResponse{}, nil
}

func main() {
	var err error
	grpcServer, err := transport.NewGRPCServer(transport.GRPCConfig{
		IP:   "0.0.0.0",
		Port: 50001,
	}, "debug", func(g *grpc.Server) {
		pb.RegisterFooServer(g, &server{})
	})
	if err != nil {
		panic(err)
	}
	defer grpcServer.Shutdown()
	err = grpcServer.Serve()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

}
