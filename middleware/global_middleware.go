package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	Name string
	Age  int
}

func middle(c *gin.Context) {
	fmt.Println("m1 ...in")
	//中间件传递数据:设置一个key-value,后续中间件中使用Get接收数据
	c.Set("name", "fengfeng")
	//传任意类型，在接收的时候做好断言即可
	c.Set("customer", User{"a", 21})

	c.Next()
	fmt.Println("m1 ...out")
}

func middle1() gin.HandlerFunc {
	return func(c *gin.Context) {
		//这里是请求来了才会执行
		fmt.Println("middle1 ...in")
	}
}

func TimeMiddleware(c *gin.Context) {
	startTime := time.Now()
	c.Next()
	since := time.Since(startTime)
	// 获取当前请求所对应的函数
	f := c.HandlerName()
	fmt.Printf("函数 %s 耗时 %d\n", f, since)
}

func main() {
	router := gin.Default()

	//全局注册中间件
	router.Use(TimeMiddleware)
	router.GET("/", func(c *gin.Context) {
		fmt.Println("index ...in")
		name, _ := c.Get("name")
		fmt.Println(name)

		if customer, exist := c.Get("customer"); exist {
			if cus, ok := customer.(User); ok {
				fmt.Println("customer:", cus)
			}

		}
		c.JSON(200, gin.H{"msg": "index"})
		fmt.Println("index ...out")
	})

	//路由分组:
	//将一系列的路由放到一个组下，统一管理
	//例如，以下的路由前面统一加上api的前缀
	r := router.Group("/api").Use(middle, middle1()) //指定哪一些分组下可以使用中间件了
	{
		r.GET("/index", func(c *gin.Context) {
			time.Sleep(2 * time.Second)
			c.String(200, "index")
		})
		r.GET("/home", func(c *gin.Context) {
			time.Sleep(3 * time.Second)
			c.String(200, "home")
		})
	}
	
	router.Run(":8080")

}
