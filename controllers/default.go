package controllers

import (
	// "fmt"
	"mybee/models"

	//"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego"
)

type MainController struct {
	baseController
}

func (this *MainController) Get() {
	// qb, _ := orm.NewQueryBuilder("mysql")
	// // 构建查询对象
	// books := qb.Select("*").From("books").Limit(10).Offset(0)
	// fmt.Println(books)

	// list := []models.Books{}
	// db := orm.NewOrm()
	// _, _ = db.QueryTable("books").RelatedSel().OrderBy("-id").Limit(10, 0).All(&list)

	page, page_size := this.getPage()
	var books models.Books
	list := books.ListPage(page, page_size)

	this.Data["page"] = list
	this.Data["title"] = beego.AppConfig.String("appname")
	this.TplName = "index/index.html"
}
