package admin

type CatController struct {
	BaseController
}

func (this *CatController) Prepare() {
	this.Layout = "admin/layout.html"
}

func (this *CatController) Index() {
	this.TplName = "article.html"
}

func (this *CatController) Modify() {
	this.TplName = "article_add.html"
}

func (this *CatController) Delete()  {

}