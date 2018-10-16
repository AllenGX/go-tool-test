package routers

import (
	"test/ajax/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/data", &controllers.DainController{})
}
