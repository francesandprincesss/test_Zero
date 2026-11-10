package app

import (
	grpcApp "greeter/app/gRPC"
	"time"
)

type App struct {
	GRPCSrv *grpcApp.App
}

func New(grpcPort int, tokenTTL time.Duration) *App {
	grpcApp := grpcApp.New(grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}