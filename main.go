package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	_"blog/models"
)



func main() {
	beego.SetStaticPath("/css","static/css/admin")
	beego.SetStaticPath("/js","static/js/admin")
	beego.SetStaticPath("/images","static/img")
	beego.SetStaticPath("/fonts","static/fonts")


	beego.AddTemplateExt("html")

	beego.Run()
}

