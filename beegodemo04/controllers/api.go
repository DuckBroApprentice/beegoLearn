package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {

	//獲取beego.AppConfig.Set("redisuser", "admin666")
	redisuer := beego.AppConfig.Set("redisuser", "admin666")
	beego.Info(redisuer)

	c.Ctx.WriteString("api接口")

}
