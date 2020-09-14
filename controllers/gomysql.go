package controllers

import (
	"database/sql"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlController struct {
	beego.Controller
}

func (this *MysqlController) ShowMysql()  {
	//连接数据库
	//操作数据库
	//关闭数据库
	//连接数据库，用sql.Open()方法,open()方法的第一个参数是驱动名称,第二个参数是用户名:密码
	//@tcp(ip:port)/数据库名称?编码方式,返回值是连接对象和错误信息
	conn,err :=sql.Open("mysql","root:Zxn123456@tcp(127.0.0.1:3306)/class1?charset=utf8")
	if err !=nil{
		beego.Info("连接失败",err)
		return
	}
	defer conn.Close()
	//
	//TODO  创建表结构
	/*_ ,err =conn.Exec("create table userInfo(id int ,name varchar(12) )")
	if err !=nil{
		beego.Info("连接失败",err)
		return

	}
	this.Ctx.WriteString("创建成功")*/


	//TODO 增加数据
	/*_ ,err =conn.Exec("insert userInfo(id,name) VALUE (?,?)",4,"周星驰")
	if err!=nil{
		beego.Info("增加数据失败",err)
		return
	}
	this.Ctx.WriteString("增加数据成功")*/
	//TODO 查询
	/*rows,err := conn.Query("select id from userInfo")
	if err!=nil{
		beego.Info("查询数据失败",err)
		return
	}
	var id  int
	for rows.Next(){
		rows.Scan(&id)
		beego.Info(id)
	}
	this.Ctx.WriteString("查询数据成功")*/

}