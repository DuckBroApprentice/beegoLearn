package routers

import (
	"beegocookie/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/product", &controllers.ProductController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/doLogin", &controllers.LoginController{}, "post:DoLogin")
	beego.Router("/loginOut", &controllers.LoginController{}, "get:LoginOut")
}
