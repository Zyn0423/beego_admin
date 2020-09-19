package routers

import (
	"github.com/astaxie/beego"
	"goadmin/controllers"
)

func init() {
  	beego.Router("/", &controllers.LoginController{},"get:ShowLogin;post:HandleLogin")  // TODO 登录
	beego.Router("/register",&controllers.RegisterController{},"get:ShowRegisterGet;post:HandleRegister") // TODO 注册
	beego.Router("/article",&controllers.ArticleController{},"get:ShowArticleGet") // TODO 登陆成功后台页面
	beego.Router("/addarticle",&controllers.ArticleController{},"get:ShowAddarticleGet;post:HandleAddarticle") //TODO 添加文章
	beego.Router("/showArticleDetail",&controllers.ArticleController{},"get:ShowArticleDetail") //TODO 查看详情
	beego.Router("/deleteArticleDetail",&controllers.ArticleController{},"get:ShowDeleteDetail") //TODO 删除详情
	beego.Router("/updataArticleDetail",&controllers.ArticleController{},"get:ShowUpdataDetail;post:HandleUpdataDetail") //TODO 修改详情
	beego.Router("/addTypeDetail",&controllers.ArticleController{},"get:ShowAddType;post:HandleAddTpye") //TODO 修改详情
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
