# Go指针用法详解

## 在base1中的实际应用
```go
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request)
```
这里的`*Engine`表示接收者是Engine的指针，允许方法修改Engine的状态。

## 指针基础
1. 声明指针：`var ptr *Type`
2. 获取地址：`ptr := &value`
3. 获取值：`value := *ptr`

## 使用场景
1. 方法需要修改接收者状态
2. 避免大对象复制
3. 表示可空值（类似Java的null）

## 与Java对比
- Go的指针更显式，需要手动操作
- Java的引用类型自动处理指针操作
- Go可以直接操作内存地址，但更安全