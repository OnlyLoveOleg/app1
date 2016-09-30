package routers

import (
	"github.com/TalLannder/app1/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
