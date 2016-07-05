// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"webscan_bate/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/verifypic",
			beego.NSInclude(
				&controllers.VerifyCodeController{},
			),
		),
		beego.NSNamespace("/tools",
			beego.NSInclude(
				&controllers.ToolsController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/manage",
			beego.NSInclude(
				&controllers.ManageController{},
			),
		),
		beego.NSNamespace("/policy",
			beego.NSInclude(
				&controllers.PolicyController{},
			),
		),
		beego.NSNamespace("/site",
			beego.NSInclude(
				&controllers.SiteController{},
			),
		),
		beego.NSNamespace("/task",
			beego.NSInclude(
				&controllers.TaskController{},
			),
		),
		beego.NSNamespace("/report",
			beego.NSInclude(
				&controllers.ReportController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
