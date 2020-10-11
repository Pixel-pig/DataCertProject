package main

import (
	"DaraCertProject/db_mysql"
	_ "DaraCertProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	//1.链接数据库
	db_mysql.Connect()

	//2.设置静态资源路径
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

