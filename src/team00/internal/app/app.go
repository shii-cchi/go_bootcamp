package app

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"team00/internal/server"
	"team00/transmitter"
)

func RunServer() {
	s := grpc.NewServer()

	srv := server.NewTransmitterServer()

	transmitter.RegisterTransmitterServiceServer(s, srv)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer l.Close()

	log.Println("Server started. Listening on port 8080...")

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
