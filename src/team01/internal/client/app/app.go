package app

import (
	"sync"
	"team01/internal/client/config"
	"team01/internal/client/service"
)

func RunClient() {
	cfg := config.SetupFlags()
	cfg.PortChan = make(chan int, 1)

	var heartbeat service.Heartbeat

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		service.DoHeartbeat(&cfg, &heartbeat)
	}()

	service.MakeRequest(&cfg, &heartbeat)

	wg.Wait()
}
