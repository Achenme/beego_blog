package admin

import (
	"blog/models"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
)

type ArticleController struct {
	BaseController
}
func (this *ArticleController) Prepare() {
	this.Layout = "admin/layout.html"
}


func (this *ArticleController) Index() {
	var article models.Article
	var list []*models.Article
	query:=article.Query()
	count,_:=query.Count()
	query.All(&list)
	Pagbar:=models.Paginator(1,1,count)
	this.Data["Pagbar"] = Pagbar
	this.Data["list"] = list
	this.TplName = "admin/article/index.html"
}

func (this *ArticleController) Add() {
	beego.ReadFromRequest(&this.Controller)
	this.TplName = "admin/article/modify.html"
}

func (this *ArticleController) Edit() {
	beego.ReadFromRequest(&this.Controller)
	id,_:= this.GetInt("id")
	var article models.Article
	article.Id = id
	if err:=article.Read();err!=nil{
		this.Abort("404")
	}
	this.Data["article"] = article
	this.TplName = "admin/article/modify.html"
}

//TODO
func (this *ArticleController) Save(){
	var article models.Article
	valid := validation.Validation{}
	this.ParseForm(&article)

	flash:=beego.NewFlash()
	b, err :=valid.Valid(&article)
	if err!=nil{
		flash.Error("接受参数无效,请重试")
	}

	if !b{
		for _, err := range valid.Errors {
			flash.Set(err.Key,err.Message)
		}
	}

	id,_:= this.GetInt("id")
	if id<=0{
		err:=article.Insert()
		if err!=nil{
			flash.Error("添加失败")
		}else{
			flash.Success("添加成功")
		}
	}else{
		article.Id = id
		if err := article.Read();err!=nil{
			flash.Error("不存在该文章")
		}else{
			this.ParseForm(&article)
			article.Id = id
			err:=article.Update()
			if err!=nil{
				flash.Error("编辑失败")
			}else{
				flash.Success("编辑成功")
			}
		}

	}

	flash.Store(&this.Controller)
	this.Redirect("/article/add", 302)
}

func (this *ArticleController) Delete()  {
	id,_:= this.GetInt("id")
	var article models.Article
	article.Id = id
	if err:=article.Read();err==nil{
		article.Delete()
	}

	this.ServeJSON()
}