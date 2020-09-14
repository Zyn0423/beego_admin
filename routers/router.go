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
	//beego.Router("/index/mysql", &controllers.MysqlController{},"get:ShowMysql")  //
	beego.Router("/index/ormsql", &controllers.OrmMysqlController{},"get:ShowGoMysql")  //  创建表
	beego.Router("/index/insert", &controllers.OrmMysqlController{},"get:ShowInsertMysql")  // 插入
	beego.Router("/index/find", &controllers.OrmMysqlController{},"get:ShowFindMysql")  // 查
	beego.Router("/index/updata", &controllers.OrmMysqlController{},"get:ShowupdataMysql")  // 更新
	beego.Router("/index/delete", &controllers.OrmMysqlController{},"get:ShowDeleteMysql")  // 删除
}
