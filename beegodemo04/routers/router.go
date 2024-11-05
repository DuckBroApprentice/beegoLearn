package routers

import (
	"beegodemo04/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api", &controllers.ApiController{})
}
