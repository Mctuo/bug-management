package result

import (
	"bug-management/conf"
	"bug-management/database"
	. "common/logs"
	"errors"
)

type StruResultReq struct {
	ProjectId int64	`json:"projectId"`
	CaseId    int64	`json:"caseId"`
	Status    string	`json:"status"`
	Assigned int64	`json:"assigned"`
	TestEnv  string	`json:"testEnv"`
	TestStep string	`json:"testStep"`
}

type StruResultResp struct {
	Id int64	`json:"id"`
	ProjectId int64	`json:"projectId"`
	CaseId    int64	`json:"caseId"`
	Status    string	`json:"status"`
	Assigned int64	`json:"assigned"`
	TestEnv  string	`json:"testEnv"`
	TestStep string	`json:"testStep"`
}

type StruResultListResp struct {
	Info []StruResultResp	`json:"info"`
}

var (
	ResultInfo =[]string{
		"case_id",
		"projectId",
		"status",
		"assigned",
		"test_env",
		"test_step",
	}
)

func CreateResult(myReq StruResultReq)error{
	IsExist := database.HasItem(conf.MyProjectTb,"id",myReq.ProjectId)
	if !IsExist{
		Error("models CreateResult projectId not exist!")
		return errors.New("项目不存在")
	}

	err := database.Insert(conf.MyResultTb,ResultInfo,myReq.CaseId,myReq.ProjectId,myReq.Status, myReq.Assigned,myReq.TestEnv,myReq.TestStep)
	if err != nil{
		Error("models CreateResult  database.Insert error:",err.Error())
		return err
	}
	return nil
}


func ResultList(myResp *StruResultListResp)error{
	ExecSql := "select test_result.* from test_result inner join test_case on test_result.case_id=test_case.id"
	Info("ResultList ExecSql =",ExecSql)


	rows,err := database.GetDB().Query(ExecSql)
	if err != nil{
		Error("ResultList database.GetDB().Query error:",err.Error())
		return err
	}
	defer rows.Close()
	for rows.Next(){
		var tmp StruResultResp
		err = rows.Scan(&tmp.Id,&tmp.CaseId,&tmp.ProjectId,&tmp.Status,&tmp.Assigned,&tmp.TestEnv,&tmp.TestStep)
		if err != nil{
			Error("ResultList  rows.Scan error:",err.Error())
			return err
		}
		myResp.Info = append(myResp.Info,tmp)
	}
	return nil
}