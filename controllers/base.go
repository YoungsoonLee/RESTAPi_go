package controllers

import (
	"fmt"
	"regexp"

	"github.com/YoungsoonLee/RESTAPi_go/libs"
	"github.com/YoungsoonLee/RESTAPi_go/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// BaseController ...
type BaseController struct {
	beego.Controller
}

// ResponseError ...
func (b *BaseController) ResponseError(code string, err error) {
	//beego.Error(err.Error())

	response := &models.RespCode{
		Code:    code,
		Message: err.Error(),
		Data:    nil,
	}
	b.Ctx.Output.JSON(response, true, true)
	// TODO: logging
	b.StopRun()
}

// ResponseHTTPError ...
func (b *BaseController) ResponseHTTPError(status int, code string, err error) {
	b.Ctx.Output.Status = status
	b.ResponseError(code, err)
}

// ResponseCommonError ...
func (b *BaseController) ResponseCommonError(e *libs.ControllerError) {
	beego.Error(fmt.Errorf(e.Message))
	b.ResponseHTTPError(e.Status, e.Code, fmt.Errorf(e.Message))
}

// ResponseServerError ...
func (b *BaseController) ResponseServerError(e *libs.ControllerError, err error) {
	beego.Error(err)
	b.ResponseHTTPError(e.Status, e.Code, fmt.Errorf(e.Message))
}

// ValidDisplayname ...
func (b *BaseController) ValidDisplayname(displayname string) {

	if len(displayname) < 4 || len(displayname) > 16 {
		//beego.Error("key: displayname, value: ", displayname, ", message: ", libs.ErrDisplayname.Message)
		b.ResponseCommonError(libs.ErrDisplayname)
	}
}

// ValidId ...
func (b *BaseController) ValidId(id string) {

	if len(id) == 0 {
		b.ResponseCommonError(libs.ErrIdAbsent)
	}
}

// ValidEmail ...
func (b *BaseController) ValidEmail(email string) {
	valid := validation.Validation{}
	v := valid.Email(email, "Email")
	if !v.Ok {
		//loggingValidError(v)
		b.ResponseCommonError(libs.ErrEmail)
	}

	v = valid.MaxSize(email, 100, "Email")
	if !v.Ok {
		//loggingValidError(v)
		b.ResponseCommonError(libs.ErrMaxEmail)
	}
}

// ValidPassword ...
func (b *BaseController) ValidPassword(password string) {
	// 8 ~ 16 letters
	if len(password) < 8 || len(password) > 16 {
		beego.Error("key: password, value: ", password, ", message: ", libs.ErrPassword.Message)
		b.ResponseCommonError(libs.ErrPassword)
	}

	valid := validation.Validation{}
	pattern := regexp.MustCompile("") //TODO: add regex for password

	v := valid.Match(password, pattern, "password")
	if !v.Ok {
		loggingValidError(v)
		b.ResponseCommonError(libs.ErrPassword)
	}
}

func loggingValidError(v *validation.Result) {
	beego.Error("key: ", v.Error.Key, ", value: ", v.Error.Value, ", message: ", v.Error.Message)
}

// ResponseSuccess ...
func (b *BaseController) ResponseSuccess(key string, value interface{}) {
	b.Ctx.Output.Status = 200

	if key == "" {
		mresponse := &models.MrespCode{
			Code:    "ok",
			Message: "success",
			Data:    value,
		}
		b.Ctx.Output.JSON(mresponse, true, true)
		//b.StopRun()
	}

	response := &models.RespCode{
		Code:    "ok",
		Message: "success",
		Data:    map[string]interface{}{},
	}

	response.Data[key] = value
	b.Ctx.Output.JSON(response, true, true)
	//b.StopRun()

}
