package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	//獲取cookie

	username, _ := c.Ctx.GetSecureCookie("123456", "username")
	c.Data["username"] = username
	c.TplName = "login.html"
}

func (c *LoginController) DoLogin() {

	username := c.GetString("usernmae")
	//把username寫入到cookie
	//fmt.Println(username)
	c.Ctx.SetSecureCookie("123456", "username", username)
	c.Redirect("/", 302) //執行跳轉
}

func (c *LoginController) LoginOut() {
	//刪除cookie  beego未內建，方法是將過期時間設為0
	c.Ctx.SetSecureCookie("123456", "username", "", 0)
	c.Redirect("/", 302) //執行跳轉

}
