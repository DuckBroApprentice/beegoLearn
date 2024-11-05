package routers

import (
	"beegodemo01/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "get:AddUser") // 第三個參數：調用方法
	beego.Router("/user/doAdd", &controllers.UserController{}, "post:DoAddUser")
	beego.Router("/user/edit", &controllers.UserController{}, "get:EditUser")
	beego.Router("/user/doEdit", &controllers.UserController{}, "post:DOEditUser")
	beego.Router("/user/getUser", &controllers.UserController{}, "get:GetUser")

	beego.Router("/goods", &controllers.GoodsController{})
	beego.Router("/goods/add", &controllers.GoodsController{}, "post:DoAdd")         //增加數據
	beego.Router("/goods/edit", &controllers.GoodsController{}, "put:DoEdit")        //修改數據
	beego.Router("/goods/delete", &controllers.GoodsController{}, "delete:DoDelete") //刪除數據
	beego.Router("/goods/xml", &controllers.GoodsController{}, "post:Xml")
}
