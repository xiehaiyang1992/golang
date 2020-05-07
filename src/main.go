package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	insertRedis()
	//redisInit()  //初始化redis
	//dbInit()     //初始化数据库
	//httpServer() //更新
}
