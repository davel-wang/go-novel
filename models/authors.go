package models

import (
	"github.com/astaxie/beego/orm"
)

type Authors struct {
	Id    int
	name  string
	Books *Books `orm:"rel(fk)"` //设置一对多关系
}

func init() {
	orm.RegisterModel(new(Authors))
}
