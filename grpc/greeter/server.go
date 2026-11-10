package greeter

import (
	"context"
	greeter "greeter/gen/go"

	"google.golang.org/grpc"
)

type Greeter interface {
	SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error)
}

type serverAPI struct {
	greeter.UnimplementedGreeterServiceServer
}

func Register(gRPC *grpc.Server) {
	greeter.RegisterGreeterServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) SayHello(ctx context.Context, req *greeter.HelloRequest) (*greeter.HelloResponse, error) {
	name := req.GetName()

	response := &greeter.HelloResponse{
		Message: "Hello, " + name,
	}

	return response, nil
}