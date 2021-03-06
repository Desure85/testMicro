package main

import (
	"albato/handler"
	pb "albato/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("albato"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterAlbatoHandler(srv.Server(), new(handler.Albato))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
