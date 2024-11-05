package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

// 8080/api/"id"
func (c *ApiController) Get() {
	//獲取動態路由的值
	id := c.Ctx.Input.Param((":id")) //Param()內要與路由(/api/:id)設置相同
	c.Ctx.WriteString("api接口---" + id)
}
