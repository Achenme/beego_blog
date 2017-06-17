package admin

import (
	"strings"
	"github.com/astaxie/beego"
	"fmt"
)

type AdminController struct {
	BaseController
}


func (this *AdminController) Login() {
	if this.isPost(){
		username:=strings.TrimSpace(this.GetString("username"))
		password:=strings.TrimSpace(this.GetString("password"))
		config_name :=beego.AppConfig.String("username")
		config_password := beego.AppConfig.String("password")
		md5_password:=this.Md5([]byte(password))
		fmt.Println(md5_password)
		fmt.Println(config_password)
		if username != config_name{
			this.Ctx.WriteString("<script>alert('你输入的账号错误');history.go(-1)</script>")
			return
		}

		if config_password!=md5_password{
			this.Ctx.WriteString("<script>alert('你输入的密码错误');history.go(-1)</script>")
			return
		}
		authkey := this.Md5([]byte(this.getClientIp() + "|" + config_password))
		this.Ctx.SetCookie("auth",username+"|"+authkey)
		this.Redirect("/article/index",302)
	}
	this.TplName = "admin/login.html"
}

func (this *AdminController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.Redirect("/admin/login", 302)
}

