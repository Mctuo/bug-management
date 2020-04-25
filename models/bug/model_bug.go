package bug

import (
	"bug-management/database"
	."common/logs"
	"database/sql"
)

type StruCreateInfoReq struct {
	BugTitle    string `json:"bug_title"`
	ProjectId   int64	`json:"projectId"`
	ModulePath string `json:"module_path"`
	Assigned    int64  `json:"assigned"`
	Severity    int64  `json:"severity"`
	Priority    int64  `json:"priority"`
	CaseId 		int64	`json:"caseId"`
	Type 		string	`json:"type"`
	FindWay	string		`json:"find_way"`
	TestEnv string		`json:"test_env"`
}

type StruBugInfoRespData struct {
	BugId		int64	`json:"bugId"`
	BugTitle    string `json:"bug_title"`
	CaseId      int64	`json:"caseId"`
	ProjectId   int64	`json:"projectId"`
	ModulePath  string `json:"module_path"`
	Assigned    int64  `json:"assigned"`
	Severity    int64  `json:"severity"`
	Priority    int64  `json:"priority"`
	Type 		string	`json:"type"`
	FindWay		string		`json:"find_way"`
	TestEnv 	string		`json:"test_env"`
}

type StruBugSolutionReq struct {
	ProjectId int64	`json:"projectId"`
	CaseId	int64	`json:"caseId"`
	Solver int64	`json:"solver"`
	SolveTime int64	`json:"solve_time"`
	Solution string	`json:"solution"`
}

type StruBugSolutionData struct {
	BugId int64	`json:"bugId"`
	CaseId	int64	`json:"caseId"`
	ProjectId int64	`json:"projectId"`
	Solver int64	`json:"solver"`
	SolveTime int64	`json:"solve_time"`
	Solution string	`json:"solution"`
}

type StruBugSolutionResp struct {
	Info []StruBugSolutionData `json:"info"`
}

type StruBugInfoResp struct {
	Info []StruBugInfoRespData `json:"info"`
}



func CreateBugInfo(myReq StruCreateInfoReq)error{
	InsertSql := "insert into bug_info (bug_title,case_id,projectId,module_path,assigned,severity,priority,type,find_way,test_env)values(?,?,?,?,?,?,?,?,?,?)"
	Info("CreateBugInfo InsertSql =",InsertSql)
	Info("CreateBugInfo InsertSql value=",myReq)

	_,err := database.GetDB().Exec(InsertSql,myReq.BugTitle,myReq.CaseId,myReq.ProjectId,myReq.ModulePath,myReq.Assigned,myReq.Severity,myReq.Priority,myReq.Type,myReq.FindWay,myReq.TestEnv)
	if err != nil{
		Error("CreateBugInfo database.GetDB().Exec error:",err.Error())
		return err
	}
	return nil
}

func GetBugInfo(title string,myResp *StruBugInfoResp)error{
	var selectSql string
	var err error
	var rows *sql.Rows

	if title == ""{
		selectSql = "select *from bug_info"
		rows,err = database.GetDB().Query(selectSql)
	}else{
		selectSql = "select *from bug_info where bug_title=?"
		rows,err = database.GetDB().Query(selectSql,title)
	}
	Info("GetBugInfo selectSql =",selectSql)
	Info("GetBugInfo selectSql value=",title)

	if err != nil{
		Error("GetBugInfo error:",err.Error())
		return err
	}

	for rows.Next(){
		var tmp StruBugInfoRespData
		err = rows.Scan(&tmp.BugId,&tmp.BugTitle,&tmp.CaseId,&tmp.ProjectId,&tmp.ModulePath,&tmp.Assigned,&tmp.Severity,&tmp.Priority,&tmp.Type,&tmp.FindWay,&tmp.TestEnv)
		if err != nil{
			Error("GetBugInfo rows.Scan error:",err.Error())
			return err
		}
		myResp.Info = append(myResp.Info,tmp)
	}
	return nil
}



func GetBugInfoByAssign(assign int64,myResp *StruBugInfoResp)error{
	selectSql := "select *from bug_info where assigned =?"

	rows,err :=database.GetDB().Query(selectSql,assign)
	if err != nil{
		Error("GetBugInfoByAssign database.GetDB().Query error:",err.Error())
		return err
	}
	for rows.Next(){
		var tmp StruBugInfoRespData
		err = rows.Scan(&tmp.BugId,&tmp.BugTitle,&tmp.CaseId,&tmp.ProjectId,&tmp.ModulePath,
			&tmp.Assigned,&tmp.Severity,&tmp.Priority,&tmp.Type,&tmp.FindWay,&tmp.TestEnv)
		if err != nil{
			Error("GetBugInfoByAssign rows.Next() error:",err.Error())
			return err
		}
		myResp.Info = append(myResp.Info,tmp)
	}
	return nil
}


func CreateBugSolution(myReq StruBugSolutionReq)error{
	InsertSql := "insert into bug_solution(projectId,case_id,solver,solve_time,solution)values(?,?,?,?,?)"
	Info("CreateBugSolution InsertSql =",InsertSql)
	Info("InsertSql InsertSql value =",myReq)

	_,err := database.GetDB().Exec(InsertSql,myReq.ProjectId,myReq.CaseId,myReq.Solver,myReq.SolveTime,myReq.Solution)
	if err != nil{
		Error("CreateBugSolution database.GetDB().Exec error:",err.Error())
		return err
	}
	return nil
}

func GetSolutionList(projectId int64,myResp *StruBugSolutionResp)error{
	var selectSql string
	var err error
	var rows *sql.Rows

	if projectId == 0{
		selectSql = "select *from bug_solution"
		rows,err = database.GetDB().Query(selectSql)
	}else{
		selectSql = "select *from bug_solution where projectId=?"
		rows,err = database.GetDB().Query(selectSql,projectId)
	}
	Info("GetSolutionList selectSql=",selectSql)
	Info("GetSolutionList selectSql value =",projectId)

	if err != nil{
		Error("GetSolutionList database.GetDB().Query error:",err.Error())
		return err
	}
	for rows.Next(){
		var tmp StruBugSolutionData
		err = rows.Scan(&tmp.BugId,&tmp.CaseId,&tmp.ProjectId,&tmp.Solver,&tmp.SolveTime,&tmp.Solution)
		if err != nil{
			Error("GetSolutionList rows.Scan error:",err.Error())
			return err
		}
		myResp.Info = append(myResp.Info,tmp)
	}
	return nil
}