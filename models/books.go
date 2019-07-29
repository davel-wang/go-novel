package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Books struct {
	Id        int
	Author_id int
	Name      string
	From      string
	Views     int
}

func (this *Books) ListPage(page int, numPerPage int) Page {
	var pages Page

	sql1 := "select * from books order by views desc limit ?," + fmt.Sprintf("%d", numPerPage)
	sql2 := "select count(*) as number from books"
	var maps, maps2 []orm.Params
	o := orm.NewOrm()
	_, err := o.Raw(sql1, numPerPage*(page-1)).Values(&maps)
	if err != nil {
		return pages
	}

	_, err = o.Raw(sql2).Values(&maps2)
	if err != nil {
		return pages
	}

	number, err := strconv.Atoi(maps2[0]["number"].(string))
	return pages.PageUtil(number, page, numPerPage, maps)
}

func (this *Books) ReadBook(book_id int) Books {
	o := orm.NewOrm()
	books := Books{Id: book_id}

	_ = o.Read(&books)
	return books
}

func init() {
	orm.RegisterModel(new(Books))
}
