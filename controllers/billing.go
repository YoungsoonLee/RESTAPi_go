package controllers

import (
	"encoding/json"

	"github.com/YoungsoonLee/RESTAPi_go/libs"
	"github.com/YoungsoonLee/RESTAPi_go/models"
)

type BillingController struct {
	BaseController
}

type PayTransaction struct {
	UID             int
	ItemID          int
	ItemName        string
	ItemDescription string
	Price           int
	PxID            string //paytransaction id
}

// GetChargeItems ...
// @Title Create Payment Category
// @Description create payment category
// @Success 200 {int} models.PaymentItem.ItemID
// @Failure 403 body is empty
// @router / [GET]
func (b *BillingController) GetChargeItems() {

	// save to db
	chargeItems, err := models.GetChargeItems()
	if err != nil {
		b.ResponseError(libs.ErrDatabase, err)
	}

	//success
	b.ResponseSuccess("", chargeItems)
}

// GetPaymentToken ...
// @Title Get PaymentToken
// @Description create payment category
// @Param	UID			json 	INT		false		"user id"
// @Param	ItemID		json 	INT		false		"item id"
// @Success 200 {int} models.PaymentItem.PayTryID
// @Failure 403 body is empty
// @router / [post]
func (b *BillingController) GetPaymentToken() {
	//
	var pt models.PaymentTry
	err := json.Unmarshal(b.Ctx.Input.RequestBody, &pt)
	if err != nil {
		b.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	// validation param

	pxid, err := models.AddPaymentTry(pt)
	if err != nil {
		b.ResponseError(libs.ErrDatabase, err)
	}

	b.ResponseSuccess("pxid", pxid)

}
