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
// @Title Create Payment Gateway
// @Description create payment gateway
// @Param	pg_description	json 	string	false		"pg description"
// @Success 200 {int} models.PaymentGateway.PgID
// @Failure 403 body is empty
// @router / [post]
func (p *PaymentGatewayController) Post() {

	var pg models.PaymentGateway
	err := json.Unmarshal(p.Ctx.Input.RequestBody, &pg)
	if err != nil {
		p.ResponseServerError(libs.ErrJSONUnmarshal, err)
	}

	// TODO: validation

	// save to db
	pgid, err := models.AddPaymentGateway(pg)
	if err != nil {
		p.ResponseServerError(libs.ErrDatabase, err)
	}

	//success
	p.ResponseSuccess("pgid", pgid)
}
