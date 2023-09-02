# 书签管理器

## 示例

```
curl -X POST http://127.0.0.1:8080/api/v1/bookmark/create  -H 'Content-Type: application/json' -d '{"name":"test","url":"baidu.com"}'
```

```
curl -X POST 'http://127.0.0.1:8080/api/v1/bookmark/page' -H 'Content-Type: application/json' -d '{}'
```