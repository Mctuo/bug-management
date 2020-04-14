package user

import (
	"bug-management/models"
	user "bug-management/models/user"
	. "common/logs"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type PersonController struct {
	beego.Controller
}

var(
	myAddr = "123.56.25.57/"
)


func(c *PersonController)UploadAvatar(){
	f,h,err:=c.GetFile("file")
	if err != nil{
		Error("PersonController UploadAvatar arg error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	defer f.Close()
	myFileName := fmt.Sprintf("%d-%s",time.Now().UnixNano()/1e6,h.Filename)

	err = c.SaveToFile("file","/usr/share/nginx/static/img/"+myFileName)
	if err != nil{
		Error("PersonController UploadAvatar error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	avatarSrc :=myAddr + myFileName
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),avatarSrc,c.Ctx)
	return
}

func (c *PersonController)SetUserInfo(){
	models.PrintClientInfo(c.Ctx)

	var myReq user.StruPersonInfoReq
	err := json.Unmarshal(c.Ctx.Input.RequestBody,&myReq)
	if err != nil{
		Error("PersonController SetUserInfo json.Unmarshal error:",err.Error())
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	if myReq.Account <=0{
		Error("PersonController SetUserInfo myReq.Account=",myReq.Account)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,"账号不合法"),nil,c.Ctx)
		return
	}
	err = user.SetUserInfo2DB(myReq)
	if err != nil{
		Error("PersonController SetUserInfo user.SetUserInfo2DB error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),nil,c.Ctx)
	return
}

func (c *PersonController)GetUserInfo(){
	models.PrintClientInfo(c.Ctx)

	myAccount,err := c.GetInt64("account")
	if err != nil{
		Error("PersonController GetUserInfo account=",myAccount)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	if myAccount <=0{
		Error("PersonController GetUserInfo account=",myAccount)
		models.HandleError(models.ErrArg,models.GetErrMsg(models.ErrArg,err.Error()),nil,c.Ctx)
		return
	}
	var myResp user.StruPersonInfoReq
	err = user.GetUserInfo(myAccount,&myResp)
	if err != nil{
		Error("PersonController GetUserInfo user.GetUserInfo error:",err.Error())
		models.HandleError(models.ErrSvr,models.GetErrMsg(models.ErrSvr,err.Error()),nil,c.Ctx)
		return
	}
	models.HandleError(models.Success,models.GetErrMsg(models.Success,""),myResp,c.Ctx)
	return
}

func(c *PersonController)GetUserList(){
	models.PrintClientInfo(c.Ctx)


}