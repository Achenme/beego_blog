package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
	"blog/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.AutoRouter(&admin.AdminController{})
	beego.AutoRouter(&admin.ArticleController{})
	beego.AutoRouter(&admin.CatController{})
	beego.AutoRouter(&controllers.IndexController{})
}
