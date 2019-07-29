package main

import (
	_ "novel/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dsn := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@tcp(" + beego.AppConfig.String("mysqlurls") + ")/" + beego.AppConfig.String("mysqldb") + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dsn)

	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()
}
