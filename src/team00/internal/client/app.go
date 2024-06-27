package client

import (
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"team00/api/generated"
	"team00/internal/db"
)

func RunClient() {
	k := flag.Float64("k", 2.0, "Anomaly coefficient")
	flag.Parse()

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	defer conn.Close()

	cl := generated.NewTransmitterServiceClient(conn)

	database, err := db.ConnectToDb()

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile("client.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatalf("failed to open client.log file: %v", err)
	}

	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	err = DetectAnomalies(cl, *k, database, logger)

	if err != nil {
		log.Fatal(err)
	}
}
