package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	clientServer "github.com/KidPudel/client-service/internal/adapters/grpc"
	eventListener "github.com/KidPudel/client-service/internal/adapters/kafka"
	"github.com/KidPudel/client-service/internal/infrastructure/kafka"
	clientUsecases "github.com/KidPudel/client-service/internal/usecases/client"
	clientPb "github.com/KidPudel/client-service/proto/client"
	deliveryPb "github.com/KidPudel/client-service/proto/delivery"
)

func main() {
	listenConfig, err := net.Listen("tcp", "localhost:50053")
	if err != nil {
		log.Fatal(err)
	}

	clientConn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	deliveryClient := deliveryPb.NewDeliveryClient(clientConn)

	ctx := context.Background()

	// kafka
	kafkaClient := kafka.NewKafkaClient()
	// event listener
	deliveryListener := eventListener.NewDeliveryListener(kafkaClient)
	go deliveryListener.ListenOnDeliveries(ctx)

	// usecases
	clientUsecase := clientUsecases.NewClientUsecase(deliveryClient)

	server := grpc.NewServer()

	clientPb.RegisterClientServer(server, clientServer.NewClientServer(clientServer.ClientServerOptions{
		ClientUsecase: clientUsecase,
	}))

	server.Serve(listenConfig)

}
