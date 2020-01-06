package routers

import (
	"bug-management/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HeathController{},"GET:Health")
}
