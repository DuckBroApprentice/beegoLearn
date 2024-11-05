package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("用戶中心")
}

func (c *UserController) AddUser() {
	c.TplName = "userAdd.html" //渲染到user.html(在views裡面)，doAdd
}
func (c *UserController) DoAddUser() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("ID必須是INT類型")
		return
	}
	username := c.GetString("username")
	password := c.GetString("password") //GetString("xxx")要與user.html內的 input ... name="xxx"一致
	c.Ctx.WriteString("用戶中心--" + username + password + strconv.Itoa(id))

	hobby := c.GetStrings("hobby")
	fmt.Printf("%v----%T", hobby, hobby)
}

func (c *UserController) EditUser() {
	c.TplName = "userEdit.html"
}

type User struct {
	Username string   `form:"usernmae" json:"username"` //反射技術
	Password string   `form:"password" json:"password"`
	Hobby    []string `form:"hobby" json:"hobby"`
}

func (c *UserController) DoEditUser() {
	u := User{}
	err := c.ParseForm(&u) //利用ParseForm獲取表單數據
	if err != nil {
		c.Ctx.WriteString("post提交失敗")
		return
	}
	fmt.Printf("%#v", u)
	c.Ctx.WriteString("解析post數據成功")
}

//在beego中，返回json數據

func (c *UserController) GetUser() {
	u := User{
		Username: "張三",
		Password: "123456",
		Hobby:    []string{"1", "2"},
	}
	//返回json數據
	c.Data["json"] = u
	c.ServeJSON()
	//
}
