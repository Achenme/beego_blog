package admin

import(

	"github.com/astaxie/beego"
	"blog/models"
	"github.com/astaxie/beego/validation"
	"time"
)


type CatController struct {
	BaseController
}


func (this *CatController) Index() {
	var tag models.Tag
	var list []*models.Tag
	query:=tag.Query()
	count,_:=query.Count()
	query.All(&list)
	Pagbar:=models.Paginator(1,1,count)
	this.Data["Pagbar"] = Pagbar
	this.Data["list"] = list

	this.TplName = "admin/cat/index.html"
}

func (this *CatController) Add() {
	beego.ReadFromRequest(&this.Controller)
	this.TplName = "admin/cat/modify.html"
}

func (this *CatController) Edit() {
	beego.ReadFromRequest(&this.Controller)
	id,_:= this.GetInt("id")
	var tag models.Tag
	tag.Id = id
	if err:=tag.Read();err!=nil{
		this.Abort("404")
	}
	this.Data["tag"] = tag
	this.TplName = "admin/cat/modify.html"

}

func (this *CatController) Save() {
	var tag models.Tag
	valid := validation.Validation{}
	this.ParseForm(&tag)

	flash:=beego.NewFlash()
	b, err :=valid.Valid(&tag)
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
		tag.Addtime = time.Now().Unix()
		err:=tag.Insert()
		if err!=nil{
			flash.Error("添加失败")
		}else{
			flash.Success("添加成功")
		}
	}else{
		tag.Id = id
		if err := tag.Read();err!=nil{
			flash.Error("不存在该文章")
		}else{
			this.ParseForm(&tag)
			tag.Id = id
			tag.Addtime = time.Now().Unix()
			err:=tag.Update()
			if err!=nil{
				flash.Error("编辑失败")
			}else{
				flash.Success("编辑成功")
			}
		}

	}

	flash.Store(&this.Controller)
	this.Redirect("/cat/add", 302)
}
func (this *CatController) Delete()  {
	id,_:= this.GetInt("id")
	var tag models.Tag
	tag.Id = id
	if err:=tag.Read();err==nil{
		tag.Delete()
	}

	this.ServeJSON()
}