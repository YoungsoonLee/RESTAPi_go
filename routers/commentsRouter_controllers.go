package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:AuthController"],
		beego.ControllerComments{
			Method: "CheckDisplayName",
			Router: `/:displayname`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:AuthController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/CreateUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Social",
			Router: `/Social`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:AuthController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:ServiceController"] = append(beego.GlobalControllerRouter["github.com/YoungsoonLee/RESTAPi_go/controllers:ServiceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
