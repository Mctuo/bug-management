package user

import (
	"bug-management/conf"
	"bug-management/database"
	. "common/logs"
)

type StruPersonInfoReq struct {
	Name 	string	`json:"name"`
	Account int64	`json:"account"`
	Mail    string	`json:"mail"`
	Job    string	`json:"job"`
	Note   string	`json:"note"`
	Avatar string	`json:"avatar"`
}

var(
	StruUserInfo=[]string{
		"name",
		"account",
		"mail",
		"job",
		"note",
		"avatar",
	}
	StruUpdateInfo=[]string{
		"name",
		"mail",
		"job",
		"note",
		"avatar",
	}
)

func SetUserInfo2DB(myReq StruPersonInfoReq)error{
	isExist := database.HasItem(conf.MyUserTb,"account",myReq.Account)
	Info("isExist is",isExist)
	if isExist == true{
		err := database.Update(conf.MyUserTb,StruUpdateInfo,"account",myReq.Name,myReq.Mail,myReq.Job,myReq.Note,myReq.Avatar,myReq.Account)
		if err != nil{
			Error("SetUserInfo2DB database.Update error:",err.Error())
			return err
		}
		return nil
	}

	err := database.Insert(conf.MyUserTb,StruUserInfo,myReq.Name,myReq.Account,myReq.Mail,myReq.Job,myReq.Note,myReq.Avatar)
	if err != nil{
		Error("SetUserInfo2DB database.Insert error:",err.Error())
		return err
	}
	return nil
}

func GetUserInfo(account int64, resp *StruPersonInfoReq)error{
	err := database.Query(conf.MyUserTb,"account",account,StruUserInfo,
		&resp.Name,&resp.Account,&resp.Mail,&resp.Job,&resp.Note,&resp.Avatar)
	if err != nil{
		Error("GetUserInfo database get error:",err.Error())
		return err
	}
	return nil
}