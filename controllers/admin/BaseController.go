package admin

import(
	"github.com/astaxie/beego"

)

var PageSize int = 10
type BaseController struct {
	beego.Controller
}

func (this *BaseController) isPost() bool{
	return this.Ctx.Request.Method == "POST"
}
