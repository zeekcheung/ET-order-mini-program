package models

import "net/http"

type ResponseModel[D any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    D      `json:"data"`
	Error   string `json:"error"`
}

func NewRespModel(code int) *ResponseModel[any] {
	return &ResponseModel[any]{
		Code:    code,
		Message: http.StatusText(code),
	}
}

func NewSuccessModel(code int, data any) *ResponseModel[any] {
	model := NewRespModel(code)
	model.Data = data
	return model
}

func NewFailModel(code int, errMsg string) *ResponseModel[any] {
	model := NewRespModel(code)
	model.Error = errMsg
	return model
}
