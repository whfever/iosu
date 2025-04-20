# Go语言中 * 符号的用法解释

## 1. 声明指针类型
```go
var ptr *int     // 声明一个指向int的指针
var srv *http.Server  // 声明一个指向http.Server的指针
```

## 2. 获取指针（取地址）
```go
num := 42
ptr := &num    // & 获取num的地址，ptr是 *int 类型
```

## 3. 解引用（获取指针指向的值）
```go
value := *ptr  // *ptr 获取指针指向的值
```

## 4. 在方法接收者中使用
```go
// *Engine 表示接收者是指针类型
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    // 方法可以修改 engine 的字段
}
```

## 5. 与Java对比
| Go | Java |
|---|---|
| `*Type` | `Type` (引用类型) |
| `&variable` | 无需显式操作 |
| `*pointer` | 自动解引用 |

## 6. 优点
1. 显式指明是否传递引用
2. 避免无意的值拷贝
3. 性能优化：避免大对象复制

## 7. 常见用例
1. 方法需要修改接收者时使用指针接收者
2. 传递大对象时使用指针避免复制
3. 值可能为空时使用指针（类似Java的null）