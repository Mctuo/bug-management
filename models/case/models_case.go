package _case

import (
	."common/logs"
	"bug-management/conf"
	"bug-management/database"
	"errors"
)

type StruCreateCaseReq struct {
	Id int64	`json:"id"`
	ProjectId int64	`json:"projectId"`
	Title string `json:"title"`
	ModulePath string `json:"modulePath"`
	Assign int64	`json:"assign"`
	Priority int64	`json:"priority"`
	TypeMethod string `json:"typeMethod"`
	TypePlan string `json:"typePlan"`
	Creator int64	`json:"creater"`
}

type StruAssignResp struct {
	ListInfo []StruCreateCaseReq `json:"listInfo"`
}

var(
	caseInfo =[]string{
		"projectId",
		"title",
		"module_path",
		"assign",
		"priority",
		"type_method",
		"type_plan",
		"creator",
	}
)

func CreateCase(myReq StruCreateCaseReq)error{
	isExist := database.HasItem(conf.MyCaseTb,"title",myReq.Title)
	if isExist{
		Error("CreateCase 已经创建case title=",myReq.Title)
		return errors.New("此title已创建，不能重复创建")
	}

	err := database.Insert(conf.MyCaseTb,caseInfo,myReq.ProjectId,myReq.Title,myReq.ModulePath,myReq.Assign,myReq.Priority,myReq.TypeMethod,myReq.TypePlan,myReq.Creator)
	if err != nil{
		Error("CreateCase database.Insert error:",err.Error())
		return err
	}
	return nil
}

func AssignList(assign int64, myResp *StruAssignResp)error{
	selectSql := "select *from test_case where assign = ?"
	Info("models AssignList selectSql=",selectSql)
	Info("models AssignList selectSql value=",assign)

	rows,err :=database.GetDB().Query(selectSql,assign)
	if err != nil{
		Error("models AssignList database.GetDB().Query error:",err.Error())
		return err
	}
	defer rows.Close()
	var tmp  StruCreateCaseReq
	for rows.Next(){
		err = rows.Scan(&tmp.Id,&tmp.ProjectId,&tmp.Title,&tmp.ModulePath,&tmp.Assign,&tmp.Priority,&tmp.TypeMethod,&tmp.TypePlan,&tmp.Creator)
		if err != nil{
			Error("models AssignList rows.Scan error:",err.Error())
			return err
		}
		myResp.ListInfo = append(myResp.ListInfo,tmp)
	}
	return nil
}

func ListByCaseId(caseId int64,myResp *StruCreateCaseReq)error{
	selectSql := "select *from test_case where case_id =?"

	Info("ListByCaseId selectSql=",selectSql)
	err := database.GetDB().QueryRow(selectSql,caseId).Scan(myResp)
	if err != nil{
		Error("ListByCaseId database.GetDB().QueryRow error:",err.Error())
		return err
	}
	return nil
}