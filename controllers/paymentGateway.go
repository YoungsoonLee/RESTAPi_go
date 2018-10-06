package controllers

import (
	"encoding/json"

	"github.com/YoungsoonLee/RESTAPi_go/libs"
	"github.com/YoungsoonLee/RESTAPi_go/models"
)

type PaymentGatewayController struct {
	BaseController
}

// Post ...
// @Title Create PG
// @Description create services
// @Param	body		body 	models.Service	true		"body for service content"
// @Success 200 {int} models.Service.Id
// @Failure 403 body is empty
// @router / [post]
func (p *PaymentGatewayController) Post() {

	var pg models.PaymentGateway
	//service.Description = s.Input().Get("description")

	json.Unmarshal(p.Ctx.Input.RequestBody, &pg)

	//fmt.Println(string(s.Ctx.Input.RequestBody[:]))
	//fmt.Println(service)

	// save to db
	pgid, err := models.AddPaymentGateway(pg)
	if err != nil {
		p.ResponseServerError(libs.ErrDatabase, err)
	}

	//success
	p.ResponseSuccess("pgid", pgid)
}
