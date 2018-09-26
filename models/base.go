package models

import (
	"fmt"
)

//
type RespCode struct {
	Code    string                 `json:"code" desc:"代码"`
	Message string                 `json:"message" desc:"描述"`
	Data    map[string]interface{} `json:"data" desc:"数据"`
}

func (rc *RespCode) Error() string {
	return fmt.Sprintf("code: %s, message: %s, data: %v", rc.Code, rc.Message, rc.Data)
}

//
func ErrorResponse(code, message string) *RespCode {
	return &RespCode{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
