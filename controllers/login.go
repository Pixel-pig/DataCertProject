package controllers

import (
	"github.com/astaxie/beego"
)

type LoginConterller struct {
	beego.Controller
}


func (l *LoginConterller) Get() {
	//加载login页面
	l.TplName = "login.html"
}