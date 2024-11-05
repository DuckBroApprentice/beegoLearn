package controllers

import (
	"github.com/astaxie/beego"
)

type CmsController struct {
	beego.Controller
}

// 8080/api/"id"
func (c *CmsController) Get() {

	cmsID := c.Ctx.Input.Param(":id") ///cms_:id
	c.Ctx.WriteString("CMS詳情---" + cmsID)

}

//路由跳轉：狀態碼一定要正確
//c.Redircet("/article/xml",302)  302為狀態碼
//c.Ctx.Redirect(302, "/article/xml")
//兩者效果一樣

/*渲染
<go>
func (n *WelcomeController) Add(){ n.TplName = "Welcom.html"}
<html>
  <meta http-equiv="refresh" content ="5; url=http://www.xxxxx" />
*/
