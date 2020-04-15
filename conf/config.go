package conf

import (
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

var(
	RunMode string
	AutoOn int

	LogFilename string
	LogMaxlines int
	LogMaxsize  int
	LogDaily    bool
	LogRotate   bool
	LogLevel    int

	RsAddr  string
	RsPwd         string
	Rsdb          int
	RsExpire      int

	//http
	HttpConnTimeOut  int
	HttpReadTimeOut  int
	HttpWriteTimeOut int
	HttpClientOn 	 int

	//mysql
	MyDbUser             string
	MyDbPwd              string
	MyDbHost             string
	MyDbPort             string
	MyDbName			 string
	MyUserTb			 string
	MyProjectTb  		 string
	MyProjectPeople		 string
	//url
	IsLoginAddr 		 string
)


func init(){
	var err error

	if MyProjectPeople = beego.AppConfig.String("mysql::myprojectPeople");MyProjectPeople ==""{
		panic("MyProjectPeople is null")
	}

	if MyUserTb = beego.AppConfig.String("mysql::myusertb");MyUserTb == ""{
		panic("MyUserTb is null")
	}

	if IsLoginAddr = beego.AppConfig.String("url::isloginaddr");IsLoginAddr==""{
		panic("IsLoginAddr is not set")
	}

	if AutoOn,err = beego.AppConfig.Int("autoon");AutoOn < 0 || err != nil{
		panic("AutoOn config error")
	}
	if MyDbUser = beego.AppConfig.String("mysql::mydbuser"); MyDbUser == "" {
		fmt.Println("MyDbUser path is not set , default is ", MyDbUser)
		panic("MyDbUser is not set , default is null")
	}

	if MyProjectTb = beego.AppConfig.String("mysql::myprojecttb");MyProjectTb ==""{
		panic("MyProjectTb is not set,default is null")
	}

	if MyDbPwd = beego.AppConfig.String("mysql::mydbpwd"); MyDbPwd == "" {
		fmt.Println("MyDbPwd path is not set , default is ", MyDbPwd)
		panic("MyDbPwd is not set , default is null")
	}

	if MyDbHost = beego.AppConfig.String("mysql::mydbhost");MyDbHost == ""{
		panic("MyDbHost is not set,default is null")
	}

	if MyDbPort = beego.AppConfig.String("mysql::mydbport"); MyDbPort == "" {
		fmt.Println("MyDbPort path is not set , default is ", MyDbPort)
		panic("MyDbPort is not set , default is null")
	}
	if MyDbName = beego.AppConfig.String("mysql::mydbname"); MyDbName == "" {
		fmt.Println("MyDbName path is not set , default is ", MyDbName)
		panic("MyDbName is not set , default is null")
	}

	if RunMode = beego.AppConfig.String("runmode"); RunMode == "" {
		RunMode = "debug"
		fmt.Println("RunMode is not set , default is  ", RunMode)
	}

	if LogFilename = beego.AppConfig.String("log::filename"); LogFilename == "" {
		panic("log file name is not set , default is null")
	}

	if LogRotate, err = beego.AppConfig.Bool("log::rotate"); err != nil {
		LogRotate = false
		fmt.Println("rotate is not set , default is ", LogRotate)
	}
	if LogMaxsize, err = beego.AppConfig.Int("log::maxsize"); err != nil || LogMaxsize < 0 {
		LogMaxsize = 1 << 28
		fmt.Println("LogMaxsize is not set , default is  , err ", 1<<28, err)
	}

	if LogMaxlines, err = beego.AppConfig.Int("log::maxlines"); err != nil || LogMaxlines < 0 {
		LogMaxlines = 1000000
		fmt.Println("LogMaxlines is not set , default is  , err ", LogMaxlines, err)
	}
	if LogDaily, err = beego.AppConfig.Bool("log::daily"); err != nil {
		LogDaily = false
		fmt.Println("LogDaily is not set , default is ", LogDaily)
	}

	if LogLevel, err = beego.AppConfig.Int("log::level"); err != nil || LogLevel < 0 {
		LogLevel = 8
		fmt.Println("LogLevel is not set , default is ", LogLevel)
	}

	if RsPwd = strings.TrimSpace(beego.AppConfig.String("redis::rspassword")); RsPwd == "" {
		RsPwd = ""
		fmt.Println("RsPwd is not set , default is null")
	}

	if RsExpire, err = beego.AppConfig.Int("redis::rsexpire"); err != nil || RsExpire < 0 {
		RsExpire = 120
		fmt.Println("RsExpire is not set , default is ", RsExpire)
	}

	if Rsdb, err = beego.AppConfig.Int("redis::rsdb"); err != nil || Rsdb < 0 {
		Rsdb = 0
		fmt.Println("Rsdb is not set , default is ", Rsdb)
	}

	if RsAddr = strings.TrimSpace(beego.AppConfig.String("redis::rsaddr")); RsAddr ==""{
		panic("RsAddr is not set, default is null")
	}

	if HttpClientOn, err = beego.AppConfig.Int("svrinfo::httpclienton"); err != nil || HttpClientOn < 0 {
		HttpClientOn = 0
		fmt.Println("HttpClientOn is not set , default is ", HttpClientOn)
	}

	if HttpConnTimeOut, err = beego.AppConfig.Int("httpcfg::conntimeout"); err != nil || HttpConnTimeOut < 0 {
		HttpConnTimeOut = 2000
		fmt.Println("HttpConnTimeOut is not set , default is ", HttpConnTimeOut)
	}

	if HttpReadTimeOut, err = beego.AppConfig.Int("httpcfg::readtimeout"); err != nil || HttpReadTimeOut < 0 {
		HttpReadTimeOut = 2000
		fmt.Println("HttpReadTimeOut is not set , default is ", HttpReadTimeOut)
	}

	if HttpWriteTimeOut, err = beego.AppConfig.Int("httpcfg::writetimeout"); err != nil || HttpWriteTimeOut < 0 {
		HttpWriteTimeOut = 2000
		fmt.Println("HttpWriteTimeOut is not set , default is ", HttpWriteTimeOut)
	}
}
