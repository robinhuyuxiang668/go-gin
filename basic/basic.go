package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由 ,使用gin的Default方法创建一个路由Handler；
	r := gin.Default()
	// 绑定路由规则，执行的函数 ,gin.Context，封装了request和response
	r.GET("/index", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	// 3.监听端口，默认在8080,不指定端口号默认为8080   http://127.0.0.1:8000/
	r.Run(":8080")
	// 启动方式二:用原生http服务的方式， router.Run本质就是http.ListenAndServe的进一步封装
	//http.ListenAndServe(":8080", router)
}
