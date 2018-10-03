package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/YoungsoonLee/RESTAPi_go/libs"

	"github.com/YoungsoonLee/RESTAPi_go/models"
)

// UserController ...
type UserController struct {
	BaseController
}

type ResetPassword struct {
	ResetToken string `json:"resetToken"`
	Password   string `json:"password"`
}

// ConfirmEmail ...
func (u *UserController) ConfirmEmail() {
	confirmToken := u.GetString(":confirmToken")
	//fmt.Println(confirmToken)

	if len(confirmToken) == 0 {
		u.ResponseCommonError(libs.ErrTokenAbsent)
	}

	// find user by email confirm token
	user, libErr := models.CheckConfirmEmailToken(confirmToken)
	if libErr == nil {
		// update
		_, err := models.ConfirmEmail(*user)
		if err != nil {
			beego.Error("email confirm update error: ", err)
			u.ResponseCommonError(libs.ErrSystem)
		}
	} else {
		if libErr.Code == "10008" {
			// alaredy confirmed
			u.ResponseSuccess("uid", strconv.FormatInt(user.Id, 10))
		} else {
			// error
			u.ResponseCommonError(libErr)
		}
	}

	// finish update confirm email.
	// havt to go to login in frontend
	u.ResponseSuccess("uid", strconv.FormatInt(user.Id, 10))
}

// ResendConfirmEmail ...
func (u *UserController) ResendConfirmEmail() {
	email := u.GetString(":email")

	// validation
	u.ValidEmail(email)

	// check email
	var user models.User
	user, err := models.FindByEmail(email)
	// if err == nil, already exists Email
	if err != nil {
		u.ResponseCommonError(libs.ErrNoUser)
	}

	// update token and send email with confirm token
	_, err = models.ResendConfirmEmail(user)
	if err != nil {
		beego.Error("email confirm update error: ", err)
		u.ResponseCommonError(libs.ErrSystem)
	}

	u.ResponseSuccess("", user)

}

// ForogtPassword ...
func (u *UserController) ForogtPassword() {
	email := u.GetString(":email")

	// validation
	u.ValidEmail(email)

	// check email
	var user models.User
	user, err := models.FindByEmail(email)
	// if err == nil, already exists Email
	if err != nil {
		u.ResponseCommonError(libs.ErrNoUser)
	}
	//fmt.Println(user)
	// send forgot password token
	_, err = models.SendPasswordResetToken(user)
	if err != nil {
		beego.Error("send password reset token error: ", err)
		u.ResponseCommonError(libs.ErrSystem)
	}

	u.ResponseSuccess("", user)
}

// IsValidResetPasswordToken ...
func (u *UserController) IsValidResetPasswordToken() {
	resetToken := u.GetString(":resetToken")
	//fmt.Println(confirmToken)

	if len(resetToken) == 0 {
		u.ResponseCommonError(libs.ErrTokenAbsent)
	}

	// find user by reset token
	user, libErr := models.CheckResetPasswordToken(resetToken)
	if libErr != nil {
		if libErr.Code == "10008" {
			// alaredy confirmed
			u.ResponseSuccess("uid", strconv.FormatInt(user.Id, 10))
		} else {
			// error
			u.ResponseCommonError(libErr)
		}
	}

	// finish update confirm email.
	// havt to go to login in frontend
	u.ResponseSuccess("uid", strconv.FormatInt(user.Id, 10))
}

// ResetPassword ...
func (u *UserController) ResetPassword() {
	var resetPassword ResetPassword
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &resetPassword)
	if err != nil {
		fmt.Println(err)
	}

	if err := models.ResetPassword(resetPassword.ResetToken, resetPassword.Password); err != nil {
		beego.Error("reset password error: ", err)
		u.ResponseCommonError(libs.ErrSystem)
	}

	u.ResponseSuccess("resetToken", resetPassword.ResetToken)
}

// GetProfile ...
func (u *UserController) GetProfile() {
	var user models.User

	uid := u.GetString(":id")

	// validation
	u.ValidId(uid)

	user, err := models.FindById(uid)
	// if err == nil, already exists displayname
	if err != nil {
		u.ResponseCommonError(libs.ErrNoUser)
	}
	u.ResponseSuccess("", user)
}

// UpdateProfile ...
func (u *UserController) UpdateProfile() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	if _, err := models.UpdateProfile(user); err != nil {
		beego.Error("update profile error: ", err)
		u.ResponseCommonError(libs.ErrSystem)
	}
	u.ResponseSuccess("", user)
}

// UpdatePassword ...
func (u *UserController) UpdatePassword() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	fmt.Println(user)

	if _, err := models.UpdatePassword(user); err != nil {
		beego.Error("update profile error: ", err)
		u.ResponseCommonError(libs.ErrSystem)
	}
	u.ResponseSuccess("", user)

}

// ---------------------------------------------------------------------------------------------------------------
// maybe not use from below
// Post ...
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {

	var user models.User

	user.Displayname = u.Input().Get("displayname")
	user.Email = u.Input().Get("email")
	user.Password = u.Input().Get("password")

	// TODO: what about the social

	// validation
	u.ValidDisplayname(user.Displayname)
	u.ValidEmail(user.Email)
	u.ValidPassword(user.Password)

	// check dup displayname
	_, err := models.FindByDisplayname(user.Displayname)
	// if err == nil, already exists displayname
	if err == nil {
		u.ResponseCommonError(libs.ErrDupDisplayname)
	}
	// check dup email
	_, err = models.FindByEmail(user.Email)
	// if err == nil, already exists Email
	if err == nil {
		u.ResponseCommonError(libs.ErrDupEmail)
	}

	// save to db
	uid, err := models.AddUser(user)
	if err != nil {
		u.ResponseServerError(libs.ErrDatabase, err)
	}

	//success
	u.ResponseSuccess("uid", strconv.FormatInt(uid, 10))
}

// @Title Login
// @Description Logs user into the system
// @Param	displayname		query 	string	true		"The displayname for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {

	displayname := u.Input().Get("displayname")
	password := u.Input().Get("password")

	// validation
	u.ValidDisplayname(displayname)
	u.ValidPassword(password)

	// Find salt, password hash for auth
	user, err := models.FindAuthByDisplayname(displayname)
	if err != nil {
		u.ResponseCommonError(libs.ErrPass)
	}

	// check password
	ok, err := user.CheckPass(password)
	if !ok || err != nil {
		// wrong password
		u.ResponseCommonError(libs.ErrPass)
	}

	// login
	et := libs.EasyToken{
		Displayname: user.Displayname,
		Uid:         user.Id,
		Expires:     time.Now().Unix() + 3600, // 1 hour
	}

	// TODO: set cookie ???

	token, err := et.GetToken()
	if token == "" || err != nil {
		u.ResponseCommonError(libs.ErrTokenOther)
	}
	//this.Data["json"]  := LoginToken{user.Displayname, user.Id, token}
	u.ResponseSuccess("login", LoginToken{user.Displayname, user.Id, token})
}

// Auth ...
// @Title Auth
// @Description validation of token
// @Success 200 {object}
// @Failure 401 unauthorized
// @router /auth [get]
func (u *UserController) Auth() {
	et := libs.EasyToken{}
	authtoken := strings.TrimSpace(u.Ctx.Request.Header.Get("Authorization"))
	valido, _, err := et.ValidateToken(authtoken)

	if !valido || err != nil {
		u.ResponseCommonError(libs.ErrExpiredToken)
	}

	u.ResponseSuccess("token", "token is valid")
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users := models.GetAllUsers()
	u.Data["json"] = users
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var user models.User
		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
