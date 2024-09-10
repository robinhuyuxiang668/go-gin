package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	fmt.Println("m1 ...in")
	c.Next()
	fmt.Println("m1 ...out")
	//c.JSON(200, gin.H{"msg": "第一个中间件拦截了"})
	//其中一个中间件响应了c.Abort()，后续中间件将不再执行，直接按照顺序走完所有的响应中间件
	//c.Abort()
}
func m2(c *gin.Context) {
	fmt.Println("m2 ...in")
	c.Next()
	fmt.Println("m2 ...out")
}

func main() {
    //gin.Default()默认使用了Logger和Recovery中间件，其中：
    //Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。 
    //Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
    //如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。
	router := gin.Default()

	router.GET("/", m1, func(c *gin.Context) {
		fmt.Println("index ...in")
		c.JSON(200, gin.H{"msg": "响应数据"})
		c.Next()
		fmt.Println("index ...out")
	}, m2)

	router.Run(":8080")
}
