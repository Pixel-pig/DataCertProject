package controllers

import (
	"crypto/sha256"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"path"
)

type UploadController struct {
	beego.Controller
}

//刷新本页面
func (this *UploadController) Get() {
	this.TplName = "home.html"
}


func (this *UploadController) Post() {
	//1.获取用户上传的文件
	file, information, err := this.GetFile("file") //返回文件，文件信息头，错误信息
	if err != nil {
		this.Ctx.WriteString("File retrieval failure")
		return
	}
	defer file.Close() //关闭上传的文件，否则出现临时文件不清除的情况  mmp错了好多次啊

	filename := information.Filename //将文件信息头的信息赋值给filename变量

	err = this.SaveToFile("file", path.Join("static/upload", filename)) //保存文件的路径。保存在static/upload中   （文件名）
	if err != nil {
		this.Ctx.WriteString("File upload failed！")
	} else {
		this.Ctx.WriteString("File upload succeed!") //上传成功后显示信息
	}
	//2.打开该文件获取该文件的事件句柄
	f, err := os.Open("./static/upload/"+filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	h := sha256.New()

	if _, err := io.Copy(h, f); err != nil {
		fmt.Println(err.Error())
		return
	}



	this.TplName = "upload.html" //停留在当前界面
}
