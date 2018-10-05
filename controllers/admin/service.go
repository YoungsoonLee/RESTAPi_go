package controllers

import (
	"github.com/YoungsoonLee/RESTAPi_go/libs"
	"github.com/YoungsoonLee/RESTAPi_go/models"
)

type ServiceController struct {
	BaseController
}

// Post ...
// @Title CreateService
// @Description create services
// @Param	body		body 	models.Service	true		"body for service content"
// @Success 200 {int} models.Service.Id
// @Failure 403 body is empty
// @router / [post]
func (s *ServiceController) Post() {

	var service models.Service

	service.Description = s.Input().Get("description")

	// save to db
	sid, err := models.AddService(service)
	if err != nil {
		s.ResponseServerError(libs.ErrDatabase, err)
	}

	//success
	s.ResponseSuccess("sid", sid)
}
