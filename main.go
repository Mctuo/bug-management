package main

import (
	"bug-management/conf"
	_ "bug-management/routers"
	"common/logs"
	"fmt"
	"github.com/astaxie/beego"
	"thirdapi/httpSim"
)

func init() {
	fmt.Println(conf.RunMode, conf.LogFilename, conf.LogMaxlines, conf.LogMaxsize, conf.LogDaily, conf.LogRotate, conf.LogLevel)
	logs.LogInit(conf.RunMode, conf.LogFilename, conf.LogMaxlines, conf.LogMaxsize, conf.LogDaily, conf.LogRotate, conf.LogLevel)
	httpSim.HttpSimInit(conf.HttpClientOn, conf.HttpConnTimeOut, conf.HttpReadTimeOut, conf.HttpWriteTimeOut)
}

func main() {
	beego.Run()
}

