package models

import (
	"github.com/astaxie/beego/orm"
)
type BooksList struct {
	Id    int
	Title string
}

func init() {
	orm.RegisterModel(new(BooksList))
}