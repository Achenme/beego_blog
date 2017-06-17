package models

import "github.com/astaxie/beego/orm"

type Relation struct {
	Id int64
	Tid int
	Aid int
	Addtime int64
}


func (m *Relation) TableName() string{
	return TableName("relation")
}

func (m *Relation) Insert() error  {
	if _,err:=orm.NewOrm().Insert(m);err!=nil{
		return err
	}
	return nil
}

func (m *Relation) Read(fields ...string) error {
	if err:=orm.NewOrm().Read(m,fields...);err!=nil{
		return err
	}
	return nil
}

func (m *Relation)  Update(fields ...string) error{
	if _,err:=orm.NewOrm().Update(m,fields...);err!=nil{
		return err
	}
	return nil
}

func (m *Relation) Delete() error  {
	if _,err:=orm.NewOrm().Delete(m);err!=nil{
		return err
	}
	return nil
}

func (m *Relation) Query() orm.QuerySeter  {
	return orm.NewOrm().QueryTable(m)
}