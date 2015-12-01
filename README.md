# 只是一个用 Go 写的炒鸡简单的 HTTP 重定向服务

## Usage

```
$ ALIASES_FILE=/tmp/test.aliases ./go-redirect
```

## Alias File

```
# 我是一个注释
# 将 viz 重定向到 vizv.com
viz: http://vizv.com
# 注意一定要带协议 http:// 或 https://
```
