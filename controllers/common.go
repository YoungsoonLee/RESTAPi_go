package controllers

import "github.com/astaxie/beego"

// Predefined const error strings.
const (
	ErrInputData    = "Data input error"
	ErrDatabase     = "Database operation error"
	ErrDupUser      = "User information already exists"
	ErrNoUser       = "User information does not exist"
	ErrPass         = "Incorrect password"
	ErrNoUserPass   = "User information does not exist or the password is incorrect"
	ErrNoUserChange = "User information does not exist or data has not changed"
	ErrInvalidUser  = "User information is incorrect"
	ErrOpenFile     = "Error opening file"
	ErrWriteFile    = "Error writing a file"
	ErrSystem       = "Operating system error"
)

// ControllerError is controller error info structer.
type ControllerError struct {
	Status   int    `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	DevInfo  string `json:"dev_info"`
	MoreInfo string `json:"more_info"`
}

// Predefined controller error values.
var (
	err404          = &ControllerError{404, 404, "page not found", "page not found", ""}
	errInputData    = &ControllerError{400, 10001, "Data input error", "Client parameter error", ""}
	errDatabase     = &ControllerError{500, 10002, "Database operation error", "Database operation error", ""}
	errDupUser      = &ControllerError{400, 10003, "User information already exists", "Duplicate database records", ""}
	errNoUser       = &ControllerError{400, 10004, "User information does not exist", "Database record does not exist", ""}
	errPass         = &ControllerError{400, 10005, "User information does not exist or the password is incorrect", "Incorrect password", ""}
	errNoUserPass   = &ControllerError{400, 10006, "User information does not exist or the password is incorrect", "Database record does not exist or password is incorrect", ""}
	errNoUserChange = &ControllerError{400, 10007, "User information does not exist or data has not changed", "Database record does not exist or data has not changed", ""}
	errInvalidUser  = &ControllerError{400, 10008, "User information is incorrect", "User information is incorrect", ""}
	errOpenFile     = &ControllerError{500, 10009, "Error opening file", "Error opening file", ""}
	errWriteFile    = &ControllerError{500, 10010, "Error writing a file", "Error writing a file", ""}
	errSystem       = &ControllerError{500, 10011, "Operating system error", "Operating system error", ""}
	errExpired      = &ControllerError{400, 10012, "Login has expired", "The token expires", ""}
	errPermission   = &ControllerError{400, 10013, "Permission denied", "Permission denied", ""}
)

// BaseController definiton.
type BaseController struct {
	beego.Controller
}

// RetError return error information in JSON.
func (base *BaseController) RetError(e *ControllerError) {
	/*
		if mode := beego.AppConfig.String("runmode"); mode == "prod" {
			e.DevInfo = ""
		}
	*/

	base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	base.Ctx.ResponseWriter.WriteHeader(e.Status)
	base.Data["json"] = e
	base.ServeJSON()

	base.StopRun()
}
