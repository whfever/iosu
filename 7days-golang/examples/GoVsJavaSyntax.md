# Go vs Java 语法对比说明

## 1. 包声明和导入
```go
// Go
package main

import (
    "fmt"
    "net/http"
)
```
```java
// Java
package com.example;

import java.io.IOException;
import java.net.http.HttpServer;
```
区别：
- Go使用 `package main` 表示可执行程序，Java包名与目录结构对应
- Go导入使用双引号，Java不需要
- Go导入可以使用括号分组，Java每行一个import

## 2. 结构体vs类
```go
// Go结构体
type Engine struct{}
```
```java
// Java类
public class Engine {}
```
区别：
- Go使用 `type` 关键字定义类型
- Go没有访问修饰符(public/private)，大写开头就是public
- Go结构体更轻量，没有继承概念

## 3. 方法定义
```go
// Go方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    // ...
}
```
```java
// Java方法
public void handle(HttpExchange exchange) throws IOException {
    // ...
}
```
区别：
- Go在方法名前用 `(engine *Engine)` 指定接收者，类似Java的this
- Go不需要声明异常抛出
- Go参数类型放在变量名后面

## 4. 错误处理
```go
// Go错误处理
if err := srv.ListenAndServe(); err != nil {
    log.Fatal(err)
}
```
```java
// Java异常处理
try {
    server.start();
} catch (IOException e) {
    e.printStackTrace();
    System.exit(1);
}
```
区别：
- Go使用返回值检查错误
- Java使用try-catch异常处理
- Go的错误是值，Java的异常是对象

## 5. 并发处理
```go
// Go并发
go func() {
    srv.ListenAndServe()
}()
```
```java
// Java并发
new Thread(() -> {
    server.start();
}).start();
```
区别：
- Go使用 `go` 关键字启动协程，更轻量
- Java使用Thread或线程池
- Go channel通信，Java共享内存通信

## 6. 变量声明
```go
// Go变量
var srv *http.Server
srv := &http.Server{} // 简短声明
```
```java
// Java变量
HttpServer server;
server = HttpServer.create();
```
区别：
- Go可以使用 `:=` 自动推导类型
- Go指针操作更简单，不需要显式内存管理
- Go变量未使用会编译错误

## 7. 闭包和函数式编程
```go
// Go闭包
quit := make(chan os.Signal, 1)
signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
```
```java
// Java Lambda
Runtime.getRuntime().addShutdownHook(new Thread(() -> {
    System.out.println("Shutting down...");
}));
```
区别：
- Go支持闭包和函数作为一等公民
- Go使用channel进行通信
- Java 8+使用Lambda表达式

## 实战建议
1. 从Java转Go要适应：
   - 没有类，使用结构体和接口
   - 错误处理方式改变
   - 包的概念更简单
   - 指针的使用更频繁但更安全
   - channel代替共享内存通信

2. Go的特色：
   - 简洁的语法
   - 内置并发支持
   - 快速的编译
   - 强大的标准库
   - 自带格式化工具gofmt