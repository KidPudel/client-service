package client

import (
	"context"
	"io"
	"log"
	"math/rand"

	"google.golang.org/protobuf/proto"

	deliveryPb "github.com/KidPudel/client-service/proto/delivery"
)

type ClientUsecase struct {
	deliveryClient deliveryPb.DeliveryClient
}

func NewClientUsecase(deliveryClient deliveryPb.DeliveryClient) ClientUsecase {
	return ClientUsecase{
		deliveryClient: deliveryClient,
	}
}

func (u ClientUsecase) StartTrackingOrder(ctx context.Context) (bool, error) {
	stream, err := u.deliveryClient.FindEachOther(ctx)
	if err != nil {
		return false, err
	}

	waitCh := make(chan int)

	// we want to receive output at the same time
	go func() {
		for {
			position, err := stream.Recv()
			if err == io.EOF {
				// end waiting
				waitCh <- 0
				break
			}
			log.Printf("delivery position lat: %d long: %d\n", *position.Lat, *position.Long)
		}
	}()

	for i := 0; i < 100; i++ {
		if err = stream.Send(&deliveryPb.Position{
			Lat:  proto.Int32(rand.Int31()),
			Long: proto.Int32(rand.Int31()),
		}); err != nil {
			return false, err
		}
	}
	err = stream.CloseSend()
	if err != nil {
		return false, err
	}

	<-waitCh

	return true, nil
}
