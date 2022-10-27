## 启动一个 gin demo

```go
package main

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
```

`gin.Default()` 初始化一个带有日志组件(logger)和恢复组件(recovery)的 gin engine 实例

日志组件将 gin 的启动日志和响应日志输出到控制台

恢复组件将 gin 不能处理的请求通过 http 500 的形式返回给调用方

`r.Run()` 将启动 gin engine 并在默认的端口 8080 上监听请求

大多数工程并没有直接启动一个裸的 gin 应用

例如通过启动一个 fx.App ,然后在 fx.App 当中利用 hook 函数启动和停止一个 ApiServer, 这个 ApiServer 通常由 gin 提供服务

这样可以自行封装 gin 的行为,更好的配合业务的使用