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
	"team00/api/generated"
	"team00/internal/db"
	"time"
)

type Statistics struct {
	count    int
	sum      float64
	sumSq    float64
	mean     float64
	variance float64
	sd       float64
}

func DetectAnomalies(cl generated.TransmitterServiceClient, k float64, database *gorm.DB, logger *log.Logger) error {
	stream, err := cl.GetFrequencyStream(context.Background(), &empty.Empty{})

	if err != nil {
		return fmt.Errorf("error calling GetFrequencyStream: %v", err)
	}

	var frequencyMessagePool = sync.Pool{
		New: func() interface{} {
			return &generated.FrequencyMessage{}
		},
	}

	var stats Statistics
	var leftBound, rightBound float64
	const countStatisticsPoints = 150

	for {
		frequencyMessage := frequencyMessagePool.Get().(*generated.FrequencyMessage)

		err := stream.RecvMsg(frequencyMessage)

		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("error receiving data: %v", err)
		}

		if stats.count < countStatisticsPoints {
			stats = makeStatistics(frequencyMessage.Frequency, stats)
			logger.Printf("Count: %d, Mean: %f, StdDev: %f", stats.count, stats.mean, stats.sd)

			if stats.count == countStatisticsPoints {
				leftBound = stats.mean - k*stats.sd
				rightBound = stats.mean + k*stats.sd
			}

		} else {
			if isAnomaly(frequencyMessage.Frequency, leftBound, rightBound) {
				database.Create(&db.Record{SessionId: frequencyMessage.SessionId, Frequency: frequencyMessage.Frequency, Timestamp: time.Unix(frequencyMessage.Timestamp.Seconds, int64(frequencyMessage.Timestamp.Nanos)).UTC()})
				logger.Printf("An anomaly has been detected! Frequency: %f", frequencyMessage.Frequency)
			}
		}

		frequencyMessagePool.Put(frequencyMessage)
	}

	return nil
}

func makeStatistics(newFrequency float64, stats Statistics) Statistics {
	stats.count++

	stats.sum += newFrequency
	stats.sumSq += math.Pow(newFrequency, 2)

	stats.mean = stats.sum / float64(stats.count)
	stats.variance = (stats.sumSq / float64(stats.count)) - math.Pow(stats.mean, 2)
	stats.sd = math.Sqrt(stats.variance)

	return stats
}

func isAnomaly(frequency, leftBound, rightBound float64) bool {
	return frequency < leftBound || frequency > rightBound
}
