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



func CreateProject(myReq StruCreateProjectReq)error{
	IsExist := database.HasItem(conf.MyProjectTb,"project_name",myReq.ProjectName)
	if IsExist{
		return errors.New("项目已创建，不能重复创建")
	}

	err := database.Insert(conf.MyProjectTb,ProjectInfo,myReq.ProjectName,myReq.Account,0,0,0,0)
	if err != nil{
		Error("CreateProject database.Insert error:",err.Error())
		return err
	}
	return nil
}