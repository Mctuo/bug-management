package result

import (
	."common/logs"
	"bug-management/models"
	"bug-management/models/result"
	"encoding/json"
	"github.com/astaxie/beego"
)

type ResultController struct {
	beego.Controller
}

func (c *ResultController)CreateResult(){
	models.PrintClientInfo(c.Ctx)

	var myReq result.StruResultReq
	err := json.Unmarshal([]byte(c.Ctx.Input.RequestBody),&myReq)
	if err != nil{
		Error("ResultController CreateResult json.Unmarshal error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}

	if myReq.ProjectId < 1{
		Error("ResultController CreateResult myReq.ProjectId=",myReq.ProjectId)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"projectId 错误"),nil,c.Ctx)
		return
	}
	if myReq.Assigned < 1{
		Error("ResultController CreateResult myReq.Assigned=",myReq.Assigned)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"assigned 错误"),nil,c.Ctx)
		return
	}

	if myReq.CaseId < 1{
		Error("ResultController CreateResult myReq.CaseId error:",myReq.CaseId)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"caseId 错误"),nil,c.Ctx)
		return
	}

	err = result.CreateResult(myReq)
	if err != nil{
		Error("ResultController CreateResult  result.CreateResult error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),nil,c.Ctx)
	return
}


func (c *ResultController)ResultList(){
	models.PrintClientInfo(c.Ctx)

	myAssign,err := c.GetInt64("assign")
	if err != nil{
		Error("ResultController ResultList myAssign error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	if myAssign < 1 {
		Error("ResultController ResultList myAssign=",myAssign)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"myAssign error"),nil,c.Ctx)
		return
	}
	var myResp result.StruResultListResp

	err = result.ResultList(myAssign,&myResp)
	if err != nil{
		Error("ResultController ResultList  result.ResultList error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),myResp.Info,c.Ctx)
	return
}

func (c *ResultController)ResultListByCaseId(){
	models.PrintClientInfo(c.Ctx)

	myCaseId ,err := c.GetInt64("caseId")
	if err != nil{
		Error("ResultListByCaseId myCaseId error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	if myCaseId < 1{
		Error("ResultListByCaseId myCaseId error,myCaseId=",myCaseId)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"caseId error"),nil,c.Ctx)
		return
	}

	var myResp result.StruResultResp
	err = result.ResultByCaseId(myCaseId,&myResp)
	if err != nil{
		Error("ResultController ResultListByCaseId   result.ResultByCaseId error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),myResp,c.Ctx)
	return
}