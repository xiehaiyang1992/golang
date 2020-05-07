package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func insertRedis() {
	redisInit() //初始化redis
	//var i int
	for i := 0; i < 999999; i++ {
		redisUtil.Set(strconv.Itoa(i), fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(i)))), 0)
		//fmt.Println(i)
	}
}
