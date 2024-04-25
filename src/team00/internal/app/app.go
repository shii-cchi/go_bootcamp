package app

import (
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"team00/internal/client"
	"team00/internal/db"
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

func RunClient() {
	k := flag.Float64("k", 1.0, "Anomaly coefficient")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	defer conn.Close()

	cl := transmitter.NewTransmitterServiceClient(conn)

	database, err := db.ConnectToDb()

	if err != nil {
		log.Fatal(err)
	}

	err = client.DetectAnomalies(cl, *k, database)

	if err != nil {
		log.Fatal(err)
	}

}
