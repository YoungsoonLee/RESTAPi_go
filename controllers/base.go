package controllers

import (
	"fmt"

	"github.com/YoungsoonLee/RESTAPi_go/libs"
	"github.com/YoungsoonLee/RESTAPi_go/models"
	"github.com/astaxie/beego"
)

// BaseController ...
type BaseController struct {
	beego.Controller
}

// ResponseError ...
func (b *BaseController) ResponseError(code string, err error) {
	response := &models.RespCode{
		Code:    code,
		Message: err.Error(),
		Data:    nil,
	}
	b.Ctx.Output.JSON(response, true, true)
	b.StopRun()
}

// ResponseHTTPError ...
func (b *BaseController) ResponseHTTPError(status int, code string, err error) {
	b.Ctx.Output.Status = status
	b.ResponseError(code, err)
}

// ResponseCommonError ...
func (b *BaseController) ResponseCommonError(e *libs.ControllerError) {
	b.ResponseHTTPError(e.Status, e.Code, fmt.Errorf(e.Message))
}

/*
func (this *BaseController) ResponseServerError(code string, err error) {
	this.ResponseHttpError(500, code, err)
}

//
func (this *BaseController) ResponseClientError(code string, err error) {
	this.ResponseHttpError(400, code, err)
}
*/

// ResponseSuccess ...
func (b *BaseController) ResponseSuccess(key string, value interface{}) {
	b.Ctx.Output.Status = 200
	response := &models.RespCode{
		Code:    "ok",
		Message: "success",
		Data:    map[string]interface{}{},
	}
	response.Data[key] = value
	b.Ctx.Output.JSON(response, true, true)
	b.StopRun()
}
