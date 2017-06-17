package admin

import (
	"blog/models"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"

	"time"
	"strings"
	"fmt"
)

type ArticleController struct {
	BaseController
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

	//var relation models.Relation

	this.Data["article"] = article
	this.TplName = "admin/article/modify.html"
}

//TODO
func (this *ArticleController) Save(){
	var relation models.Relation
	var tag_model models.Tag
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
	fmt.Println(id)
	tag:=this.GetString("tag")
	if id<=0{

		article.Addtime = time.Now().Unix()
		err:=article.Insert()
		if err!=nil{
			flash.Error("添加失败")
		}else{
			if tag != ""{

				tagArr:=strings.Split(tag,",")
				for _,v:=range tagArr{
					tag_model.Name = v
					if err:=tag_model.Read("Name");err!=nil{
						tag_model.Addtime = time.Now().Unix()
						tag_model.Insert()
					}
					relation.Aid = article.Id
					relation.Tid = tag_model.Id
					relation.Addtime = time.Now().Unix()
					relation.Insert()
					relation = models.Relation{}
					tag_model = models.Tag{}
				}
			}

			flash.Success("添加成功")
		}
	}else{
		article.Id = id
		if err := article.Read();err!=nil{
			flash.Error("不存在该文章")
		}else{
			this.ParseForm(&article)
			article.Id = id
			article.Addtime = time.Now().Unix()
			err:=article.Update()
			if err!=nil{
				flash.Error("编辑失败")
			}else{
				if tag != ""{

					tagArr:=strings.Split(tag,",")
					for _,v:=range tagArr{
						tag_model.Name = v
						if err:=tag_model.Read("Name");err!=nil{
							tag_model.Addtime = time.Now().Unix()
							tag_model.Insert()
						}
						relation.Aid = article.Id
						relation.Tid = tag_model.Id
						relation.Addtime = time.Now().Unix()
						relation.Insert()
						relation = models.Relation{}
						tag_model = models.Tag{}
					}
				}

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