package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type Article struct {
	Title   string
	Content string
}

func (c *MainController) Get() {
	//1.模板中綁定基本數據 字符串 數值 布爾值
	c.Data["website"] = "beego教程" //<html> {{.Website}}
	c.Data["title"] = "你好beego"   //<html> {{.Email}}
	c.Data["num"] = 123
	c.Data["flag"] = true

	//2.模板中綁定結構體數據
	article := Article{
		Title:   "I'm golang",
		Content: "beego",
	}
	c.Data["article"] = article //<html {{.article}} 是Data["article"]

	//3.模板中自定義變量
	/*
		<html>
			<h2>3.模板中自定義變量</h2>
			{{$title := .title}}  //自定義
			//將.title值賦值給$title 或$acb
			{{$title}} 或 {{$acb}}  //輸出
	*/

	//3.模板中循環遍歷range (切片)
	c.Data["sliceList"] = []string{"php", "javj", "golang"}
	//<html>
	//{{.sliceList[0]}}錯誤語法

	/*
		<html>
			<ul>
			{{range $key,$val := .sliceList}} //賦值給key val變數
			<li>{{$key}}--{{$val}}</li>  //輸出
			{{end}}  //就一定要的結尾
			</ul>
	*/
	//4.模板中循環遍歷range (Map)
	userinfo := make(map[string]interface{})
	userinfo["usernmae"] = "張三"
	userinfo["age"] = 20
	userinfo["sex"] = "man"
	c.Data["userinfo"] = userinfo
	// {{.userinfo["username"]}} 同樣是錯誤的寫法

	//5.模板中循環遍歷range ([]sturct)
	c.Data["articleList"] = []Article{
		{
			Title:   "新聞111",
			Content: "新聞內容111",
		},
		{
			Title:   "新聞222",
			Content: "新聞內容222",
		},
		{
			Title:   "新聞333",
			Content: "新聞內容3334",
		},
		{
			Title:   "新聞444",
			Content: "新聞內容444",
		},
	}
	/*
		<html>
				<h2>6.循環遍歷range ([]struct)</h2>
			<ul>
			  {{range $key,$val := .articleList}}
			  <li>{{$key}}--{{$val.Title}}--{{$val.Contnet}}</li>  //賦值後才可以用dot運算子
			  {{end}}
			</ul>
	*/

	//6.模板中循環遍歷range ([]sturct) 二
	c.Data["cmsList"] = []struct { //匿名結構體
		Title string
	}{
		{
			Title: "news111",
		}, {
			Title: "news222",
		}, {
			Title: "news333",
		}, {
			Title: "news444",
		},
	}

	//7.模板中的條件判斷
	c.Data["isLogin"] = true //isLogin等於true
	/*<html>
	{{if .isLogin}}
	<p>isLogin等於true</p>
	{{end}}
	*/
	c.Data["isHome"] = false //isHome等於false
	/*<html>
	{{if .isHome}}
	<p>isHome等於true</p>
	{{else}}
	<p>isHome等於false</p>
	{{end}}
	*/
	c.Data["isAbout"] = true //isAbout等於true 111
	/*<html>
	{{if .isHome}}  //false
	<p>isHome等於true</p>
	{{else}}
	{{if .isAbout}} //嵌套
	<p>isAbout等於true 111</p>
	{{end}}
	{{end}}
	*/

	//8.if條件判斷 eq/ne/lt/le/gt/ge
	/*配合if中使用
	eq : arg1 == arg2
	ne : arg1 != arg2
	lt : arg1 < arg2
	le : arg1 <= arg2
	gt : arg1 > arg2
	ge : arg1 >= arg2
	*/
	c.Data["n1"] = 12
	c.Data["n2"] = 6

	/*
		{{if gt .n1 .n2}}
		<p>n1大於n2</p>
		{{end}}
	*/
	//9.define自定義模板 template使用模板
	/*
		<h2>
		  10.define自定義模板
		</h2>
		{{define "aaa"}}
		<h4>這是一個自定義的代碼塊</h4>
		<p>123</p>
		<p>123523</p>
		{{end}}

		<div>
		  {{template "aaa" .}}  是在div內輸出 可多次調用
		</div>
	*/

	c.Data["unix"] = 1587880013 //驗證模板函式是否可重複調用
	//在index.view裡寫入<p>{{.unix | unixToDate}}</p>
	//main.go:beego.AddFuncMap("unixToDate", models.UnixToDate)

	c.TplName = "index.html"
}
