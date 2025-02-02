package presentation

import (
	"context"

	hellopb "github.com/victorlin12345/ddd-template/internal/pkg/gen"
	"google.golang.org/grpc"
)

type HelloController struct {
	hellopb.UnimplementedMyServiceServer
}

func NewHelloController(server *grpc.Server) hellopb.MyServiceServer {
	ctr := &HelloController{}
	hellopb.RegisterMyServiceServer(server, ctr)

	return ctr
}

func (c *HelloController) SayHello(context.Context, *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{}, nil
}
