package handler

import (
	"time"
)

type RequestString struct {
	DbRequest string `json:"db_request"`
}

type Heartbeat struct {
	NodesList         []Node `json:"nodes_list"`
	ReplicationFactor int    `json:"replication_factor"`
}

type Node struct {
	Port       int       `json:"port"`
	Role       string    `json:"role"`
	LastActive time.Time `json:"last_active"`
}

const ReplicationFactor = 2
