package main

import (
	"github.com/YoungsoonLee/RESTAPi_go/models"
	_ "github.com/YoungsoonLee/RESTAPi_go/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	models.RegisterDB()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

		orm.Debug = true
	}

	orm.RunSyncdb("default", false, true)

	beego.Run()
}
