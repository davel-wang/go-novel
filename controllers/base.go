package controllers

import (
	"strings"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func init() {
	langTypes := strings.Split(beego.AppConfig.String("lang_types"), "|")
	for _, lang := range langTypes {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

type baseController struct {
	beego.Controller
	i18n.Locale
}

func (this *baseController) Prepare() {
	this.Lang = ""
	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5]
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}
	if len(this.Lang) == 0 {
		this.Lang = "en-US"
	}
	this.Data["Lang"] = this.Lang
}

func (this *baseController) getPage() (int, int) {
	page_size, err := beego.AppConfig.Int("page_size")
	if err != nil {
		page_size = 10
	}
	page_max, err := beego.AppConfig.Int("page_max")
	if err != nil {
		page_max = 1000
	}

	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	} else {
		if page < 1 {
			page = 1
		}

		if page > page_max {
			page = page_max
		}
	}
	return page, page_size
}
