package src

import "net/http"

type Response struct {
	Status   string      `json:"status"`
	Data     interface{} `json:"data,omitempty"`
	Error    *ErrorInfo  `json:"error,omitempty"`
	MetaData interface{} `json:"meta_data,omitempty"`
}

type ErrorInfo struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Status: "success",
		Data: data,
	}
}

func NewInternalServerException(err error) *Response {
	return &Response{
		Status: "error",
		Error: &ErrorInfo{
			Code: http.StatusInternalServerError,
			Description: err.Error(),
		},
	}
}