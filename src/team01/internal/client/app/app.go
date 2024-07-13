package app

import (
	"sync"
	"team01/internal/client/config"
	"team01/internal/client/service"
)

func RunClient() {
	cfg := config.SetupFlags()

	var heartbeat service.Heartbeat
	var failedRequests service.FailedRequests

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		service.DoHeartbeat(&cfg, &heartbeat, &failedRequests)
	}()

	service.MakeRequest(&cfg, &heartbeat, &failedRequests)

	wg.Wait()
}
