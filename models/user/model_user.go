package user

import (
	"bug-management/conf"
	"bug-management/database"
	. "common/logs"
	"errors"
)

type StruPersonInfoReq struct {
	Name 	string	`json:"name"`
	Account int64	`json:"account"`
	Mail    string	`json:"mail"`
	Job    string	`json:"job"`
	Note   string	`json:"note"`
	Avatar string	`json:"avatar"`
}

type StruPersonInfoResp struct {
	Info []StruPersonInfoReq `json:"info"`
}

type StruAddUserReq struct {
	ProjectId  int64	`json:"projectId"`
	Account    []int64	`json:"account"`
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

func GetUserList(project int64,myResp *StruPersonInfoResp)error{
	Rows,err := database.GetDB().Query("select *from person_info where account IN (select account from project_people where id=?)",project)
	if err != nil{
		Error("GetUserList Query error:",err.Error())
		return err
	}
	defer Rows.Close()
	var id int
	var tmp  StruPersonInfoReq
	for Rows.Next() {
		err = Rows.Scan(&id, &tmp.Name, &tmp.Account, &tmp.Mail, &tmp.Job, &tmp.Note, &tmp.Avatar)
		if err != nil {
			Error("GetUserList Rows.Scan error:", err.Error())
			return err
		}
		myResp.Info = append(myResp.Info,tmp)
	}
	return nil
}



func AddUser(myReq StruAddUserReq)error{
	isExist := database.HasItem(conf.MyProjectTb,"id",myReq.ProjectId)
	if !isExist{
		return errors.New("项目不存在")
	}
	insertSql, err := database.GetDB().Prepare(`insert into project_people(id,member_num)values(?,?)`)
	if err != nil{
		Error("AddUser database.GetDB().Prepare error:",err.Error())
		return err
	}

	tx,err := database.GetDB().Begin()
	defer func() {
		if err != nil{
			err1 := tx.Rollback()
			if err1 != nil{
				Error("AddUser Rollback error:",err1.Error())
			}
		}
	}()
	for _,v := range myReq.Account{
		_,err = tx.Stmt(insertSql).Exec(myReq.ProjectId,v)
		if err != nil{
			Error("AddUser sql Exec error:",err.Error())
			return err
		}
	}

	peopleNum := len(myReq.Account)
	_,err = database.GetDB().Exec("update project set project_people = project_people + ? where id=?",peopleNum,myReq.ProjectId)
	if err != nil{
		Error("AddUser Exec2 error:",err.Error())
		return err
	}

	err  = tx.Commit()
	return err
}

func DeleteUser(myReq StruAddUserReq)error{
	isExist := database.HasItem(conf.MyProjectTb,"id",myReq.ProjectId)
	if !isExist{
		return errors.New("项目不存在")
	}
	deleteSql,err := database.GetDB().Prepare("delete from project_people where id=? and member_num=?")
	if err != nil{
		Error("DeleteUser database.GetDB().Prepare error:",err.Error())
		return err
	}
	tx,err := database.GetDB().Begin()
	defer func() {
		if err != nil{
			err1 := tx.Rollback()
			if err1 != nil{
				Error("DeleteUser Rollback error:",err1.Error())
			}
		}
	}()
	for _,v := range myReq.Account{
		_,err = tx.Stmt(deleteSql).Exec(myReq.ProjectId,v)
		if err != nil{
			Error("DeleteUser tx.Stmt Exec error:",err.Error())
			return err
		}
	}

	peopleNum := len(myReq.Account)
	_,err = database.GetDB().Exec("update project set project_people = project_people - ? where id=?",peopleNum,myReq.ProjectId)
	if err != nil{
		Error("AddUser Exec2 error:",err.Error())
		return err
	}
	err = tx.Commit()
	return err
}