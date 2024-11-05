package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	//設置cookie，未設定過期時間(第三個參數)
	//c.Ctx.SetCookie("username", "zhangsan")

	//設置cookie-2
	// c.Ctx.SetCookie("username", "lisi", 10)

	//設置cookie訪問路徑
	// c.Ctx.SetCookie("username", "wangwu", 1000, "/article")
	//"/article"表示在article路由及其子路由可訪問

	//cookie的路徑Domain (二級域名下面也可訪問cookie)
	//itying.com   www.itying.com  a.itying.com  b.itying.com  目的：後面三個域之間共享cookie
	//c.Ctx.SetCookie("username", "wangwu1111", 1000, "/", ".itying.com")

	//設置中文cookie
	//c.Ctx.SetCookie("username", "張三", 1000) 無法獲取張三
	//c.SetSecureCookie("123456", "username", "張三", 1000) //第一個參數為密鑰(加密字符串) Get要與之相同

	c.TplName = "index.html"
}
