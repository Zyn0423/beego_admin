package main

import (
	_ "goadmin/routers"
	"github.com/astaxie/beego"
	_"goadmin/models"
)

func main() {
	beego.Run()
}

