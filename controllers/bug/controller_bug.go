package bug

import (
	"bug-management/models"
	"bug-management/models/bug"
	"encoding/json"
	"github.com/astaxie/beego"
	."common/logs"

)

type BugController struct {
	beego.Controller
}

func(c *BugController)BugCreate(){
	models.PrintClientInfo(c.Ctx)

	var myReq bug.StruCreateInfoReq
	err := json.Unmarshal(c.Ctx.Input.RequestBody,&myReq)
	if err != nil{
		Error("BugCreate json.Unmarshal error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg, err.Error()),nil,c.Ctx)
		return
	}

	if myReq.Assigned < 1{
		Error("BugCreate myReq.Assigned=",myReq.Assigned)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg, "assign error"),nil,c.Ctx)
		return
	}

	if myReq.BugTitle == ""{
		Error("BugCreate  myReq.BugTitle is null")
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg, "bugTitle error"),nil,c.Ctx)
		return
	}

	err = bug.CreateBugInfo(myReq)
	if err != nil{
		Error("BugController BugCreate  bug.CreateBugInfo error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,"ok"),nil,c.Ctx)
	return
}


func (c *BugController)BugInfoByTitle(){
	models.PrintClientInfo(c.Ctx)

	myTitle := c.GetString("bug_title")

	var myResp bug.StruBugInfoResp
	err := bug.GetBugInfo(myTitle,&myResp)
	if err != nil{
		Error("BugController BugInfo  bug.BugInfo error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,"ok"),myResp.Info,c.Ctx)
	return
}


func(c *BugController)AssignBug(){
	models.PrintClientInfo(c.Ctx)

	assigned,err := c.GetInt64("assigned")
	if err != nil{
		Error("AssignBug assigned error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg, err.Error()),nil,c.Ctx)
		return
	}
	if assigned < 1{
		Error("AssignBug assigned error,assigned=",assigned)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"assigned error"),nil,c.Ctx)
		return
	}
	var myResp bug.StruBugInfoRespData
	err = bug.GetBugInfoByAssign(assigned,&myResp)
	if err != nil{
		Error("BugController BugInfo  bug.BugInfo error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,"ok"),myResp,c.Ctx)
	return

}

func (c *BugController)CreateBugSolution(){
	models.PrintClientInfo(c.Ctx)

	var myReq bug.StruBugSolutionReq
	err := json.Unmarshal(c.Ctx.Input.RequestBody,&myReq)
	if err != nil{
		Error("CreateBugSolution json.Unmarshal error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg, err.Error()),nil,c.Ctx)
		return
	}

	err = bug.CreateBugSolution(myReq)
	if err != nil{
		Error("CreateBugSolution BugInfo  bug.CreateBugSolution error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,"ok"),nil,c.Ctx)
	return
}

func (c *BugController)GetSolutionList(){
	models.PrintClientInfo(c.Ctx)

	myProjectId ,err := c.GetInt64("projectId",0)
	if err != nil{
		Error("GetSolutionList myProjectId error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	if myProjectId < 0{
		Error("GetSolutionList myProjectId=",myProjectId)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"invaild projectId"),nil,c.Ctx)
		return
	}

	var myResp bug.StruBugSolutionResp
	err = bug.GetSolutionList(myProjectId,&myResp)
	if err != nil{
		Error("BugController BugInfo  bug.BugInfo error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,"ok"),myResp.Info,c.Ctx)
	return
}