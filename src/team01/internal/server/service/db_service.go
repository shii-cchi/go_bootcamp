package service

import (
	"net/http"
	"team01/internal/server/db"
)

type ResponseData struct {
	Code  int
	Error string
	Data  db.JsonData `json:"data,omitempty"`
}

func DoRequest(reqData RequestData, database *db.Database) ResponseData {
	var resData ResponseData

	switch reqData.Operation {
	case "SET":
		data := database.GetData(reqData.ID)

		if data.Name == "" {
			resData.Code = http.StatusCreated
		} else {
			resData.Code = http.StatusOK
		}

		database.SetData(reqData.ID, reqData.Data)
		resData.Data = reqData.Data

	case "GET":
		data := database.GetData(reqData.ID)

		if data.Name == "" {
			resData.Code = http.StatusNotFound
			resData.Error = "Not found"
		} else {
			resData.Code = http.StatusOK
			resData.Data = data
		}

	case "DELETE":
		data := database.GetData(reqData.ID)

		if data.Name == "" {
			resData.Code = http.StatusNotFound
			resData.Error = "Not found"
		} else {
			resData.Code = http.StatusOK
			resData.Data = data
		}

		database.DeleteData(reqData.ID)
	}

	return resData
}
