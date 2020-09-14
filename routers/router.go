package routers

import (
	"goadmin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	//beego.Router("/index", &controllers.IndexController{})
	//beego.Router("/index", &controllers.IndexController{},"get:ShowGet;post:Post")  // TODO 指定路由
	//beego.Router("/index/?:id", &controllers.IndexController{},"get:ShowGet;post:Post")  // TODO 正则路由
	beego.Router("/index/mysql", &controllers.MysqlController{},"get:ShowMysql")  //
}
