package database

import (
	"bug-management/conf"
	. "common/logs"
	"database/sql"
	"errors"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

var pbDb *sql.DB

func init() {
	err, db := OpenDB("mysql", conf.MyDbUser, conf.MyDbPwd, conf.MyDbHost, conf.MyDbPort, conf.MyDbName)
	if err != nil {
		fmt.Println(err)
	}
	pbDb = db
}

func GetDB()*sql.DB{
	return pbDb
}

func OpenDB(driverName, usrName, psswd, hostName, port, dbName string) (error, *sql.DB) {
	db_pbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Asia%%2FShanghai&timeout=30s&readTimeout=35s&writeTimeout=35s", usrName, psswd, hostName, port, dbName)

	db, err := sql.Open(driverName, db_pbLink)
	if err != nil {
		fmt.Println("sql open failed :", err.Error())
		return err, nil
	}
	return err, db
}

func Insert(dbtName string, colNameList []string, args ...interface{}) error {
	if pbDb == nil {
		Error("Insert pbDb is null ")
		return errors.New("not open db")
	}

	err := dbExec("INSERT "+dbtName, colNameList, "", args...)
	if err != nil {
		Error("Insert error:  ", err)
		return err
	}

	return nil
}

func Update(dbtName string, colNameList []string, keyName string, args ...interface{}) error {

	if pbDb == nil {
		Error("Update pbDb is null ")
		return errors.New("not open db")
	}

	err := dbExec("UPDATE "+dbtName, colNameList, keyName, args...)
	if err != nil {
		Error("Update error: ", err)
		return err
	}

	return nil
}

func Query(dbtName, whereKey string, whereVal interface{}, queryKeys []string, queryVals ...interface{}) error {

	if pbDb == nil {
		Error("Query pbDb is null ")
		return errors.New("not open db")
	}
	queryFieldStr := "*"

	for idx, key := range queryKeys {
		if idx == 0 {
			queryFieldStr = key
		} else {
			queryFieldStr = fmt.Sprintf("%s,%s", queryFieldStr, key)
		}
	}

	qureyCmd := fmt.Sprintf("SELECT %s FROM %s WHERE %s=?", queryFieldStr, dbtName, whereKey)
	Info("Query qureyCmd=[", qureyCmd, "].")
	Info("Query whereVal=[", whereVal, "].")
	row := pbDb.QueryRow(qureyCmd, whereVal)

	err := row.Scan(queryVals...)
	if err != nil {
		Error("Query row scan failed: ", err.Error())
		return err
	}

	return nil
}

func dbExec(cmd string, columnList []string, keyName string, args ...interface{}) error {
	if pbDb == nil {
		Error("dbExec pbDb is null ")
		return errors.New("not open db")
	}

	jobCmd := cmd + " SET "
	for idx := 0; idx < len(columnList); idx++ {
		if idx == 0 {
			jobCmd = fmt.Sprintf("%s %s=?", jobCmd, columnList[idx])
		} else {
			jobCmd = fmt.Sprintf("%s,%s=?", jobCmd, columnList[idx])
		}
	}

	if keyName != "" {
		jobCmd = fmt.Sprintf("%s WHERE %s=?", jobCmd, keyName)
	}

	stmt, err := pbDb.Prepare(jobCmd)
	if err != nil {
		Error("dbExec sql Prepare error: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		Error("dbExec sql stmt.Exec() failed: ", err)
		return err
	}

	return nil
}


func HasItem(dbtName, queryKey string, queryVal interface{}) bool {
	if pbDb == nil {
		Error("HasItem pbDbR is null ")
		return false
	}

	var db_pbSelectCount string
	db_pbSelectCount = fmt.Sprintf("SELECT COUNT(*)  FROM %s WHERE %s=?", dbtName, queryKey)

	var count int

	err := pbDb.QueryRow(db_pbSelectCount, queryVal).Scan(&count)
	if err != nil {
		Error("HasItem db.QueryRow error ", err)
		return false
	}

	if count == 0 {
		return false
	}

	return true
}