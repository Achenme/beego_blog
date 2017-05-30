package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
	"blog/controllers/admin"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&admin.AdminController{})
	beego.AutoRouter(&admin.ArticleController{})
	beego.AutoRouter(&admin.CatController{})
}
