package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"team01/internal/server/repository"
)

type RequestData struct {
	Operation string
	ID        uuid.UUID
	ItemData  repository.ItemData
}

func ParseRequest(reqString string) (RequestData, error) {
	var reqData RequestData

	parts := strings.SplitN(reqString, " ", 3)

	if len(parts) < 2 {
		return RequestData{}, errors.New("invalid request format: must contain at least operation and ID")
	}

	switch parts[0] {
	case "SET", "GET", "DELETE":
		reqData.Operation = parts[0]
	default:
		return RequestData{}, fmt.Errorf("invalid operation: must be SET, GET, or DELETE, got %s", reqData.Operation)
	}

	if reqData.Operation != "SET" {
		if len(parts) != 2 {
			return RequestData{}, errors.New("invalid request format: GET or DELETE operations requires just ID")
		}
	}

	id, err := uuid.Parse(parts[1])
	if err != nil {
		return RequestData{}, errors.New("invalid UUID")
	}
	reqData.ID = id

	if reqData.Operation == "SET" {
		if len(parts) != 3 {
			return RequestData{}, errors.New("invalid request format: SET operation requires data")
		}

		err := json.NewDecoder(strings.NewReader(strings.Trim(parts[2], "'"))).Decode(&reqData.ItemData)

		if err != nil {
			return RequestData{}, errors.New("invalid request format")
		}
	}

	return reqData, nil
}
