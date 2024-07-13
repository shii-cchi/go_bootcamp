package service

import (
	"sync"
	"time"
)

const (
	heartbeatTick    = 1 * time.Second
	maxRetryAttempts = 10
	retryDelay       = 2 * time.Second
)

type RequestString struct {
	DbRequest string `json:"db_request"`
}

type FailedRequest struct {
	RequestString []byte
}

type FailedRequests struct {
	failedRequests []FailedRequest
	mutex          sync.Mutex
}

type ResponseData struct {
	Code        int      `json:"code"`
	RequestType string   `json:"request_type"`
	Error       string   `json:"error,omitempty"`
	ItemData    ItemData `json:"item_data"`
}

type ItemData struct {
	Name string `json:"name,omitempty"`
}

type Heartbeat struct {
	NodesList         []NodeSummary `json:"nodes_list"`
	ReplicationFactor int           `json:"replication_factor"`
}

type NodeSummary struct {
	Port int    `json:"port"`
	Role string `json:"role"`
}
