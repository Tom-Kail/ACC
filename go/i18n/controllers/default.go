package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	beego.AddFuncMap("i18n", i18n.Tr)
}
