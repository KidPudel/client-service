package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	clientServer "github.com/KidPudel/client-service/internal/adapters/grpc"
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

	// usecases
	clientUsecase := clientUsecases.NewClientUsecase(deliveryClient)

	server := grpc.NewServer()

	clientPb.RegisterClientServer(server, clientServer.NewClientServer(clientServer.ClientServerOptions{
		ClientUsecase: clientUsecase,
	}))

	server.Serve(listenConfig)

}
