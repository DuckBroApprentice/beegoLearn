package routers

import (
	"beegodemo02/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//動態路由
	//  /api/:id
	beego.Router("/api/:id", &controllers.ApiController{})

	//正則路由
	//cms_:id([0-9]+).html  動態匹配0-9一個或多個數字  id([0-9]+).html可看作一個整體
	//Ex:  8080/cms_15182.html   //15182可當成商品ID  //只能傳入數字
	beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/doLogin", &controllers.LoginController{}, "post:DoLogin")

	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/doRegister", &controllers.RegisterController{}, "post:DoRegister")

}
