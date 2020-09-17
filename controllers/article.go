package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goadmin/models"
	"path"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController)ShowArticleGet()  {
	// 获取数据库数据
	o :=orm.NewOrm()
	qt :=o.QueryTable("Article") //选择表
	var article[]  models.Article  //定义一个切片容器
	_,err := qt.All(&article) //查询数据
	if err !=nil{
		beego.Info("查询数据失败",err)
		return

	}
	this.Data["article"]=article
	this.TplName="index.html"
}

func (this *ArticleController)ShowAddarticleGet()  {  //文章列表

	this.TplName="add.html"
}
func (this *ArticleController)HandleAddarticle()  {
	const filesize  = 5000000
	articName:=this.GetString("articleName")
	content:=this.GetString("content")
	f,h,err:=this.GetFile("uploadname")  //TODO 获取上传图片
	if err !=nil{
		beego.Info("上传图片失败",err)
		return
	}
	defer f.Close()
	//判断文件格式
	ext :=path.Ext(h.Filename)

	if ext!= ".png" && ext!= ".jpg"&&ext!= ".jpeg" {
		beego.Info("上传的图片格式错误")
		return
	}
	if h.Size >filesize {
		beego.Info("上传的图片大于",filesize,"字节")
		return
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	beego.Info(t+ext)
	// TODO 一直上传不上去"./static/img/"  加点
	this.SaveToFile("uploadname","./static/img/"+t+ext)
	if err !=nil{
		beego.Info("保存图片失败",err)
		return
	}
	//创建数据库对象
	o:=orm.NewOrm()
	//整理数据
	var article =models.Article{}
	article.Title = articName
	article.Content = content
	article.Img = "./static/img/"+t+ext
	//插入数据
	_,err =o.Insert(&article)
	if err !=nil{
		 beego.Info("插入数据失败",err)

	}
	this.Redirect("article",302)



}