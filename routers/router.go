package routers

import (
	"github.com/astaxie/beego"
	"goadmin/controllers"
	"github.com/astaxie/beego/context"
)

var FliterFunc = func (ctx *context.Context)  {
	user :=ctx.Input.Session("userName")
	if user == nil{
		ctx.Redirect(302,"/")
	}
}
func init() {
	beego.InsertFilter("/Article/*",beego.BeforeRouter,FliterFunc)
  	beego.Router("/", &controllers.LoginController{},"get:ShowLogin;post:HandleLogin")  // TODO 登录
	beego.Router("/register",&controllers.RegisterController{},"get:ShowRegisterGet;post:HandleRegister") // TODO 注册
	beego.Router("/Article/article",&controllers.ArticleController{},"get:ShowArticleGet;post:HandleArticlePost") // TODO 登陆成功后台页面
	beego.Router("/Article/addarticle",&controllers.ArticleController{},"get:ShowAddarticleGet;post:HandleAddarticle") //TODO 添加文章
	beego.Router("/Article/showArticleDetail",&controllers.ArticleController{},"get:ShowArticleDetail") //TODO 查看详情
	beego.Router("/Article/deleteArticleDetail",&controllers.ArticleController{},"get:ShowDeleteDetail") //TODO 删除详情
	beego.Router("/Article/updataArticleDetail",&controllers.ArticleController{},"get:ShowUpdataDetail;post:HandleUpdataDetail") //TODO 修改详情
	beego.Router("/Article/addTypeDetail",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddTpye") //TODO 修改详情
	beego.Router("/Article/deleteTypeDetail",&controllers.ArticleController{},"get:ShowDeleteTypeDetail") //TODO 删除详情
	beego.Router("/Article/logout",&controllers.ArticleController{},"get:ShowLogout") //TODO 删除详情
	//beego.Router("/index", &controllers.IndexController{})
	//beego.Router("/index", &controllers.IndexController{},"get:ShowGet;post:Post")  // TODO 指定路由
	//beego.Router("/index/?:id", &controllers.IndexController{},"get:ShowGet;post:Post")  // TODO 正则路由
	//beego.Router("/index/mysql", &controllers.MysqlController{},"get:ShowMysql")  //
	//beego.Router("/index/ormsql", &OrmMysqlController{},"get:ShowGoMysql")     //  创建表
	//beego.Router("/index/insert", &OrmMysqlController{},"get:ShowInsertMysql") // 插入
	//beego.Router("/index/find", &OrmMysqlController{},"get:ShowFindMysql")     // 查
	//beego.Router("/index/updata", &OrmMysqlController{},"get:ShowupdataMysql") // 更新
	//beego.Router("/index/delete", &OrmMysqlController{},"get:ShowDeleteMysql") // 删除

		//beego.Router("/register",&RegisterController{},"get:ShowRegisterGet;post:ShowRegisterPost")
}
