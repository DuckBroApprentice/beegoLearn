package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	//在控制器訪問配置參數
	//獲取app.cof裡面的信息
	//beego.AppConfig
	mysqluser := beego.AppConfig.String("mysqluser")
	adminSlice := beego.AppConfig.Strings("admin")
	beego.Info(mysqluser, adminSlice)

	//獲取beego.AppConfig.Set("redisuser", "admin666")
	redisuer := beego.AppConfig.Set("redisuser", "admin666")
	beego.Info(redisuer)

	//c.Ctx.WriteString("獲取配置信息")

	//獲取其他配置文件信息
	redispass := beego.AppConfig.String("redipass")
	beego.Info(redispass)
	//c.TplName = "index.html" //輸出rootpass root
	//嘗試不渲染,app.conf:EnableErrorsRender = false 顯示：該網頁無法正常運作
}
