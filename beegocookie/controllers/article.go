package controllers

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {

	//獲取普通cookie
	// 	username := c.Ctx.GetCookie("username") //與SetCookie的第一個參數一樣
	// 	c.Data["username"] = username
	// 	c.TplName = "article.html"
	//獲取加密cookie
	username, _ := c.Ctx.GetSecureCookie("123456", "username")
	c.Data["username"] = username
	c.TplName = "article.html"
}
