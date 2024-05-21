package service

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
	"team01/internal/server/db"
)

type RequestData struct {
	Operation string
	ID        uuid.UUID
	Data      db.JsonData
}

func ParseRequest(reqString string) (RequestData, error) {
	var reqData RequestData

	parts := strings.SplitN(reqString, " ", 3)

	if len(parts) < 2 {
		return RequestData{}, errors.New("invalid request format")
	}

	if parts[0] != "SET" && parts[0] != "GET" && parts[0] != "DELETE" {
		return RequestData{}, errors.New("invalid operation")
	}

	if len(parts) == 3 && parts[0] != "SET" || len(parts) == 2 && parts[0] == "SET" {
		return RequestData{}, errors.New("invalid request format")
	}

	id, err := uuid.Parse(parts[1])

	if err != nil {
		return RequestData{}, errors.New("invalid UUID")
	}

	if parts[0] == "SET" {
		var data db.JsonData

		err = json.NewDecoder(strings.NewReader(parts[2])).Decode(&data)

		if err != nil {
			return RequestData{}, errors.New("invalid request format")
		}

		reqData.Data = data
	}

	reqData.Operation = parts[0]

	reqData.ID = id

	return reqData, nil
}
