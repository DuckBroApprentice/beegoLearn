package controllers

import (
	"github.com/astaxie/beego"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) Get() {
	//獲取cookie
	username := c.Ctx.GetCookie("username") //與SetCookie的第一個參數一樣
	c.Data["username"] = username

	c.TplName = "product.html"
}
