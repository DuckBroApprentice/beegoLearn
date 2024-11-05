package controllers

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() { //get
	c.Data["title"] = "hello beego"
	c.Data["num"] = 12
	c.TplName = "goods.tpl"
}

/*
RestFull設計指南主要是對API接口進行了規定，要求獲取數據使用Get，增加數據使用Post，修改數據使用Put，刪除數據使用Delete
*/

func (c *GoodsController) DoAdd() { //post請求
	c.Ctx.WriteString("執行增加操作")
}

type Product struct {
	Title   string `form:"title" xml:"title"` //用於ParseForm的標籤
	Content string `form:"content" xml:"content"`
}

func (c *GoodsController) DoEdit() { //put請求
	//title := c.GetString("title")
	p := Product{}

	if err := c.ParseForm(&p); err != nil {
		c.Ctx.WriteString("獲取數據失敗")
		return
	}
	fmt.Printf("%#v", p)
	c.Ctx.WriteString("執行修改操作")
}

func (c *GoodsController) DoDelete() { //delete
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("參數錯誤")
	}
	c.Ctx.WriteString("執行刪除操作--" + strconv.Itoa(id))
}

//接收post傳送的XML數據

func (c *GoodsController) Xml() { //delete

	p := Product{}

	// str := string(c.Ctx.Input.RequestBody)
	// beego.Info(str)
	// c.Ctx.WriteString(str)

	err := xml.Unmarshal(c.Ctx.Input.RequestBody, p) //Ctx.Input.RequestBody 獲取xml數據 Unmarshal將xml填入struct
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	}
	c.Data["json"] = p
	c.ServeJSON()
}
