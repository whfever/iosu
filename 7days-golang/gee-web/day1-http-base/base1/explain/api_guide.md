# Web服务器 API 说明

## 端点 (Endpoints)
1. GET / 
   - 功能：显示当前请求的URL路径
   - 返回格式：`URL.Path = "/"`

2. GET /hello
   - 功能：显示所有HTTP请求头
   - 返回格式：`Header["User-Agent"] = ["curl/7.54.0"]`

## 测试方法
```bash
# 测试根路径
curl http://localhost:9999/

# 测试hello路径
curl http://localhost:9999/hello
```

## 目录结构
```
base1/
├── main.go          # 主程序代码
├── explain/         # 说明文档
│   ├── pointer_guide.md    # 指针用法说明
│   ├── go_vs_java.md      # Go和Java对比
│   └── api_guide.md       # API文档
```