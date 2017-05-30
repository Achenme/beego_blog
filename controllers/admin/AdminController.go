package admin


type AdminController struct {
	BaseController
}


func (this *AdminController) Login() {
	
}

func (this *AdminController) Logout() {
	this.TplName = "admin/logout.html"
}

