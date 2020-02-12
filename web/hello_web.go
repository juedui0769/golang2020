package main

import (
	"log"
	"net/http"
	"web/web"
)

func main() {
	invokeSayHelloName()
}



// 2020年2月12日
func invokeSayHelloName() {
	http.HandleFunc("/", web.SayHelloName)       // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}


