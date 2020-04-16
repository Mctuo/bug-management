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
	}
)

var(
	projectinfo =[]string{
		"project_name",
		"account",
		"project_people",
		"task_total",
		"task_unfinished",
		"task_finished",
	}
)

type StruCreateProjectReq struct {
	ProjectName string `json:"project_name"`
	Account int64	`json:"account"`
}

type StruProject struct {
	ProjectName string `json:"project_name"`
	Account int64	`json:"account"`
	ProjectPeople int	`json:"project_people"`
	TotalTask int	`json:"total_task"`
	TaskFinished int	`json:"task_finished"`
	TaskUnFinished int	`json:"task_unfinished"`
}

type StruProjectListResp struct {
	ListInfo []StruProject `json:"listInfo"`
}


func CreateProject(myReq StruCreateProjectReq)error{
	IsExist := database.HasItem(conf.MyProjectTb,"project_name",myReq.ProjectName)
	if IsExist{
		return errors.New("项目已创建，不能重复创建")
	}

	err := database.Insert(conf.MyProjectTb,ProjectInfo,myReq.ProjectName,myReq.Account,1,0,0,0)
	if err != nil{
		Error("CreateProject database.Insert error:",err.Error())
		return err
	}
	return nil
}

func ProjectList(account int64,myResp *StruProjectListResp)error{
	var tmp StruProject

	rows,err :=database.GetDB().Query("select *from project where account =?",account)
	if err != nil{
		Error("ProjectList database.GetDB().Query error:",err.Error())
		return err
	}
	var id int
	for rows.Next(){
		err = rows.Scan(&id,&tmp.ProjectName,&tmp.Account,&tmp.ProjectPeople,&tmp.TotalTask,&tmp.TaskUnFinished,&tmp.TaskFinished)
		if err != nil{
			Error("ProjectList rows.Scan error:",err.Error())
			return err
		}
		myResp.ListInfo = append(myResp.ListInfo,tmp)
	}
	return nil
}