package grpc

import (
	"context"
	"log"

	"google.golang.org/protobuf/proto"

	pb "github.com/KidPudel/client-service/proto/client"
)

type ClientUsecase interface {
	StartTrackingOrder(ctx context.Context) (bool, error)
}

type ClientServerOptions struct {
	ClientUsecase ClientUsecase
}

type ClientServer struct {
	pb.ClientServer
	options ClientServerOptions
}

func NewClientServer(options ClientServerOptions) *ClientServer {
	return &ClientServer{
		options: options,
	}
}

func (server *ClientServer) CheckOrder(ctx context.Context, _ *pb.Empty) (*pb.Status, error) {
	success, err := server.options.ClientUsecase.StartTrackingOrder(ctx)
	if err != nil {
		return nil, err
	}
	log.Println(success)
	return &pb.Status{
		Success: proto.Bool(success),
	}, nil
}
