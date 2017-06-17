package controllers

import (
	"blog/models"
	//"sync"
	"sync"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	var article models.Article
	var list []*models.Article
	var relation *models.Relation
	//var relationlist []*models.Relation
	//var tag models.Tag
	name:=this.GetString("name")
	page,_:=this.GetInt("page")

	query:=article.Query()
	count,_ := query.Count()
	if len(name) != 0 {
		query = query.Filter("name__contains", name)
	}
	query.All(&list)

	var a map[int][]string
	a = make(map[int][]string)


	var nn sync.WaitGroup
	for k,_:=range list{
		nn.Add(1)
		go func(k int) {
			defer nn.Done()
			relationlist := []*models.Relation{}
			query:=relation.Query()
			query.Filter("aid",list[k].Id).All(&relationlist)
			var n sync.WaitGroup
			for r,_:=range relationlist{
				n.Add(1)
				go func(){
					defer n.Done()
				tag := models.Tag{}
					tid:=relationlist[r].Tid
					tag.Id = tid
					if err:=tag.Read();err==nil{
						a[list[k].Id] = append(a[list[k].Id],tag.Name)
					}

				}()

			}
			n.Wait()
		}(k)
	}
	nn.Wait()
	this.Data["count"] = count
	this.Data["paginator"] = models.Paginator(page,PageSize,count)
	this.Data["list"] = list
	this.Data["a"] = a
	this.TplName = "index.html"
}

func (this *IndexController) Detail(){

	id,_:=this.GetInt("id")
	var article models.Article
	article.Id = id
	if err:=article.Read();err!=nil{
		this.Abort("404")
	}
	//article.Content = RenderMarkdown(article.Content)
	this.Layout = "layout.html"
	this.Data["article"] = article
	this.TplName = "detail.html"
}

func (this *IndexController) Axis()  {
	this.TplName = "axis.html"
}

