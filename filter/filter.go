package filter

import (
	"bug-management/conf"
	. "common/logs"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"io/ioutil"
	"net/http"
	"time"
)

type StruIsLoginResp struct {
	Code int 	 		`json:"code"`
	Msg  string  		`json:"msg"`
	Data interface{}	`json:"data"`
}

func SsoFilter()beego.FilterFunc{
	return func(ctx* context.Context){
		myRequestURI := ctx.Request.RequestURI
		Info("SsoFilter RequestURI:", myRequestURI)

		autoOn := conf.AutoOn

		if autoOn == 1{
			SsoToken := ctx.GetCookie("SsoToken")

			if len(SsoToken) < 1{
				ctx.ResponseWriter.WriteHeader(401)
				return
			}

			if !CheckToken(SsoToken){
				ctx.ResponseWriter.WriteHeader(403)
				return
			}
		}
	}
}

func CheckToken(SsoToken string)bool{
	client := http.Client{Timeout:5*time.Second}

	addr := fmt.Sprintf("%s",conf.IsLoginAddr)
	myReq,err := http.NewRequest("GET",addr,nil)
	if err != nil{
		Error("CheckToken http.NewRequest error:",err.Error())
		return false
	}
	myReq.Header.Set("cookie",fmt.Sprintf("SsoToken=%s",SsoToken))

	Resp,err :=client.Do(myReq)
	if err != nil{
		Error("SsoToken CheckToken client.Do error:",err.Error())
		return false
	}
	var myResp StruIsLoginResp

	RespBody,err :=ioutil.ReadAll(Resp.Body)
	if err != nil{
		Error("CheckToken ioutil.ReadAll error:",err.Error())
		return false
	}
	defer Resp.Body.Close()
	err = json.Unmarshal(RespBody,&myResp)
	if err != nil{
		Error("CheckToken json.Unmarshal error:",err.Error())
		return false
	}
	if myResp.Code != 0{
		Error("CheckToken  myResp.Code =", myResp.Code)
		return false
	}
	return true
}