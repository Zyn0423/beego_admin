package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goadmin/models"
	"math"
	"path"
	"strconv"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController)ShowArticleGet()  {   //TODO 文章列表
	// 获取数据库数据
	pageSize := 1   // TODO 定义1页展示多少数据
	o :=orm.NewOrm()
	qt :=o.QueryTable("Article") //选择表
	var article[]  models.Article  //定义一个切片容器
	pageIndex:=this.GetString("pageIndex")
	pageIndex1,err:=strconv.Atoi(pageIndex)//字符串转int
	if err!=nil{
		beego.Info("错误pageIndex1：",pageIndex1)
		pageIndex1 =1
	}
	//------------>
	//首先查数据库有多少条数据
	count,err:=qt.Count()
	if err !=nil{
		beego.Info("查询多少条数据失败",err)
		return
	}
	start := pageSize *(pageIndex1 -1)  //todo  0 ->2->4    每页展示数据*（当前页 -1）=数据库拿的数据
	_,err=qt.Limit(pageSize,start).All(&article)
	if err !=nil{
		beego.Info("查询数据失败",err)
		return
	}
	firstPage :=false  //todo 固定上一页按钮
	if pageIndex1 <2 {
		firstPage =true
	}
	countPage := math.Ceil(float64(count)/float64(pageSize))  //向上取整  总共有多少页	//TODO 总数/展示页数 = 总共有多少页
	nextPage :=false
	if pageIndex1 >int(countPage){
		nextPage =true
	}
	this.Data["firstPage"] =firstPage  //todo 5.方向标 控制上页面的超链接显示
	this.Data["nextPage"] =nextPage  //todo 5.方向标 控制下页面的超链接显示
	this.Data["countPage"] =countPage  //todo 2.总共有多少页
	this.Data["article"]=article		//TODO 4.传出数据
	this.Data["pageIndex"] = pageIndex1   //TODO 3. pageIndex1  而不是pageIndex
	this.Data["count"] =count  //todo 1.总共多少条数据
	this.TplName="index.html"
}

func (this *ArticleController)ShowAddarticleGet()  {  //文章列表

	this.TplName="add.html"
}
func (this *ArticleController)HandleAddarticle()  {
	const filesize  = 5000000
	articName:=this.GetString("articleName")
	content:=this.GetString("content")
	artype := this.GetString("select")

	beego.Info("----------->文章类型：",artype)
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
	//article.ArticleType =
	//article.User = ""
	//插入数据
	_,err =o.Insert(&article)
	if err !=nil{
		 beego.Info("插入数据失败",err)

	}
	this.Redirect("article",302)

}

func (this *ArticleController) ShowArticleDetail() {
	//获取数据
	id_,er:=this.GetInt("articleId")   // TODO 新的方法获取
	//数据校验
	if er != nil{
		beego.Info("传递的链接错误")
	}
	beego.Info("获取的ID",id_)
	o :=orm.NewOrm()
	//获取数据
	var article models.Article
	article.Id2 = id_
	err :=o.Read(&article)
	if err!=nil{
		beego.Info("ID获取数据失败")
		return
	}

	//修改阅读量
	article.Count+=1
	o.Update(&article)  //更新数据
	//返回视图页面
	this.Data["article"] = article
	this.TplName = "content.html"

}

func (this *ArticleController)ShowDeleteDetail()  {  //TODO 列表详情页删除
	//先获取前端传来的数据
	id,_:=this.GetInt("id")
	//创建数据容器
	var article models.Article
	article.Id2 = id
	//创建数据库对象
	o :=orm.NewOrm()
	_,err :=o.Delete(&article)
	if err!= nil{
		beego.Info("删除数据失败")
		return
	}
	this.Redirect("/article",302)

}
func (this *ArticleController)ShowUpdataDetail()  {
	//获取前端传来数据
	id ,_:=this.GetInt("id1")
	var article models.Article
	article.Id2 = id
	o :=orm.NewOrm()
	err := o.Read(&article)
	if err!= nil{
		beego.Info("查询数据失败")
		return
	}
	this.Data["article"] = article
	this.TplName="update.html"
}

func (this *ArticleController)HandleUpdataDetail()  {   //TODO 修改文章列表的指定编辑信息
	const filesize  = 5000000
	articName:=this.GetString("articleName")
	content:=this.GetString("content")
	if articName == "" && content == ""{
		beego.Info("没有数据请填写数据")
		return
	}
	id,_ :=this.GetInt("id")

	f,h,err:=this.GetFile("uploadname")  //TODO 获取上传图片
	beego.Info(articName,content)
	if err !=nil{
		beego.Info("上传图片失败",err)
		//创建数据库对象
		o:=orm.NewOrm()
		//TODO 更新数据前先查数据
		var article =models.Article{Id2:id}
		err =o.Read(&article)
		if err !=nil{
			beego.Info("查询数据失败",err)
			return
		}
		//整理数据
		article.Title = articName
		article.Content = content
		//插入数据
		beego.Info(article)
		_,err =o.Update(&article)
		if err !=nil{
			beego.Info("更新数据失败",err)

		}
		this.Redirect("article",302)

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
	// TODO 一直上传不上去"./static/img/"  加点
	this.SaveToFile("uploadname","./static/img/"+t+ext)
	if err !=nil{
		beego.Info("保存图片失败",err)
		return
	}
	//创建数据库对象
	o:=orm.NewOrm()
	//TODO 更新数据前先查数据
	var article =models.Article{Id2:id}
	err =o.Read(&article)
	if err !=nil{
		beego.Info("查询数据失败",err)
		return
	}
	//整理数据
	article.Title = articName
	article.Content = content
	article.Img = "./static/img/"+t+ext
	//插入数据
	beego.Info(article)
	_,err =o.Update(&article)
	if err !=nil{
		beego.Info("更新数据失败",err)
	}
	this.Redirect("article",302)
}

func (this *ArticleController)ShowAddType()  {
	var article[] models.ArticleType
	o :=orm.NewOrm()
	_,err :=o.QueryTable("ArticleType").All(&article)
	if err!=nil{
		beego.Info("查询数据失败",err)
		return
	}
	this.Data["types"] = article
	this.TplName = "addType.html"
	//this.Redirect("/addTypeDetail",302)

}
func (this *ArticleController)HandleAddTpye()  {
	//获取前端传来的数据
	 typsinfo:=this.GetString("typeName")
	//进行数据判断是否为空
	if typsinfo ==""{
		beego.Info("数据不能为空")
		return
	}
	o :=orm.NewOrm()
	var articletype models.ArticleType
	articletype.TypeName=typsinfo
	_,err :=o.Insert(&articletype)
	if err != nil{
		beego.Info("插入数据失败")
		return
	}

	//创建数据库对象并把数据插入到数据库中
	this.Redirect("addTypeDetail",302)
}