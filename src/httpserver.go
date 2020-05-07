package main

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

func httpServer() {
	//静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./src/static"))))
	//非静态资源
	http.HandleFunc("/", welcome)
	//mysql测试
	http.HandleFunc("/mysql", mysql)
	//redis测试
	http.HandleFunc("/redis", httpRedis)
	err := http.ListenAndServe("127.0.0.1:80", nil)
	if err != nil {
		fmt.Print(err)
	}
}

//欢迎页面
func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome"))
	fmt.Println(r.URL)
}

//mysql
func mysql(w http.ResponseWriter, r *http.Request) {
	queryId := strconv.Itoa(rand.Int() % 10000)
	row := dbUtil.QueryRow("select md5 from md5 where id =" + queryId)
	var md5 string
	row.Scan(&md5)
	fmt.Println(queryId)
	w.Write([]byte(queryId + ":" + md5))
}

//redis
func httpRedis(w http.ResponseWriter, r *http.Request) {
	queryId := strconv.Itoa(rand.Int() % 10)
	var md5str string
	md5str, _ = redisUtil.Get(queryId).Result()
	if len(md5str) != 0 {
		fmt.Print(len(md5str))
		fmt.Print("已存在")
	} else {
		md5str = fmt.Sprintf("%x", md5.Sum([]byte(queryId)))
		redisUtil.Set(queryId, md5str, 0)
	}
	fmt.Println(queryId + ":" + md5str)
	w.Write([]byte(md5str))
}
