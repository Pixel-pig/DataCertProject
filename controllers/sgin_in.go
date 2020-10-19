package controllers

import (
	"DaraCertProject/models"
	"github.com/astaxie/beego"
)

type Sgin_inConterller struct {
	beego.Controller
}

/**
 *执行登录操作
 */
func (s *Sgin_inConterller) Post() {
	//1.解析form表单中的数据
	var user models.User
	err := s.ParseForm(&user)
	if err != nil {
		panic("数据解析错误")
	}

	//2.跟获取到的数据与数据库中的数据进行比较
	u, err := user.QueryUserInfo()
	if err != nil {
		s.TplName = "login.html"
		return
	}

	//3.加载核心页面（现在没有）
	s.Data["phone"] = u.Phone
	s.TplName ="home.html"
}

/**
 *注册页面直接登录功能,加载页面
 */
func (s *Sgin_inConterller) Get() {
	//加载login页面
	s.TplName = "home.html"
}
