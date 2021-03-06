package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/YoungsoonLee/RESTAPi_go/libs"
	"github.com/YoungsoonLee/RESTAPi_go/models"
)

type BillingController struct {
	BaseController
}

// Xsolla struct
type XSuser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Ip    string `json:"ip"`
}

type XSpurchaseDetail struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

type XSpurchase struct {
	Total XSpurchaseDetail
}

type XStransaction struct {
	ID          string    `json:"id"`           // TxID from Xsolla
	ExternalID  string    `json:"external_id"`  // PxID
	PaymentDate time.Time `json:"payment_date"` // transaction_at
}

type XSollaData struct {
	Signature        string        `json:"signature"`
	NotificationType string        `json:"notification_type"`
	User             XSuser        `json:"user"`
	Purchase         XSpurchase    `json:"purchase"`
	Transaction      XStransaction `json:"transaction"`
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

	pt, err = models.AddPaymentTry(pt)
	if err != nil {
		b.ResponseError(libs.ErrDatabase, err)
	}

	fmt.Println(pt)

	b.ResponseSuccess("", pt)

}

// CallbackXsolla ...
// @Title Get xsolla callback data
// @Description ...
// ...
func (b *BillingController) CallbackXsolla() {
	var xsollaData XSollaData
	err := json.Unmarshal(b.Ctx.Input.RequestBody, &xsollaData)
	if err != nil {
		b.ResponseError(libs.ErrJSONUnmarshal, err)
	}

	fmt.Println("xsollaData: ", xsollaData)
	fmt.Println("before unmarshall: ", string(b.Ctx.Input.RequestBody[:]))
}
