// package main 声明这是一个可执行程序
// 如果是库，则使用其他包名
package main

// 导入需要的包
// 使用括号分组导入多个包
import (
    "context"        // 用于处理上下文，如超时控制
    "fmt"           // 用于格式化输入输出
    "log"           // 用于日志记录
    "net/http"      // 提供 HTTP 客户端和服务器实现
    "os"            // 提供操作系统功能
    "os/signal"     // 用于处理系统信号
    "syscall"       // 系统调用
    "time"          // 时间相关功能
)

// type 关键字定义新类型
// struct{} 表示空结构体，类似 Java 的空类
type Engine struct{}

// (engine *Engine) 是接收者，类似 Java 的 this
// ServeHTTP 方法名
// (w http.ResponseWriter, req *http.Request) 是参数列表，类型在变量名后
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    // switch 语句，类似 Java 的 switch
    switch req.URL.Path {
    case "/":
        // fmt.Fprintf 格式化输出到 Writer
        // %q 表示带引号的字符串
        fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
    case "/hello":
        // range 用于遍历，类似 Java 的 for-each
        for k, v := range req.Header {
            fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
        }
    default:
        fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
    }
}

func main() {
    // := 是短变量声明，自动推导类型
    // new() 分配内存并返回指针
    engine := new(Engine)
    
    // & 取地址运算符
    // 复合字面值初始化结构体
    srv := &http.Server{
        Addr:    ":9999",      // 字段名: 值
        Handler: engine,
    }

    // go 关键字启动新的协程（轻量级线程）
    go func() {
        // err != nil 是 Go 的错误处理方式
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    // make 创建 channel（通道）
    // 用于协程间通信
    quit := make(chan os.Signal, 1)
    
    // 注册信号处理
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    
    // <-quit 从通道接收数据
    <-quit
    log.Println("Shutting down server...")

    // context.WithTimeout 创建超时上下文
    // defer 在函数返回时执行
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // 优雅关闭服务器
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

    log.Println("Server exiting")
}