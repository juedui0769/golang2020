package web

import (
	"fmt"
	"net/http"
	"strings"
)

func SayHelloName(writer http.ResponseWriter, req *http.Request) {
	req.ParseForm()       // 解析参数，默认是不会解析的
	fmt.Println(req.Form) // 这些信息时输出到服务器端的打印信息
	fmt.Println("path", req.URL.Path)
	fmt.Println("scheme", req.URL.Scheme)
	fmt.Println(req.Form["url_long"])
	for k, v := range req.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(writer, "Hello golang!") // 这个是写入到 writer 的是输出到客户端的
}
