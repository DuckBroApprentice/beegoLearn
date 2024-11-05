package main

import (
	"beegodemo03/models"
	_ "beegodemo03/routers"

	"github.com/astaxie/beego"
)

/*
在配置文件修改
conf/app.conf
TemplateLeft="<<"
TemplateRight=">>"
*/
//或是在main.go內設定

func main() {
	// beego.BConfig.WebConfig.TemplateLeft = "<<"
	// beego.BConfig.WebConfig.TemplateRight = ">>"
	beego.AddFuncMap("hi", models.Hello) //自定義模板函式
	//{{.title | hi}} 這是文章列表world

	//[Add FuncMap] 讓使用者在範本中註冊一個函數。
	//UnixToDate 函式寫在models/tools.go內
	beego.AddFuncMap("unixToDate", models.UnixToDate)
	//先在model/tools.go封裝方法，後在入口注冊，使該方法能在各個模板裡調用(即公共的功能)
	//注意：要在模板(views)裡調用需要在入口注冊(控制器則不用在入口注冊)
	//<p>{{.unix | unixToDate}}</p>
	//在控制器(controllers)調用就跟一般調用函式一樣(import包 包.方法)

	//配置靜態資源目錄  默認static 通常不需要這樣做
	beego.SetStaticPath("/down", "download") //將download的url設為down,path 一個 download
	/*
		SetStaticPath sets static directory path and proper url pattern in beego application.
		if beego.SetStaticPath("static","public"), visit /static/* to load static file in folder "public".
	*/
	//localhost:8080/down/222.zip  down是自定義的靜態資源目錄
	//localhost:8080/static/222.zip  static是默認路徑，非本例操作
	beego.Run()
}
