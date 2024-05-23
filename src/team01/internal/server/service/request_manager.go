package service

import (
	"net/http"
	"team01/internal/server/repository"
)

type ResponseData struct {
	Code        int                 `json:"code"`
	RequestType string              `json:"request_type"`
	Error       string              `json:"error,omitempty"`
	ItemData    repository.ItemData `json:"item_data"`
}

func DoRequest(reqData RequestData, store *repository.Store) ResponseData {
	var resData ResponseData

	switch reqData.Operation {
	case "SET":
		resData.RequestType = "SET"

		if data := store.GetData(reqData.ID); data.Name == "" {
			resData.Code = http.StatusCreated
		} else {
			resData.Code = http.StatusOK
		}

		store.SetData(reqData.ID, reqData.ItemData)

	case "GET":
		resData.RequestType = "GET"

		if data := store.GetData(reqData.ID); data.Name == "" {
			resData.Code = http.StatusNotFound
			resData.Error = "Not found"
		} else {
			resData.Code = http.StatusOK
			resData.ItemData = data
		}

	case "DELETE":
		resData.RequestType = "DELETE"

		if data := store.GetData(reqData.ID); data.Name == "" {
			resData.Code = http.StatusNotFound
			resData.Error = "Not found"
		} else {
			resData.Code = http.StatusOK
			store.DeleteData(reqData.ID)
		}
	}

	return resData
}
