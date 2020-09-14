package models

import ("github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
)


type User struct {
	//TODO 这里字段大写不然回报错goadmin/controllers.User` needs a primary key field, default is to use 'id' if not set
	Id int
	UserName string
	Password string
}

func init()  {
	//1.连接数据库
	orm.RegisterDataBase("default","mysql","root:Zxn123456@tcp(127.0.0.1:3306)/new_web?charset=utf8")
	//2.注册表  用orm.RegisterModel()函数，参数是结构体对象，如果有多个表，可以用 , 隔开，多new几个对象：
	orm.RegisterModel(new(User))
	//3.生成表 //1.数据库别名 //2.是否强制更新 //3.创建表过程是否可见
	orm.RunSyncdb("default",false,true)

}