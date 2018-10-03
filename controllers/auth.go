package controllers

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"github.com/YoungsoonLee/RESTAPi_go/libs"
	"github.com/YoungsoonLee/RESTAPi_go/models"
)

type AuthController struct {
	BaseController
}

type LoginToken struct {
	Displayname string `json:"displayname"`
	Uid         int64  `json:"uid"`
	Token       string `json:"token"`
}

type Social struct {
	Provider            string `json:"provider"`
	ProviderAccessToken string `json:"accessToken"`
	Email               string `json:"email"`
	ProviderID          string `json:"providerId"`
	Picture             string `json:"picture"`
}

type AuthedData struct {
	Uid         int64
	Displayname string
	Balance     int
	Pciture     string
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

// CreateUser ...
// @Title CreateUser except social
// @Description create users
// @Param	displayname		query 	string	true		"The displayname"
// @Param	email			query 	string	true		"The email"
// @Param	password		query 	string	true		"The password"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /CreateUser [post]
func (c *AuthController) CreateUser() {
	var user models.User
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

	if user.Provider == "facebook" && user.Password == "" {
		c.ResponseCommonError(libs.ErrLoginFacebook)
	}
	if user.Provider == "google" && user.Password == "" {
		c.ResponseCommonError(libs.ErrLoginGoogle)
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

// CheckLogin ...
func (c *AuthController) CheckLogin() {

	et := libs.EasyToken{}
	authtoken := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	valido, displayname, err := et.ValidateToken(authtoken)

	beego.Info("Check Login: ", displayname, valido)

	if !valido || err != nil {
		c.ResponseCommonError(libs.ErrExpiredToken)
	}

	// get userinfo
	//fmt.Println("check login: ", displayname)
	var user models.User
	user, err = models.FindByDisplayname(displayname)
	if err != nil {
		c.ResponseCommonError(libs.ErrNoUser)
	}

	//fmt.Println(user.Displayname, user.Picture)

	c.ResponseSuccess("", AuthedData{user.Id, user.Displayname, 0, user.Picture})
}

// Social ...
// @Title CreateUser or SigninUser for social FB and G+
// @Description create social users or signin
// @Param	provider		query 	string	true		"The provider (FB, G+)"
// @Param	accessToken		query 	string	true		"The accessToken"
// @Success 200 {int}
// @Failure 403 body is empty
// @router /Social [post]
func (c *AuthController) Social() {
	var social Social
	json.Unmarshal(c.Ctx.Input.RequestBody, &social)

	// TODO: validation
	// unless provier is null or accessToken is null, get error
	//fmt.Println(social)

	var user models.User
	user, err := models.FindByEmail(social.Email)

	// if err == nil, already exists Email
	if err == nil {
		// make login
		//fmt.Println("already exists email ", user)
		//update social info, it can login local and social both.
		if len(user.Provider) == 0 || user.Provider != social.Provider {
			user.Provider = social.Provider
			user.ProviderAccessToken = social.ProviderAccessToken
			user.ProviderID = social.ProviderID
			user.Picture = social.Picture

			c.updateSocialInfo(user)
		}

		c.makeLogin(&user)

	} else {
		// add social user
		user.Provider = social.Provider
		user.ProviderAccessToken = social.ProviderAccessToken
		user.ProviderID = social.ProviderID
		user.Email = social.Email
		user.Picture = social.Picture
		c.createSocialUser(user)
	}

}

// @Title Logout
// @Description Logs user into the system
// @Param	displayname		query 	string	true		"The displayname for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (c *AuthController) Logout() {

}

func (c *AuthController) createSocialUser(user models.User) {

	uid, displayname, err := models.AddSocialUser(user)
	if err != nil {
		c.ResponseServerError(libs.ErrDatabase, err)
	}

	user.Id = uid
	user.Displayname = displayname
	c.makeLogin(&user)
}

func (c *AuthController) updateSocialInfo(user models.User) {
	uid, displayname, err := models.UpdateSocialInfo(user)
	if err != nil {
		c.ResponseServerError(libs.ErrDatabase, err)
	}

	user.Id = uid
	user.Displayname = displayname
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

	// TODO: add balance to LoginToken
	c.ResponseSuccess("", LoginToken{user.Displayname, user.Id, token})
}
