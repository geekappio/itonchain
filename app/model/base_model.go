package model

type BaseRequest struct {
}

type BaseResponse struct {
}

type PageRequest struct {
	BaseRequest
	PageNum int `json:"pageNum" binding:"required"`
	PageSize int `json:"pageSize"`
}

type ResponseModel struct {
	ReturnCode string      `json:"returnCode" binding:"required"`
	ReturnMsg  string      `json:"returnMsg"`
	ReturnData interface{} `json:"returnData"`
}
