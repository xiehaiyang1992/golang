package main

import (
	"fmt"
	//"math"
	"math/rand"
	"net/http"
)

func main() {
	startup()
}

func startup() {
	http.HandleFunc("/", welcome)
	err := http.ListenAndServe("127.0.0.1:80", nil)
	if err != nil {
		fmt.Print(err)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
	//var ram int
	ram := rand.Int()
	fmt.Print(ram)
	fmt.Print(r.URL)
}
