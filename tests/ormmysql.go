package test
//
//import (
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/orm"
//	_ "github.com/go-sql-driver/mysql"
//	"goadmin/models"
//)
//
//type OrmMysqlController struct {
//	beego.Controller
//}
//
//
////todo 创建表1
///*func (this *OrmMysqlController) ShowGoMysql() {
//	//1.连接数据库  aliasName 别名
//	orm.RegisterDataBase("default","mysql","root:Zxn123456@tcp(127.0.0.1:3306)/class1?charset=utf8")
//	//2.注册表
//	orm.RegisterModel(new(User))
//	//3.生成表   force 强制刷新表结构   verbose 查看创建表的过程
//	orm.RunSyncdb("default",false,true)
//	this.Ctx.WriteString("创建表成功")
//}*/
////todo 创建表2
//func (this *OrmMysqlController) ShowGoMysql()  {
//	orm.RunSyncdb("default",false,true)
//	this.Ctx.WriteString("创建表成功")
//}
//// todo 插入数据
//func (this *OrmMysqlController) ShowInsertMysql()  {
//	//先获取一个ORM对象,用orm.NewOrm()即可获得
//	o :=orm.NewOrm()
//	//定义一个要插入数据库的结构体对象  var user User
//	user := models.User{}
//	//给定义的对象赋值
//	user.Name="吴彦祖"
//	//执行插入操作，o.Insert()插入，参数是结构体对象，返回值是插入的id和错误信息。
//	_ , err :=o.Insert(&user)
//	if err != nil{
//		beego.Info("插入数据失败")
//		return
//	}
//	this.Ctx.WriteString("插入数据成功")
//}
//
////TODO 查询数据
//func (this *OrmMysqlController) ShowFindMysql()  {
//	//先获得一个ORM对象
//	o :=orm.NewOrm()
//	//获取查询对象  定义一个要获取数据的结构体对象  var user User
//
//	user := models.User{}
//	//给结构体对象赋值，相当于给查询条件赋值
//	user.Id = 1
//	//执行操作函数  查询,用o.Read()，第一个参数是对象地址，第二个参数是指定查询字段，返回值只有错误信息
//	//TODO 如果查询字段是查询对象的主键的话，可以不用指定查询字段
//	err := o.Read(&user)
//	if err != nil{
//		beego.Info("查询数据失败")
//		return
//	}
//	beego.Info(user.Name)
//	this.Ctx.WriteString("查询数据成功")
//
//}
////TODO 更新数据
//func (this *OrmMysqlController) ShowupdataMysql() {
//	//创建连接对象
//	 o :=orm.NewOrm()
//	 //获取查询对象
//	 user := models.User{Id: 1}
//	 //查询要更新的对象是否存在
//	 err :=o.Read(&user)
//	 if err!= nil{
//	 	beego.Info("查询数据失败")
//		 return
//	 }
//	 //如果查找到了要更新的对象,就给这个对象赋新值
//	 user.Name= "张三丰"
//	 //执行更新操作,用o.Update()函数，参数是结构体对象指针，返回值是更新的条目数和错误信息
//	 _,err=o.Update(&user)
//	 if err !=nil{
//	 	beego.Info("更新数据失败")
//		 return
//	 }
//
//	this.Ctx.WriteString("更新数据成功")
//
//}
//
////TODO 删除数据
//
//func (this *OrmMysqlController) ShowDeleteMysql() {
//	//创建对象
//	o:=orm.NewOrm()
//	//获取查询对象
//	user := models.User{Id: 1}
//	//执行删除操作
//	_,err :=o.Delete(&user)
//	if err !=nil{
//		beego.Info("删除数据失败")
//		return
//	}
//	this.Ctx.WriteString("删除数据成功")
//
//}