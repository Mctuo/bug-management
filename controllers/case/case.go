package _case

import (
	"bug-management/models"
	_case "bug-management/models/case"
	. "common/logs"
	"encoding/json"
	"github.com/astaxie/beego"
)

type CaseController struct {
	beego.Controller
}

func(c *CaseController)CreateCase(){
	models.PrintClientInfo(c.Ctx)

	var myReq _case.StruCreateCaseReq
	err := json.Unmarshal([]byte(c.Ctx.Input.RequestBody),&myReq)
	if err != nil{
		Error("CaseController CreateCase json.Unmarshal error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}

	if myReq.ProjectId < 1{
		Error("CaseController error,ProjectId= ",myReq.ProjectId)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"projectId错误"),nil,c.Ctx)
		return
	}
	if myReq.Creator < 1{
		Error("CaseController error,creater= ",myReq.Creator)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"creator错误"),nil,c.Ctx)
		return
	}

	if myReq.Title == ""{
		Error("CaseController error,title is null")
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"title为空"),nil,c.Ctx)
		return
	}
	err = _case.CreateCase(myReq)
	if err != nil{
		Error("CaseController _case.CreateCase error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,"ok"),nil,c.Ctx)
	return
}

func (c *CaseController)AssignList(){
	models.PrintClientInfo(c.Ctx)

	myAssign,err := c.GetInt64("assign")
	if err != nil{
		Error("CaseController AssignList error",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	if myAssign < 1{
		Error("CaseController AssignList error ,myAssign= ",myAssign)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"Assign错误"),nil,c.Ctx)
		return
	}
	var myResp _case.StruAssignResp
	err = _case.AssignList(myAssign,&myResp)
	if err != nil{
		Error("CaseController _case.AssignList error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,"ok"),myResp.ListInfo,c.Ctx)
	return
}


func (c *CaseController)GetCaseId(){
	models.PrintClientInfo(c.Ctx)

	CaseId,err := c.GetInt64("caseId")
	if err != nil{
		Error("GetCaseId CaseId error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	if CaseId < 1 {
		Error("GetCaseId CaseId error,CaseId=",CaseId)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"caseId error"),nil,c.Ctx)
		return
	}

	var myResp _case.StruCreateCaseReq
	err = _case.ListByCaseId(CaseId,&myResp)
	if err != nil{
		Error("CaseController  _case.ListByCaseId error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,"ok"),myResp,c.Ctx)
	return
}