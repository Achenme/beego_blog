package models

import(
	"github.com/astaxie/beego/orm"
)


type Article struct {
	Id	int `form:"-"`
	Name 	string `form:"name" valid:"Required"`
	Content	string `form:"content"`
	Addtime int64
}


func (m *Article) TableName() string  {
	return TableName("Article")
}

func (m *Article) Insert() error  {
	if _,err:=orm.NewOrm().Insert(m);err!=nil{
		return err
	}
	return nil
}

func (m *Article)  Read(fields ...string) error {
	if err:=orm.NewOrm().Read(m,fields...);err!=nil{
		return err
	}
	return nil
}

func (m *Article) Update(fields ...string) error  {
	if _,err:=orm.NewOrm().Update(m,fields...);err!=nil{
		return err
	}
	return nil
}

func (m *Article) Delete() error  {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}