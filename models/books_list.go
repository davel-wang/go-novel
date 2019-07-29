package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

type BooksList struct {
	Id       int
	Title    string
	Add_time int
}

func init() {
	orm.RegisterModel(new(BooksList))
}

func (this *BooksList) ListPage(book_id int, page int, numPerPage int) Page {
	var pages Page

	sql1 := "select * from books_list where book_id=" + fmt.Sprintf("%d", book_id) + " order by id desc limit ?," + fmt.Sprintf("%d", numPerPage)
	sql2 := "select count(*) as number from books_list where book_id=" + fmt.Sprintf("%d", book_id)
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

func (this *BooksList) ReadByTitle(title string) BooksList {
	sql1 := "select * from books_list where title='" + title + "' limit 1"
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
