package routers

import (
	"go-discover/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/registerApp",&controllers.MainController{},"get:RegisterApp")
	beego.Router("/api/listApp",&controllers.MainController{},"get:ListApp")
}
