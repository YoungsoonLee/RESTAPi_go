package models

import (
	"fmt"
)

// 消息格式
type RespCode struct {
	Code    string                 `json:"code" desc:"代码"`
	Message string                 `json:"message" desc:"描述"`
	Data    map[string]interface{} `json:"data" desc:"数据"`
}

func (rc *RespCode) Error() string {
	return fmt.Sprintf("code: %s, message: %s, data: %v", rc.Code, rc.Message, rc.Data)
}

// 错误的响应
func ErrorResponse(code, message string) *RespCode {
	return &RespCode{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
