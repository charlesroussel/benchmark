package main

import (
	"context"
	"fmt"
	"net"

	"github.com/ogury/profiling/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedAdServerServer
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAdServerServer(srv, &server{})
	reflection.Register(srv)

	fmt.Println("Running on 0.0.0.0:8080")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) Score(
	_ context.Context,
	request *proto.AdRequest) (*proto.AdResponse, error) {
	return &proto.AdResponse{
		Id: request.Id,
	}, nil
}
