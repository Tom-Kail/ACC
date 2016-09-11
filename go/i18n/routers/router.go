package routers

import (
	"github.com/Tom-Kail/acc/go/i18n/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
