package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	albato "albato/proto"
)

type Albato struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Albato) Call(ctx context.Context, req *albato.Request, rsp *albato.Response) error {
	log.Info("Received Albato.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Albato) Stream(ctx context.Context, req *albato.StreamingRequest, stream albato.Albato_StreamStream) error {
	log.Infof("Received Albato.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&albato.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Albato) PingPong(ctx context.Context, stream albato.Albato_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&albato.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

// CustomUserRequest is a single request handler called via client.CustomUserRequest or the generated client code
func (e *Albato) CustomUserRequest(ctx context.Context, req *albato.UserRequest, rsp *albato.UserResponse) error {
	log.Info("Received Albato.CustomUserRequest request")
	rsp.Combined = "Combined string " + req.Msg + req.Msg2
	return nil
}
