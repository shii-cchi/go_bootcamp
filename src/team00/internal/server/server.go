package server

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"math/rand"
	"sync"
	"team00/api/generated"
	"time"
)

type sessionParams struct {
	sessionID string
	mean      float64
	sd        float64
}

func generateSessionParams() *sessionParams {
	return &sessionParams{
		sessionID: uuid.New().String(),
		mean:      rand.Float64()*20 - 10,
		sd:        rand.Float64()*1.2 + 0.3,
	}
}

func (s *TransmitterServer) GetFrequencyStream(req *empty.Empty, stream generated.TransmitterService_GetFrequencyStreamServer) error {
	params := generateSessionParams()
	s.logger.Printf("New session: %s, mean: %f, std: %f\n", params.sessionID, params.mean, params.sd)

	for {
		now := time.Now().UTC()

		ts := &timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		}

		var frequencyMessagePool = sync.Pool{
			New: func() interface{} {
				return &generated.FrequencyMessage{}
			},
		}

		frequencyMessage := frequencyMessagePool.Get().(*generated.FrequencyMessage)
		frequencyMessage.SessionId = params.sessionID
		frequencyMessage.Frequency = rand.NormFloat64()*params.sd + params.mean
		frequencyMessage.Timestamp = ts

		err := stream.Send(frequencyMessage)

		if err != nil {
			return err
		}

		frequencyMessagePool.Put(frequencyMessage)

		time.Sleep(100 * time.Millisecond)
	}
}
