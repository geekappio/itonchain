package api

type BaseRequest struct {
}

type BaseResponse struct {
}

type ResponseHead struct {
	ReturnCode 		string			`json:"returnCode"`
	ReturnMsg 		string			`json:"returnMsg"`
	ReturnData		interface{}	`json:"returnData"`
}
