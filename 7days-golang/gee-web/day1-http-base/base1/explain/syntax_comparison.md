# Go vs Java 语法对比（基础HTTP服务器）

## 1. 路由处理对比
### Go方式：独立函数
```go
func indexHandler(w http.ResponseWriter, req *http.Request)
func helloHandler(w http.ResponseWriter, req *http.Request)
```

### Java方式：类方法
```java
class Engine implements HttpHandler {
    public void handle(HttpExchange exchange)
}
```

## 2. 关键语法差异

### 函数声明
- Go: `func name(参数 类型) 返回值`
- Java: `访问修饰符 返回值类型 方法名(类型 参数)`

### 错误处理
- Go: 使用多返回值和错误检查
  ```go
  if err := http.ListenAndServe(":9999", nil); err != nil
  ```
- Java: 使用try-catch
  ```java
  try {
      server.start();
  } catch (IOException e)
  ```

### 路由注册
- Go: 直接使用全局函数
  ```go
  http.HandleFunc("/", indexHandler)
  ```
- Java: 通过server对象方法
  ```java
  server.createContext("/", new Engine())
  ```

## 3. 主要概念对应关系

| Go概念 | Java概念 |
|--------|----------|
| `http.ResponseWriter` | `HttpExchange` |
| `*http.Request` | `HttpExchange` |
| `http.HandleFunc` | `createContext` |
| `log.Fatal` | `System.exit(1)` |
| `fmt.Fprintf` | `PrintWriter.format` |

## 4. 代码组织方式
- Go：偏向功能性编程，使用独立函数
- Java：偏向面向对象，使用类封装

## 5. 优点对比
### Go的优点
1. 代码更简洁
2. 路由注册更直观
3. 参数类型更清晰

### Java的优点
1. 面向对象特性更强
2. 异常处理更系统化
3. 类型系统更严格

## 6. 开发环境设置
### IDE可访问性支持
- Go:
  - VS Code: 使用 "editor.accessibilitySupport": "on"
  - GoLand: 内置屏幕阅读器支持
- Java:
  - Eclipse: 内置JAWS支持
  - IntelliJ: 支持多种屏幕阅读器

### 开发工具建议
- Go: VS Code + Go扩展
- Java: Eclipse或IntelliJ IDEA

### 调试工具对比
- Go: Delve调试器，支持键盘导航
- Java: Java调试器，完整的键盘快捷键支持