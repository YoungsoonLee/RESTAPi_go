// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/YoungsoonLee/RESTAPi_go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/user",
			beego.NSRouter("/confirmEmail/:confirmToken", &controllers.UserController{}, "post:ConfirmEmail"),
			beego.NSRouter("/resendConfirmEmail/:email", &controllers.UserController{}, "post:ResendConfirmEmail"),
			beego.NSRouter("/forgotPassword/:email", &controllers.UserController{}, "post:ForogtPassword"),
			beego.NSRouter("/isValidResetPasswordToken/:resetToken", &controllers.UserController{}, "post:IsValidResetPasswordToken"),
		),

		/*
			beego.NSNamespace("/confirm",
				beego.NSInclude(&controllers.ConfirmController{}),
			),
		*/
		//beego.NSRouter("/auth/checkDisplayName/:displayname", &controllers.AuthController{}, "get:CheckDisplayName"),

		beego.NSNamespace("/auth",
			beego.NSRouter("/checkDisplayName/:displayname", &controllers.AuthController{}, "get:CheckDisplayName"),
			beego.NSRouter("/register", &controllers.AuthController{}, "post:CreateUser"),
			beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
			beego.NSRouter("/checkLogin", &controllers.AuthController{}, "get:CheckLogin"),
			beego.NSRouter("/social", &controllers.AuthController{}, "post:Social"),
			beego.NSRouter("/logout", &controllers.AuthController{}, "post:Logout"),
		),

		//adimn
		beego.NSNamespace("/admin",
			beego.NSNamespace("/service",
				beego.NSInclude(&controllers.ServiceController{}),
			),
		),
	)
	beego.AddNamespace(ns)
}
