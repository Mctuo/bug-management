package project

import (
	"bug-management/models"
	pro "bug-management/models/project"
	. "common/logs"
	"encoding/json"
	"github.com/astaxie/beego"
)

type ProjectController struct {
	beego.Controller
}

func (c *ProjectController)ProjectInfo(){
	models.PrintClientInfo(c.Ctx)

	//myAccount,err := c.GetInt64("account")
	//if err != nil{
	//	Error("PersonController ProjectInfo myAccount error:",err.Error())
	//	models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
	//	return
	//}
	//if myAccount <= 0{
	//	Error("ProjectController ProjectInfo error:account=",myAccount)
	//	models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"账号错误"),nil,c.Ctx)
	//	return
	//}
	myProjectId,err := c.GetInt64("projectId")
	if err != nil{
		Error("PersonController ProjectInfo myProjectId error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	if myProjectId < 1{
		Error("PersonController ProjectInfo  error,myProjectId=",myProjectId)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"非法projectId"),nil,c.Ctx)
		return
	}
	var myResp pro.StruProject

	err = pro.ProjectDesc(myProjectId,&myResp)
	if err != nil{
		Error("PersonController project.ProjectList error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),myResp,c.Ctx)
	return
}


func (c *ProjectController)ProjectList(){
	models.PrintClientInfo(c.Ctx)

	myAccount,err := c.GetInt64("account")
	if err != nil{
		Error("PersonController ProjectList myAccount error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	if myAccount <= 0{
		Error("ProjectController ProjectList error:account=",myAccount)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"账号错误"),nil,c.Ctx)
		return
	}
	var myResp pro.StruProjectListResp
	err = pro.ProjectList(myAccount,&myResp)
	if err != nil{
		Error("PersonController project.ProjectList error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),myResp.ListInfo,c.Ctx)
	return
}


func (c *ProjectController)CreateProject(){
	models.PrintClientInfo(c.Ctx)

	var myReq  pro.StruCreateProjectReq
	err := json.Unmarshal([]byte(c.Ctx.Input.RequestBody),&myReq)
	if err != nil{
		Error("ProjectController CreateProject json.Unmarshal error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	if myReq.Account <= 0{
		Error("ProjectController CreateProject error:account=",myReq.Account)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"账号错误"),nil,c.Ctx)
		return
	}
	if myReq.ProjectName ==""{
		Error("ProjectController CreateProject error:ProjectName=",myReq.ProjectName)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"项目名称为空"),nil,c.Ctx)
		return
	}
	if myReq.Creater < 1{
		Error("ProjectController CreateProject myReq.Creater =",myReq.Creater)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"创建人错误"),nil,c.Ctx)
		return
	}
	err = pro.CreateProject(myReq)
	if err != nil{
		Error("ProjectController CreateProject pro.CreateProject error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),nil,c.Ctx)
	return
}

func (c *ProjectController)DeleteProject(){
	models.PrintClientInfo(c.Ctx)

	var myReq pro.StruDeleteProjectReq
	err := json.Unmarshal([]byte(c.Ctx.Input.RequestBody),&myReq)
	if err != nil{
		Error("ProjectController DeleteProject json.Unmarshal error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}

	if myReq.Creater <= 0{
		Error("ProjectController DeleteProject error:creater=",myReq.Creater)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"创建人错误"),nil,c.Ctx)
		return
	}
	if myReq.ProjectId < 1{
		Error("ProjectController DeleteProject error:projectId=",myReq.ProjectId)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"项目Id错误"),nil,c.Ctx)
		return
	}

	err = pro.ProjectCreater(myReq.Creater)
	if err != nil{
		Error("ProjectController DeleteProject pro.ProjectCreater error:",err.Error())
		models.HandleError(models.ErrPerson,models.GetErrMsg(models.ErrPerson,err.Error()),nil,c.Ctx)
		return
	}


}