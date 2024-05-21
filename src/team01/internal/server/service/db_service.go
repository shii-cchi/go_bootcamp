package service

type ResponseData struct {
	Code  int
	Error string
}

func DoRequest(reqData RequestData) ResponseData {
	var resData ResponseData

	switch reqData.Operation {
	case "SET":
		resData.Code, resData.Error = setData(reqData.ID, reqData.Data)
	case "GET":
		resData.Code, resData.Error = getData(reqData.ID)
	case "DELETE":
		resData.Code, resData.Error = deleteData(reqData.ID)
	}

	return resData
}
