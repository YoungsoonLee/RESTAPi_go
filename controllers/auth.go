package controllers

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/YoungsoonLee/RESTAPi_go/libs"
	"github.com/YoungsoonLee/RESTAPi_go/models"
)

type AuthController struct {
	BaseController
}

type LoginToken struct {
	Displayname string `json:"user"`
	Uid         int64  `json:"uid"`
	Token       string `json:"token"`
}

// CheckDisplayName ...
// @Title CheckDisplayName
// @Description create services
// @Param	body		body 		true		""
// @Success 200 {string} displayname
// @Failure 403 body is empty
// @router /:displayname [get]
func (c *AuthController) CheckDisplayName() {

	displayname := c.GetString(":displayname")
	// validation
	c.ValidDisplayname(displayname)

	_, err := models.FindByDisplayname(displayname)
	// if err == nil, already exists displayname
	if err == nil {
		c.ResponseCommonError(libs.ErrDupDisplayname)
	}

	//success
	c.ResponseSuccess("displayname", displayname)
}

// Local ...
// @Title CreateUser
// @Description create users
// @Param	displayname		query 	string	true		"The displayname"
// @Param	email			query 	string	true		"The email"
// @Param	password		query 	string	true		"The password"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /local [post]
func (c *AuthController) Local() {

	var user models.User

	//user.Displayname = c.Input().Get("displayname")
	//user.Email = c.Input().Get("email")
	//user.Password = c.Input().Get("password")

	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	//fmt.Println(string(c.Ctx.Input.RequestBody[:]))
	//fmt.Println(user)

	// validation
	c.ValidDisplayname(user.Displayname)
	c.ValidEmail(user.Email)
	c.ValidPassword(user.Password)

	// check dup displayname
	_, err := models.FindByDisplayname(user.Displayname)
	// if err == nil, already exists displayname
	if err == nil {
		c.ResponseCommonError(libs.ErrDupDisplayname)
	}
	// check dup email
	_, err = models.FindByEmail(user.Email)
	// if err == nil, already exists Email
	if err == nil {
		c.ResponseCommonError(libs.ErrDupEmail)
	}

	// save to db
	uid, err := models.AddUser(user)
	if err != nil {
		c.ResponseServerError(libs.ErrDatabase, err)
	}

	//success
	//c.ResponseSuccess("uid", strconv.FormatInt(uid, 10))

	// auto login
	user.Id = uid
	c.makeLogin(&user)
}

// @Title Login
// @Description Logs user into the system
// @Param	displayname		query 	string	true		"The displayname for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (c *AuthController) Login() {

	//displayname := c.Input().Get("displayname")
	//password := c.Input().Get("password")

	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	// validation
	inputPass := user.Password
	c.ValidDisplayname(user.Displayname)
	c.ValidPassword(user.Password)

	//fmt.Println(user.Displayname, user.Password)

	// Find salt, password hash for auth
	user, err := models.FindAuthByDisplayname(user.Displayname)
	if err != nil {
		c.ResponseCommonError(libs.ErrPass)
	}

	// check password
	ok, err := user.CheckPass(inputPass)
	if !ok || err != nil {
		// wrong password
		c.ResponseCommonError(libs.ErrPass)
	}

	// login
	// TODO: set cookie ???
	/*
		et := libs.EasyToken{
			Displayname: user.Displayname,
			Uid:         user.Id,
			Expires:     time.Now().Unix() + 3600, // 1 hour
		}
			token, err := et.GetToken()
			if token == "" || err != nil {
				c.ResponseCommonError(libs.ErrTokenOther)
			}
			//this.Data["json"]  := LoginToken{user.Displayname, user.Id, token}
			c.ResponseSuccess("login", LoginToken{user.Displayname, user.Id, token})
	*/
	c.makeLogin(&user)
}

func (c *AuthController) makeLogin(user *models.User) {
	// login
	et := libs.EasyToken{
		Displayname: user.Displayname,
		Uid:         user.Id,
		Expires:     time.Now().Unix() + 3600, // 1 hour
	}

	// TODO: set cookie ???

	token, err := et.GetToken()
	if token == "" || err != nil {
		c.ResponseCommonError(libs.ErrTokenOther)
	}
	//this.Data["json"]  := LoginToken{user.Displayname, user.Id, token}
	c.ResponseSuccess("", LoginToken{user.Displayname, user.Id, token})
}

// CheckLogin ...
func (c *AuthController) CheckLogin() {
	//fmt.Println(string(c.Ctx.Input.RequestBody[:]))

	authtoken := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	fmt.Println(authtoken)
}
