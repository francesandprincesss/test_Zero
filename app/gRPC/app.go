package grpcapp

import (
	"fmt"
	"greeter/grpc/greeter"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	gRPCServer		*grpc.Server
	port    			int
}

func New(port int) *App {
	gRPCServer := grpc.NewServer()

 	greeter.Register(gRPCServer)

	return &App{
		gRPCServer: gRPCServer,
		port: port,			
	}
}

func (a *App) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf(":%w", err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf(":%w", err)
	}

	return nil
}