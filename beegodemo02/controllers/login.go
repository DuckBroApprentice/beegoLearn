package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//路由跳轉：狀態碼一定要正確
//c.Redircet("/article/xml",302)  302為狀態碼
//c.Ctx.Redirect(302, "/article/xml")
//兩者效果一樣
/*渲染
<go>
func (n *WelcomeController) Add(){ n.TplName = "Welcom.html"}
<html>
  <meta http-equiv="refresh" content ="5; url=http://www.xxxxx" /> content = "5 ...  "5秒後跳轉
*/

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) DoLogin() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Println(username, password)

	//執行跳轉

	c.Redirect("/", 302) //首頁
	//c.Redirect("/cms_123.html", 302) 跳轉到cms_123.thml
	//c.Ctx.Redirect(302, "/")  同c.Redirect只是參數要留意

}
