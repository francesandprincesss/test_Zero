package main

import (
	"greeter/app"
	"time"
)

func main() {
	application := app.New(44044, 10 * time.Minute)

	application.GRPCSrv.Run()
}