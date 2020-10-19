package routers

import (
	"DaraCertProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//加载注册页面
    beego.Router("/", &controllers.MainController{})

    //执行注册操作，并提示（注册成功or失败的信息），成功后加载登录页面
    beego.Router("/register", &controllers.RegisterConterller{})

    //执行登录操作
    beego.Router("/sgin_in",&controllers.Sgin_inConterller{})

    //文件上传
	beego.Router("/upload", &controllers.UploadController{},"*:Post")
}
