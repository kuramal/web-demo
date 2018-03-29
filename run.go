package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func hostIPGet() string {

	var ips string

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips += ipnet.IP.String() + "    "
			}

		}
	}

	return ips
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	/*
		r.ParseForm()       //解析参数，默认是不会解析的
		fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
	*/
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Printf("%v  %v   \n", r.RemoteAddr, time.Now())
	fmt.Fprintf(w, "Hello world!    ")                //这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "%v\n", time.Now())                //这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "............    %v", hostIPGet()) //这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "\n")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage\n")
}

func main() {
	http.HandleFunc("/", home)              //设置访问的路由
	http.HandleFunc("/hello", sayhelloName) //设置访问的路由
	/*	err := http.ListenAndServe(":8080", nil) //设置监听的端口
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	*/
	err := http.ListenAndServeTLS(":443", "/root/server.crt", "/root/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
