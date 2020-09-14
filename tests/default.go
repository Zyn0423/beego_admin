
package test
//
//import (
//	"github.com/astaxie/beego"
//)
//
//type MainController struct {
//	beego.Controller
//}
//
//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.Data["key_"] = "goadmin"
//	c.TplName = "index_1.html"
//}
//
//type IndexController struct {
//	 beego.Controller
//}
//
//func (c *IndexController) Post()  {
//
//	c.Data["key_"] = "Post修改成功"
//	c.TplName = "index_1.html"
//}
//
//func (this *IndexController)ShowGet()  {
//	id :=this.GetString(":id")
//	beego.Info(id)
//	this.Data["key_"] = "ShowGet修改成功"
//	this.TplName = "index_1.html"
//}