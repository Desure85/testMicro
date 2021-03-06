package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	faker "faker/proto"
	fakerData "syreclabs.com/go/faker"
)

type Faker struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Faker) Call(ctx context.Context, req *faker.Request, rsp *faker.Response) error {
	log.Info("Received Faker.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Faker) Stream(ctx context.Context, req *faker.StreamingRequest, stream faker.Faker_StreamStream) error {
	log.Infof("Received Faker.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&faker.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Faker) PingPong(ctx context.Context, stream faker.Faker_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&faker.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

// FakeAddress is a bidirectional stream handler called via client.FakeAddress or the generated client code
func (e *Faker) FakeAddress(ctx context.Context, req *faker.FakeSeed, rsp *faker.FakeAddressResponse) error {
	log.Info("Received Faker.FakeAddress request")
	if req.Seed != 0 {
		fakerData.Seed(req.Seed)
	}
	rsp.City = fakerData.Address().City()
	rsp.StreetAddress = fakerData.Address().StreetAddress()
	rsp.SecondaryAddress = fakerData.Address().SecondaryAddress()
	rsp.StreetName = fakerData.Address().StreetName()
	rsp.BuildingNumber = fakerData.Address().BuildingNumber()
	rsp.Postcode = fakerData.Address().Postcode()
	return nil
}
