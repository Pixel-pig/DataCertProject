package controllers

import (
	"DaraCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterConterller struct {
	beego.Controller
}

//加载login页面既刷新页面
func (r *RegisterConterller) Get() {
	r.TplName = "login.html"
}

/**
* 重写一个post的方法来执行注册的操作
*/
func (r *RegisterConterller) Post() {
	//1.初始化一个 user 对象，将form表单中的数据解析到结构体中
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		panic("数据解析失败")
	}

	//2.将结构体中的数据存入MySQL数据库中
	_, err = user.SaveUserInfo()
	if err != nil {
		fmt.Println(err.Error())
	}

	//3.加载登录页面
	r.TplName = "login.html"
}
