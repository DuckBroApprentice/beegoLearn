package controllers

import (
	"beegodemo03/models"
	"fmt"

	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

//localhost:8080/api
func (c *ApiController) Get() {
	unix := 1589980013

	str := models.UnixToDate(unix)

	s1 := "2020-05-19 22:20:13"
	u := models.DateToUnix(s1)
	fmt.Println(u)

	fmt.Println(models.GetUnix())
	fmt.Println(models.GetDate())

	//md5加蜜  不需要背但要知道
	// h := md5.New()
	// io.WriteString(h, "123456")
	// fmt.Printf("%x\n", h.Sum(nil)) //e10adc3949ba59abbe56e057f20f883e
	//加密後的值與想要後的值是否一樣
	/*
		data := []byte("123456")
		fmt.Printf("%x\n", md5.Sum(data)) //e10adc3949ba59abbe56e057f20f883e
		//接下來嘗試封裝的效果
	*/
	fmt.Println(models.Md5("123456"))

	c.Ctx.WriteString("api接口--" + str)

}
