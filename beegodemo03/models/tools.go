package models

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/beego/beego"
)

/*
var timestamp int64 = 1587880013  //時間戳
t := time.Unix(timestamp,0)  //日期對象
fmt.Println(t.Format("2006-01-02 03:04:05"))  //日期格式化輸出
*/

// main.go調用beego.AddFuncMap("unixToDate", models.UnixToDate)
func UnixToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0) //不傳入奈秒(0)
	return t.Format("2006-01-02 15:04:05")
}

// 在控制器內調用
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:04:05" //定義模板字符串
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		beego.Info(err)
		return 0
	}
	return t.Unix()
}

func GetUnix() int64 {
	return time.Now().Unix()
}
func GetDate() string {
	template := "2006-01-02 15:04:05"
	t := time.Now()
	return t.Format(template)
}

func Hello(in string) (out string) {
	out = in + "world"
	return
}

func Md5(str string) string {
	data := []byte(str)
	return fmt.Sprint("%x\n", md5.Sum(data))
}
