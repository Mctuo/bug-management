package routers

import (
	"bug-management/controllers/health"
	"bug-management/controllers/project"
	"bug-management/controllers/user"
	"github.com/astaxie/beego/plugins/cors"
	sso"bug-management/filter"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//	AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		//	ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowHeaders:  []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "IVC-Auth"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},

		AllowCredentials: true,
	}))
	beego.InsertFilter("/bug/api/*",beego.BeforeRouter,sso.SsoFilter())

    beego.Router("/health", &health.HealthController{},"GET:Health")

    beego.Router("/bug/api/getnowtimes",&health.HealthController{},"GET:GetNowTimes")

	beego.Router("/bug/api/uploadavator",&user.PersonController{},"POST:UploadAvatar")

    beego.Router("/bug/api/setusrinfo",&user.PersonController{},"POST:SetUserInfo")

	beego.Router("/bug/api/getuserinfo",&user.PersonController{},"GET:GetUserInfo")

	beego.Router("/bug/api/getuserlist",&user.PersonController{},"GET:GetUserList")

	beego.Router("/bug/api/project/create",&project.ProjectController{},"POST:CreateProject")

	beego.Router("/bug/api/project/list",&project.ProjectController{},"GET:ProjectList")
}
