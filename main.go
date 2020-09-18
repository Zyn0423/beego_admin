package main

import (
	_ "goadmin/routers"
	"github.com/astaxie/beego"
	_"goadmin/models"
)

func main() {
	//todo 视图函数  key 是前端｜后的名称 fn 定义的方法
	beego.AddFuncMap("ShowPrePage",HandlePrePage)
	beego.AddFuncMap("ShowNextPage",HandleNextPage)
	beego.Run()
}
func HandlePrePage(data int) (int) {
	return data -1
}
func HandleNextPage(data int) (int) {
	return data+1
}

