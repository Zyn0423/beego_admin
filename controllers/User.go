package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goadmin/models"
	"time"
)

//TODO 1.登录控制器
type LoginController struct {
	beego.Controller
}

//登录Get
func (this *LoginController)ShowLogin(){
	userName:=this.Ctx.GetCookie("userName")
	//beego.Info("获取userName--->",userName)
	if userName !=""{
		this.Data["check"] = "checked"
		this.Data["name"] =userName
	}
	this.TplName = "login.html"
}
//登录Post
func (this *LoginController)HandleLogin()  {
	//1.获取浏览器用户数据
	name :=this.GetString("userName")
	password :=this.GetString("password")
	check :=this.GetString("remember")
	beego.Info("------>check",check)
	//2.处理数据
	user :=models.User{}
	if name == "" || password =="" {
		beego.Info("输入的账号密码不能为空")
		this.TplName="login.html"
		return
	}
	user.UserName =name
	//查询数据库
	o :=orm.NewOrm()
	err :=o.Read(&user,"user_name")
	if err != nil{
		beego.Info("查询账号失败")
		this.TplName="login.html"
		return
	}

	// 密码验证
	if user.Password != password {
		beego.Info("密码错误")
		this.TplName="login.html"
		return
	}
	if check == "on"{  //判断是否勾选☑️
		this.Ctx.SetCookie("userName",name,time.Second*3600)
		beego.Info("POST --->",check,name)
	}
	this.Redirect("/article",302)
}




//TODO 2.注册控制器
type RegisterController struct {
	beego.Controller
}
//注册Get
func (this *RegisterController) ShowRegisterGet() {
	this.TplName = "register.html"

}
//注册Post
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
	this.Redirect("/",302)    //todo 浏览器重定向发送请求，不可以发送数据    TplName 服务器端的不发送请求 可以发数据
}