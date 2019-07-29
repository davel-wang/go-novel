package models

import (
	"github.com/astaxie/beego/orm"
)

type Page struct {
	PageNo       int
	PageSize     int
	TotalPage    int
	TotalCount   int
	PreviousPage int
	NextPage     int
	FirstPage    bool //是否是第一页
	LastPage     bool //是否是最后一页
	List         []orm.Params
}

func (this *Page) PageUtil(count int, pageNo int, pageSize int, list []orm.Params) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	FirstPage := pageNo == 1
	LastPage := pageNo == tp

	PreviousPage := 0
	NextPage := 0
	if !FirstPage {
		PreviousPage = pageNo - 1
	}

	if !LastPage {
		NextPage = pageNo + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: FirstPage, LastPage: LastPage, PreviousPage: PreviousPage, NextPage: NextPage, List: list}
}
