package models

import (
	"github.com/astaxie/beego"
	"fmt"
	_"github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/orm"
)


func init()  {
	orm.Debug = true
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "a.db")
	orm.RegisterModel(new(Article), new(Tag),new(Relation))

}

func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}