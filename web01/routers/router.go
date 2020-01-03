package routers

import (
	"web01/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{}, "get:GetLogin;post:HandleLogin")
	beego.Router("/login", &controllers.MainController{})
}
