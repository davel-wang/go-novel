package routers

import (
	"mybee/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index/:id", &controllers.MainController{})
}
