package client

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"gorm.io/gorm"
	"io"
	"log"
	"math"
	"sync"
	"team00/internal/db"
	"team00/transmitter"
)

type Statistics struct {
	count    int
	sum      float64
	sumSq    float64
	mean     float64
	variance float64
	sd       float64
}

func (s *Statistics) Update(newValue float64) {
	s.count++
	s.sum += newValue
	s.sumSq += math.Pow(newValue, 2)

	s.mean = s.sum / float64(s.count)
	s.variance = (s.sumSq / float64(s.count)) - math.Pow(s.mean, 2)
	s.sd = math.Sqrt(s.variance)
}

var transmissionPool = sync.Pool{
	New: func() interface{} {
		return &transmitter.Transmission{}
	},
}

func DetectAnomalies(cl transmitter.TransmitterServiceClient, k float64, database *gorm.DB) error {
	stream, err := cl.TransmitStream(context.Background(), &empty.Empty{})

	if err != nil {
		return fmt.Errorf("error calling TransmitStream: %v", err)
	}

	stats, err := calcStatistics(stream)

	if err != nil {
		return err
	}

	for {
		transmission := transmissionPool.Get().(*transmitter.Transmission)

		err := stream.RecvMsg(transmission)

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("error receiving data: %v", err)
		}

		leftBound := stats.mean - k*stats.sd
		rightBound := stats.mean + k*stats.sd

		if transmission.Frequency < leftBound || transmission.Frequency > rightBound {
			database.Create(&db.Record{SessionId: transmission.SessionId, Frequency: transmission.Frequency, Timestamp: transmission.Timestamp.Seconds})
			log.Printf("An anomaly has been detected! Frequency: %f", transmission.Frequency)
		}

		transmissionPool.Put(transmission)
	}

	return nil
}

func calcStatistics(stream transmitter.TransmitterService_TransmitStreamClient) (Statistics, error) {
	var stats Statistics

	for {
		if stats.count > 150 {
			break
		}

		transmission := transmissionPool.Get().(*transmitter.Transmission)

		err := stream.RecvMsg(transmission)

		if err == io.EOF {
			break
		}

		if err != nil {
			return Statistics{}, fmt.Errorf("error receiving data: %v", err)
		}

		stats.Update(transmission.Frequency)

		transmissionPool.Put(transmission)

		log.Printf("Count: %d, Mean: %f, StdDev: %f", stats.count, stats.mean, stats.sd)
	}

	return stats, nil
}
