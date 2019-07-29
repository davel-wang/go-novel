package routers

import (
	"novel/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/index/:page:int", &controllers.MainController{}, "*:Index")                    //书籍列表
	beego.Router("/book/:book_id:int/:page:int", &controllers.MainController{}, "*:BooksIist")    //书籍章节列表
	beego.Router("/book/:book_id:int/info/:id:int", &controllers.MainController{}, "*:BooksInfo") //书籍章节详情
}
