package admin

import(
	"github.com/astaxie/beego"
	"crypto/md5"
	"encoding/hex"
	"strings"

)

var PageSize int = 10
type BaseController struct {
	beego.Controller
	username string
}

func (this *BaseController) Prepare() {
	controllerName, actionName := this.GetControllerAndAction()
	if !(controllerName == "AdminController" && (strings.ToLower(actionName )== "login" || strings.ToLower(actionName) =="logout")) {
		this.auth()
		this.Layout = "admin/layout.html"
	}


}

func (this *BaseController) isPost() bool{
	return this.Ctx.Request.Method == "POST"
}

func (this *BaseController) Md5(b []byte) string  {
	h:=md5.New()
	h.Write(b)
	x:=h.Sum(nil)

	y:=make([]byte,32)
	hex.Encode(y,x)
	return string(y)
}


func (this *BaseController) auth(){

	arr:=strings.Split(this.Ctx.GetCookie("auth"),"|")

	if len(arr) == 2{

		config_username:=beego.AppConfig.String("username")
		config_password:=beego.AppConfig.String("password")
		check_password :=this.Md5([]byte(this.getClientIp() + "|" + config_password))
		if arr[0] == config_username && check_password == arr[1]  {
			this.username = arr[0]
		}

	}

	if len(this.username) == 0{
		this.Redirect("/admin/login",302)
	}

}

func (this *BaseController) getClientIp() string  {
	s:=strings.Split(this.Ctx.Request.RemoteAddr,":")
	return s[0]
}