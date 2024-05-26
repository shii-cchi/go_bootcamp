package service

type RequestString struct {
	DbRequest string `json:"db_request"`
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
