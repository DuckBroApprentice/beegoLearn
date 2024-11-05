package main

import (
	_ "beegodemo04/routers"

	"github.com/astaxie/beego"
)

func main() {

	//配置redis數據庫的連接地址
	//Set可以配置全局配置的值
	beego.AppConfig.Set("redisuser", "admin666") //這樣就能在別處獲取redisuser

	//加載其他配置文件 conf/redis.conf
	//beego.LoadAppConfig("ini", "conf/redis.conf")

	beego.Run()
}

//以下通常用不到，大規模的團隊才會用到
/*
[dev]
httpport = 8000
[prod]
httpport = 8088
[test]
httpport = 8888
for運行
bee run -runmode dev
bee run -runmode prod
bee run -runmode test
for訪問配置訊息
beego.AppConft.String("dev::mysqluser")
beego.AppConft.String("prod::mysqluser")
beego.AppConft.String("test::mysqluser")
*/

/*
加載其他配置文件 conf/redis.conf
beego.LoadAppConfig("ini","conf/redis.conf")
第一個參數為加載模式，beego目前支持INI,XML,JSON,YAML格式的配置文件解析
，但是默認採用了INI格式解析，用戶可以通過簡單的配置就可以獲得很大的靈活性。
*/
