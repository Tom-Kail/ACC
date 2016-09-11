package routers

import (
	"github.com/Tom-Kail/acc/go/beego_test/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
