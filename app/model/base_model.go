package model

import "github.com/geekappio/itonchain/app/enum"

type BaseRequest struct {
}

type BaseResponse struct {
}

type PageRequest struct {
	BaseRequest
	PageNum  int `json:"pageNum" binding:"required"`
	PageSize int `json:"pageSize"`
}

type ResponseModel struct {
	ReturnCode string      `json:"returnCode" binding:"required"`
	ReturnMsg  string      `json:"returnMsg"`
	ReturnData interface{} `json:"returnData"`
}

// Create response model instance with SUCCESS result.
func NewSuccessResponseModel() *ResponseModel {
	return &ResponseModel{
		ReturnCode: enum.SYSTEM_SUCCESS.GetRespCode(),
	}
}

// Create response model instance with SUCCESS result and data.
func NewSuccessResponseModelWithData(returnData interface{}) *ResponseModel {
	response := &ResponseModel{
		ReturnCode: enum.SYSTEM_SUCCESS.GetRespCode(),
	}

	if returnData != nil {
		response.ReturnData = returnData
	}

	return response;
}

// Create response model instance with FAILURE or ERROR result.
func NewFailedResponseModel( errorCode enum.ErrorCode,  errorMsg string) *ResponseModel {
	return &ResponseModel{
		ReturnCode: errorCode.GetRespCode(),
		ReturnMsg:  errorMsg,
	}
}

// Create response model instance with FAILURE or ERROR result.
func NewFailedResponseModelWithData( errorCode enum.ErrorCode,  errorMsg string, returnData interface{}) *ResponseModel {
	response :=&ResponseModel{
		ReturnCode: errorCode.GetRespCode(),
		ReturnMsg:  errorMsg,
	}

	if returnData != nil {
		response.ReturnData = returnData
	}

	return response
}

