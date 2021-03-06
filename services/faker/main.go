package main

import (
	"faker/handler"
	pb "faker/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("faker"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterFakerHandler(srv.Server(), new(handler.Faker))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
