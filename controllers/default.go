package controllers

import (
	// "fmt"
	"novel/models"

	//"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

type MainController struct {
	baseController
}

func (this *MainController) Index() {
	page, page_size := this.getPage()
	var books models.Books
	list := books.ListPage(page, page_size)
	this.Data["page"] = list
	this.Data["title"] = beego.AppConfig.String("appname")
	this.TplName = "index/index.html"
}

//书籍章节列表
func (this *MainController) BooksIist() {
	page, page_size := this.getPage()
	var booksList models.BooksList

	book_id, err := this.GetInt(":book_id")
	if err != nil {
		book_id = 0
	}

	list := booksList.ListPage(book_id, page, page_size)
	this.Data["page"] = list
	this.Data["title"] = beego.AppConfig.String("appname")
	this.TplName = "book/list.html"
}

//书籍章节详情
func (this *MainController) BooksInfo() {
	book_id, err := this.GetInt(":book_id")
	if err != nil {
		book_id = 0
	}
	var books models.Books
	book := books.ReadBook(book_id)

	books_list_id, err := this.GetInt(":id")
	if err != nil {
		books_list_id = 0
	}

	cache_key := "book_list_" + string(books_list_id)
	filecache, err := cache.NewCache("file", `{}`)

	content := filecache.Get(cache_key)

	this.Data["book"] = book
	this.Data["content"] = content
	this.Data["title"] = beego.AppConfig.String("appname")
	this.TplName = "book/info.html"
}
