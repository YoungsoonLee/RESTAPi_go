package controllers

import (
	"fmt"
	"strconv"

	"github.com/YoungsoonLee/RESTAPi_go/libs"

	"github.com/YoungsoonLee/RESTAPi_go/models"
)

type ConfirmController struct {
	BaseController
}

// @Title Get
// @Description get user by token
// @Param	token		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :token is empty
// @router /:token [get]
func (c *ConfirmController) Get() {
	confirmToken := c.GetString(":token")

	var user *models.User
	var libErr *libs.ControllerError

	if confirmToken != "" {
		user, libErr = models.CheckEmailConfirmToken(confirmToken)
		if libErr == nil {
			// update
			_, err := models.ConfirmEmail(*user)
			fmt.Println("oops: ", user.Id, err)
			//
		} else {
			// error
			c.ResponseCommonError(libErr)
		}

		/*
			user, err := models.ConfirmEmail(confirmToken)
			if err != nil {
				c.Data["json"] = err.Error()
			} else {
				c.Data["json"] = user
			}
		*/
	}

	//success
	c.ResponseSuccess("uid", strconv.FormatInt(user.Id, 10))

}
