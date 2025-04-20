// main.go 详细注释版
package main

import (
    "fmt"           // 提供格式化I/O功能
    "log"           // 提供日志功能
    "net/http"      // 提供HTTP服务器功能
)

// indexHandler处理根路径请求
// 对应Java中的 handle 方法第一个case
func indexHandler(w http.ResponseWriter, req *http.Request) {
    // fmt.Fprintf 类似 Java 的 PrintWriter.format
    fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// helloHandler处理/hello路径请求
// 对应Java中的 handle 方法第二个case
func helloHandler(w http.ResponseWriter, req *http.Request) {
    // range循环遍历header，类似Java的forEach
    for k, v := range req.Header {
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }
}

func main() {
    // HandleFunc注册路由，类似Java的createContext
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/hello", helloHandler)
    
    // ListenAndServe启动服务器，类似Java的server.start()
    // log.Fatal相当于Java的System.exit(1)加错误日志
    log.Fatal(http.ListenAndServe(":9999", nil))
}