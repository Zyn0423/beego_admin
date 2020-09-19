package models

import ("github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"time"
)


type User struct {
	//TODO 这里字段大写不然回报错goadmin/controllers.User` needs a primary key field, default is to use 'id' if not set
	Id int
	UserName string
	Password string
	Article []*Article `orm:"rel(m2m)"`  //TODO 多对多
}

type Article struct {
	Id2 int `orm:"pk;auto"`  //pk主键
	Title string `orm:"size(20)"`//文章标题
	Content string `orm:"size(500)"` //内容500字以内
	Img string `orm:"size(50);nil"`//图路径50以内，默认为空
	Time time.Time `orm:"type(datatime);auto_now_add"`//时间设置
	Count int `orm:"default(0)"`//统计可以为0
	ArticleType *ArticleType `orm:"rel(fk)"`  //fk 外键
	User []*User `orm:"reverse(many)"`
}

type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`
	Article []*Article `orm:"reverse(many)"`  // TODO 一对多
}
func init()  {
	//1.连接数据库
	orm.RegisterDataBase("default","mysql","root:Zxn123456@tcp(127.0.0.1:3306)/new_web?charset=utf8")
	//2.注册表  用orm.RegisterModel()函数，参数是结构体对象，如果有多个表，可以用 , 隔开，多new几个对象：
	orm.RegisterModel(new(User),new(Article),new(ArticleType))
	//3.生成表 //1.数据库别名 //2.是否强制更新 //3.创建表过程是否可见
	orm.RunSyncdb("default",false,true)

}