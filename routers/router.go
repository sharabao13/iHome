package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"iHome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("api/v1.0/areas", &controllers.AreaController{}, "get:GetArea")
}
