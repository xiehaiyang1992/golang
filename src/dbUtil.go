package main

import (
	"database/sql"
	"fmt"
)

var dbUtil *sql.DB

func dbInit() {
	dbUtil, _ = sql.Open("mysql", "haiyang:haiyang@tcp(119.29.183.232:3306)/go")
	dbUtil.SetMaxOpenConns(3)
	dbUtil.SetMaxIdleConns(3)
	err := dbUtil.Ping()
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println("数据库连接池已连接")
	}
}
