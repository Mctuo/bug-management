package project

import (
	"bug-management/conf"
	"bug-management/database"
	. "common/logs"
	"errors"
)

var (
	ProjectInfo =[]string{
		"project_name",
		"account",
		"project_people",
		"task_total",
		"task_finished",
		"task_unfinished",
		"creater",
	}
)


type StruDeleteProjectReq struct {
	ProjectId int64	`json:"projectId"`
	Creater   int64	`json:"creater"`
}

type StruCreateProjectReq struct {
	ProjectName string `json:"project_name"`
	Account int64	`json:"account"`
	Creater int64	`json:"creater"`
}

type StruProject struct {
	ProjectId int64	`json:"projectId"`
	ProjectName string `json:"project_name"`
	Account int64	`json:"account"`
	ProjectPeople int	`json:"project_people"`
	TotalTask int	`json:"total_task"`
	TaskFinished int	`json:"task_finished"`
	TaskUnFinished int	`json:"task_unfinished"`
	Creater int64	`json:"creater"`
}

type StruProjectListResp struct {
	ListInfo []StruProject `json:"listInfo"`
}

func ProjectDesc(projectId int64, myResp *StruProject)error{
	sqlStr := "select* from project where  id= ?"
	Info("ProjectDesc Query sqlStr=",sqlStr)
	Info("ProjectDesc Query value=",projectId)
	rows,err := database.GetDB().Query(sqlStr,projectId)
	if err != nil{
		Error("ProjectDesc database.GetDB().Query error:",err.Error())
		return err
	}
	defer rows.Close()
	for rows.Next(){
		err = rows.Scan(&myResp.ProjectId,&myResp.ProjectName,&myResp.Account,&myResp.ProjectPeople,&myResp.TotalTask,&myResp.TaskUnFinished,&myResp.TaskFinished,&myResp.Creater)
		if err != nil{
			Error("ProjectDesc rows.Scan error:",err.Error())
			return err
		}
	}
	return nil
}

func CreateProject(myReq StruCreateProjectReq)error{
	IsExist := database.HasItem(conf.MyProjectTb,"project_name",myReq.ProjectName)
	if IsExist{
		return errors.New("项目已创建，不能重复创建")
	}

	tx,err := database.GetDB().Begin()

	err = database.Insert(conf.MyProjectTb,ProjectInfo,myReq.ProjectName,myReq.Account,1,0,0,0,myReq.Account)
	if err != nil{
		tx.Rollback()
		Error("CreateProject database.Insert error:",err.Error())
		return err
	}

	execSql := "insert into project_people(id,account) select id,account from project where project_name =?"

	_,err = database.GetDB().Exec(execSql,myReq.ProjectName)
	if err != nil{
		tx.Rollback()
		Error("models CreateProject database.GetDB().Exec error:",err.Error())
		return err
	}
	tx.Commit()
	return nil
}


func ProjectList(account int64,myResp *StruProjectListResp)error{
	var tmp StruProject

	rows,err :=database.GetDB().Query("select *from project where account =?",account)
	if err != nil{
		Error("ProjectList database.GetDB().Query error:",err.Error())
		return err
	}
	defer rows.Close()

	for rows.Next(){
		err = rows.Scan(&tmp.ProjectId,&tmp.ProjectName,&tmp.Account,&tmp.ProjectPeople,&tmp.TotalTask,&tmp.TaskUnFinished,&tmp.TaskFinished,&tmp.Creater)
		if err != nil{
			Error("ProjectList rows.Scan error:",err.Error())
			return err
		}
		myResp.ListInfo = append(myResp.ListInfo,tmp)
	}
	return nil
}


func ProjectCreater(creater int64)error{
	rows,err :=database.GetDB().Query("select creater from project where id =?",creater)
	if err != nil{
		Error("ProjectDelete database.GetDB().Query error:",err.Error())
		return err
	}
	defer rows.Close()
	var tmp int64
	err = rows.Scan(&tmp)
	if err != nil{
		Error("ProjectDelete rows.Scan error:",err.Error())
		return err
	}
	if tmp != creater{
		Error("ProjectCreater 创建人错误,删除失败!!!")
		return errors.New("您属于非创建人")
	}
	return nil
}

func DeleteProject (projectId int64)error{
	tx,err := database.GetDB().Begin()
	if err != nil{
		Error("DeleteProject database.GetDB().Begin() error:",err.Error())
		return err
	}
	ExecSql1 := "delete from project where id=?"
	Info("DeleteProject ExecSql:",ExecSql1)
	Info("DeleteProject ExecSql value=",projectId)
	_,err = database.GetDB().Exec(ExecSql1,projectId)
	if err != nil{
		Error("DeleteProject  database.GetDB().Exec error:",err.Error())
		tx.Rollback()
		return err
	}

	ExecSql2 := "delete from project_people where id=?"
	Info("DeleteProject ExecSql2 =",ExecSql2)
	Info("DeleteProject ExecSql2 value=",projectId)
	_,err = database.GetDB().Exec(ExecSql2,projectId)
	if err != nil{
		Error("DeleteProject database.GetDB().Exec2:",err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}