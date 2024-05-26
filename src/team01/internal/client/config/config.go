package config

import (
	"flag"
	"log"
)

type ClientConfig struct {
	Host     string
	Port     int
	PortChan chan int
}

func SetupFlags() ClientConfig {
	var host string
	var port int

	flag.StringVar(&host, "H", "", "the host address")
	flag.IntVar(&port, "P", 0, "the port of leader node")

	flag.Parse()

	if host == "" || port == 0 {
		flag.Usage()
		log.Fatal("error: host and/or port are not specified")
	}

	if host != "127.0.0.1" {
		log.Fatal("wrong host address")
	}

	return ClientConfig{
		Host: host,
		Port: port,
	}
}
