package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	// 输出到文件
	f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//不想看到这些debug日志，那么我们可以改为release模式
	//gin.SetMode(gin.ReleaseMode)
	//router := gin.Default()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "/"})
	})

	//router := gin.New()
	////自定义LOG输出样式
	//router.Use(gin.LoggerWithFormatter(LoggerWithFormatter))
	//router.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{"msg": "/"})
	//})
	//router.Run(":8080")
}

// LoggerWithFormatter 输出有颜色的log
func LoggerWithFormatter(params gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	statusColor = params.StatusCodeColor()
	methodColor = params.MethodColor()
	resetColor = params.ResetColor()

	return fmt.Sprintf(
		"[ feng ] %s  | %s %d  %s | \t %s | %s | %s %-7s %s \t  %s\n",
		params.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, params.StatusCode, resetColor, // 状态码
		params.ClientIP,                        // 客户端ip
		params.Latency,                         // 请求耗时
		methodColor, params.Method, resetColor, // 请求方法
		params.Path, // 路径
	)
}
