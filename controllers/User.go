package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goadmin/models"
)


type MainController struct {
	beego.Controller
}
type RegisterController struct {
	beego.Controller
}


func (this *MainController)Get(){
	this.TplName = "login.html"
}

func (this *RegisterController) ShowRegisterGet() {
	this.TplName = "register.html"

}
func (this *RegisterController)HandleRegister()  {
	// 1. 拿浏览器数据
	//2.数据处理
	//3. 插入数据库
	//4.返回视图
	name :=this.GetString("userName")
	password :=this.GetString("password")
	if name == "" || password =="" {
		beego.Info("输入的账号密码不能为空")
		this.Data["err_info"] = "输入的账号密码不能为空"
		this.TplName="register.html"
		return
	}
	o:=orm.NewOrm()
	user:=models.User{}
	user.UserName=name
	user.Password=password
	_,err:=o.Insert(&user)
	if err!=nil{
		beego.Info("插入数据失败")
	}

	beego.Info(name,password)
	this.Ctx.WriteString("插入数据成功")
	//this.TplName="login.html"
}