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
	err = pro.CreateProject(myReq)
	if err != nil{
		Error("ProjectController CreateProject pro.CreateProject error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),nil,c.Ctx)
	return
}