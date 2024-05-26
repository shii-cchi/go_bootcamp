package config

import (
	"flag"
	"log"
)

type ServerConfig struct {
	CurrentPort int
	LeaderPort  int
}

func SetupFlags() ServerConfig {
	var port, leaderPort int

	flag.IntVar(&port, "P", 0, "the port of this node")
	flag.IntVar(&leaderPort, "L", 0, "the port of leader node")

	flag.Parse()

	if port == 0 || leaderPort == 0 {
		flag.Usage()
		log.Fatal("error: ports are not specified")
	}

	return ServerConfig{
		CurrentPort: port,
		LeaderPort:  leaderPort,
	}
}
