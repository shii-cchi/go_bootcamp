package server

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"os"
	"sync"
	"team00/transmitter"
	"time"
)

type TransmitterServer struct {
	transmitter.UnimplementedTransmitterServiceServer
	once      sync.Once
	sessionId string
	mean      float64
	sd        float64
	logger    *log.Logger
}

func NewTransmitterServer() *TransmitterServer {
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalf("Failed to open server.log file: %v", err)
	}

	logger := log.New(file, "", log.LstdFlags)

	return &TransmitterServer{
		logger: logger,
	}
}

func (s *TransmitterServer) TransmitStream(req *empty.Empty, stream transmitter.TransmitterService_TransmitStreamServer) error {
	s.once.Do(func() {
		s.sessionId = uuid.New().String()
		s.mean = rand.Float64()*(1.5-0.3) + 0.3
		s.sd = float64(rand.Intn(20) - 10)

		s.logger.Printf("sessionId: %s, mean: %f, sd: %f", s.sessionId, s.mean, s.sd)
	})

	for {
		frequency := rand.NormFloat64()*s.sd + s.mean

		now := time.Now().UTC()

		ts := &timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		}

		transmission := &transmitter.Transmission{
			SessionId: s.sessionId,
			Frequency: frequency,
			Timestamp: ts,
		}

		err := stream.Send(transmission)

		if err != nil {
			return err
		}

		time.Sleep(100 * time.Millisecond)
	}
}
