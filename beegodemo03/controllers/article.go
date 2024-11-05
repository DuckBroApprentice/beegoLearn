package controllers

import (
	"beegodemo03/models"
	"fmt"
	"time"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.Data["title"] = "你好beego"
	//header.html <h1>這是一個頭部---{{.title}}</h1> 無法透過template輸出{{.title}}
	//顯示：這是一個頭部---
	//{{template "/public/header.html" .}}  引入外部模板的正確寫法
	//顯示：這是一個頭部---"你好beego"

	now := time.Now() //time.Time型別
	c.Data["now"] = now
	//<html> {{date .now "Y-m-d H:i:s"}} 2024-11-02 10:42:33

	c.Data["title"] = "這是文章列表"
	//<html> {{substr .title 0 4}} 這是文章

	c.Data["html"] = "<h2>這是一個後台渲染的 h2</h2>"
	//{{.html}} <h2>這是一個後台渲染的 h2</h2>
	/*兩種寫法
	<p>{{.html | html2str}}</p> 這是一個後台渲染的 h2
	<p>{{html2str  .html}}</p> 這是一個後台渲染的 h2
	*/
	/*引入css
	p>{{.html | str2html}}</p>  這是一個後台渲染的 h2 (字變大)
	<p>{{str2html  .html}}</p>  這是一個後台渲染的 h2 (字變大)
		index.css
			h2{
	    	color :red
			}
		views/article.html
		link rel="stylesheet" href="/static/css/index.css"
		或
		{{assets_css "/static/css/index.css"}}
		調用index.css的設置
		讓字變紅色
	*/
	/*
		{{htmlquote .html}}
		 &lt;h2&gt;這是一個後台渲染的&nbsp;h2&lt;/h2&gt;
	*/

	//引入js
	/*
		static/js/base.js
		<script src="/static/js/base.js"></script>
		或
		{{assets_js "/static/js/base.js"}}
		彈出一個對話框: 111

	*/
	/*config
	獲取config的值，使用方法 {{config configType configKey defaultValue}}可選
	{{config "String" 'httpport" ""}}  首字母大寫要注意
	{{config "String" "httpport" ""}} 8080
	{{config "String" "appname" ""}} beegodemo03
	*/
	userinfo := make(map[string]interface{})
	userinfo["username"] = "qek s"
	userinfo["age"] = 20
	userinfo["a"] = map[string]float64{
		"c": 4,
	}
	c.Data["userinfo"] = userinfo
	//{{map_get .userinfo "username"}} "qek s"
	//{{map_get .userinfo "age"}} 20
	//{{map_get .userinfo "a" "c"}} 4

	c.Data["unix"] = 1587880013 //渲染unix
	//{{.unix | unixToDate}} 時間戳轉成時分秒
	//2024-11-02 11:22:55
	//main.go:beego.AddFuncMap("unixToDate", models.UnixToDate)

	//字符串截取
	var str = "你好這是文章列表"
	slice := []rune(str)
	fmt.Println(string(slice[4:])) //文章列表

	c.Data["str"] = string(slice[4:])
	//{{.str}}  文章列表

	c.TplName = "article.html"
}

func (c *ArticleController) Add() {
	unix := 1589898013
	str := models.UnixToDate(unix)

	c.Ctx.WriteString("新聞add--" + str)
}
