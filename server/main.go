package main

import (
	"context"
	//"grpc_tutorial/proto"
	"net"

	gen "github.com/rayhan/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	gen.RegisterAddServiceServer(srv, &server{})
	//.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *gen.Request) (*gen.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a + b

	return &gen.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *gen.Request) (*gen.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a * b

	return &gen.Response{Result: result}, nil
}
