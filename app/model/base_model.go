package model

type BaseRequest struct {
}

type BaseResponse struct {
}

type ResponseModel struct {
	ReturnCode string      `json:"returnCode" binding:"required"`
	ReturnMsg  string      `json:"returnMsg"`
	ReturnData interface{} `json:"returnData"`
}
