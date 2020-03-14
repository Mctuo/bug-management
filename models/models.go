package models

import (
	. "common/logs"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/context"
	"strings"
)

type SErrEventResp struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

const(
	Success = 0
	ErrArg = 2000
	ErrSvr = 5000
)

func GetErrMsg(errcode int, key string) string {
	switch errcode {
	case Success:
		return fmt.Sprintf("OK")

	case ErrArg:
		return fmt.Sprintf("参数错误:%s",key)

	case ErrSvr:
		return fmt.Sprintf("服务器内部错误:%s",key)
	}
	return key
}

func HandleError(status int, message string, data interface{}, ctx *context.Context) string {
	resp := SErrEventResp{
		Code:    status,
		Message: message,
		Data:    data,
	}
	res, err := json.Marshal(resp)
	if err != nil {
		return ""
	}
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json;charset=utf-8")
	ctx.WriteString(string(res))
	return string(res)
}

func PrintClientInfo(ctx *context.Context) (string, string, string) {
	myMethod := ctx.Request.Method
	myUrl := ctx.Request.URL.String()
	myBody := string(ctx.Input.RequestBody)
	myRemoteAddr := string(ctx.Request.RemoteAddr)
	myUserAgent := string(ctx.Request.UserAgent())
	myXRealIP := string(ctx.Request.Header.Get("X-Real-IP"))
	myXForwardedFor := string(ctx.Request.Header.Get("X-Forwarded-For"))
	myXForwardedFors := strings.Split(string(myXForwardedFor), ",")
	myRealIP := strings.TrimSpace(myXForwardedFors[0])
	if myRealIP == "" {
		myRealIP = myXRealIP
	}
	Info("ClientInfo [", myMethod, "]. url=[", myUrl, "]. body=[", myBody, "]. realIp=[", myRealIP, "]. XRealIP=[", myXRealIP, "]. RemoteAddr=[", myRemoteAddr, "]. XForwardedFor=[", myXForwardedFor, "]. UserAgent=[", myUserAgent, "]. ")
	return myUrl, myRealIP, myUserAgent
}