package main

import (
    "context"        // 上下文处理
    "fmt"           // 格式化I/O
    "log"           // 日志
    "net/http"      // HTTP服务
    "os"            // 操作系统功能
    "os/signal"     // 信号处理
    "syscall"       // 系统调用
    "time"          // 时间相关
)

// Engine 结构体实现了 http.Handler 接口
// 相当于 Java 中实现 HttpHandler 接口的类
type Engine struct{}

// ServeHTTP 方法实现了 http.Handler 接口
// 相当于 Java 中的 handle 方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.URL.Path {
    case "/":
        fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
    case "/hello":
        // range 类似 Java 的 for-each
        for k, v := range req.Header {
            fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
        }
    default:
        fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
    }
}

func main() {
    // new() 返回指针，类似 Java 的 new
    engine := new(Engine)
    
    // 结构体初始化，类似 Java 的构造函数
    srv := &http.Server{
        Addr:    ":9999",
        Handler: engine,
    }

    // go 关键字启动协程（轻量级线程）
    // 类似 Java 的 new Thread(() -> {...}).start()
    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    // 创建信号通道，用于优雅关闭
    // 类似 Java 的 Runtime.getRuntime().addShutdownHook()
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Println("正在关闭服务器...")

    // 创建5秒超时的上下文
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()  // 函数返回时调用，类似 Java 的 finally

    // 优雅关闭服务器
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("服务器被强制关闭:", err)
    }

    log.Println("服务器已退出")
}