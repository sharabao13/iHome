package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"iHome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetArea")
	beego.Router("/api/v1.0/houses/index", &controllers.HouseIndexController{}, "get:GetHouseIndex")
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionData;delete:DeleteSessionData")
	beego.Router("/api/v1.0/users", &controllers.UserController{}, "post:Reg")
	beego.Router("/api/v1.0/sessions", &controllers.SessionController{}, "post:Login")
	//avatar api/v1.0/user/avatar
	beego.Router("/api/v1.0/user/avatar", &controllers.UserController{}, "post:Avatar")
	//api/v1.0/user
	beego.Router("/api/v1.0/user", &controllers.UserController{}, "get:GetUserData")
	//api/v1.0/user/name
	beego.Router("/api/v1.0/user/name", &controllers.UserController{}, "put:UpdateUserName")
	//api/v1.0/user/auth
	beego.Router("/api/v1.0/user/auth", &controllers.UserController{}, "get:GetUserData;post:PostRealName")

}
