package controllers

import (
	"DaraCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type Sgin_inConterller struct {
	beego.Controller
}

func (s *Sgin_inConterller) Post() {
	//1.解析form表单中的数据
	var user models.User
	err := s.ParseForm(&user)
	if err != nil {
		panic("数据解析错误")
	}

	//2.跟获取到的数据与数据库中的数据进行比较
	err = user.QueryUserInfo()
	if err != nil {
		fmt.Println(err.Error())
	}

	//3.加载核心页面（现在没有）
	s.Ctx.WriteString("核心页面")
}
