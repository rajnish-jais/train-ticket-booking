package server

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"train-ticket-booking/adapters"
	"train-ticket-booking/api"
	"train-ticket-booking/config"
	pb "train-ticket-booking/proto"
)

// StartServer initializes and starts the gRPC server using config
func StartServer() {
	cfg := config.MustLoadConfig("config/config.yaml")

	listener, err := net.Listen("tcp", ":"+cfg.Server.Port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", cfg.Server.Port, err)
	}

	// Create storage adapter
	storage := adapters.NewInMemoryAdapter()
	service := api.NewTicketService(storage)

	grpcServer := grpc.NewServer()
	pb.RegisterTrainTicketServiceServer(grpcServer, service)

	fmt.Printf("🚀 gRPC Server started on port %s\n", cfg.Server.Port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
