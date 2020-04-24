package routers

import (
	"bug-management/controllers/bug"
	_case "bug-management/controllers/case"
	"bug-management/controllers/health"
	"bug-management/controllers/project"
	"bug-management/controllers/result"
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
	//1.用户user
	beego.Router("/bug/api/uploadavator",&user.PersonController{},"POST:UploadAvatar")

    beego.Router("/bug/api/setusrinfo",&user.PersonController{},"POST:SetUserInfo")

	beego.Router("/bug/api/getuserinfo",&user.PersonController{},"GET:GetUserInfo")

	beego.Router("/bug/api/getuserlist",&user.PersonController{},"GET:GetUserList")

	beego.Router("/bug/api/user/add",&user.PersonController{},"POST:AddUser")

	beego.Router("/bug/api/user/delete",&user.PersonController{},"POST:DeleteUser")

	//2.项目project
	beego.Router("/bug/api/project/create",&project.ProjectController{},"POST:CreateProject")

	beego.Router("/bug/api/project/list",&project.ProjectController{},"GET:ProjectList")

	beego.Router("/bug/api/project/info",&project.ProjectController{},"GET:ProjectInfo")

	beego.Router("/bug/api/project/delete",&project.ProjectController{},"POST:DeleteProject")

	//3.用例case
	beego.Router("/bug/api/case/create",&_case.CaseController{},"POST:CreateCase")

	beego.Router("bug/api/case/assign/list",&_case.CaseController{},"GET:AssignList")

	beego.Router("/bug/api/case/listbycaseid",&_case.CaseController{},"GET:GetCaseId")

	//4.result
	beego.Router("/bug/api/result/create",&result.ResultController{},"POST:CreateResult")

	beego.Router("/bug/api/result/listbyassign",&result.ResultController{},"GET:ResultList")

	beego.Router("/bug/api/result/listbycaseid",&result.ResultController{},"GET:ResultListByCaseId")

	//bug_info
	beego.Router("/bug/api/buginfo/create",&bug.BugController{},"POST:BugCreate")

	beego.Router("/bug/api/buginfo/listbytitle",&bug.BugController{},"GET:BugInfoByTitle")

	beego.Router("/bug/api/buginfo/listbyassign",&bug.BugController{},"GET:AssignBug")


	//bug_solution
	beego.Router("/bug/api/bugsolve/create",&bug.BugController{},"POST:CreateBugSolution")

	beego.Router("/bug/api/bugsolve/list",&bug.BugController{},"GET:GetSolutionList")
}
