package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"team00/api/generated"
)

type TransmitterServer struct {
	generated.UnimplementedTransmitterServiceServer
	logger *log.Logger
}

func NewTransmitterServer() *TransmitterServer {
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Failed to open server.log file: %v", err)
	}

	logger := log.New(file, "", log.LstdFlags)

	return &TransmitterServer{
		logger: logger,
	}
}

func RunServer() {
	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer l.Close()

	grpcServer := grpc.NewServer()

	server := NewTransmitterServer()

	generated.RegisterTransmitterServiceServer(grpcServer, server)

	log.Println("Server is running on port 8080...")

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
