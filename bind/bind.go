package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"` //POST form提交必填form否则接收不到
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex" binding:"oneof=male female"`
}

// Should Bind
// 可以绑定json，query，param，yaml，xml
// 如果校验不通过会返回错误
func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {

		var userInfo UserInfo
		err := c.ShouldBindJSON(&userInfo)
		if err != nil {
			c.JSON(200, gin.H{"msg": "你错了"})
			fmt.Println("ShouldBindJSON err:", err)
			return
		}
		c.JSON(200, userInfo)

	})

	//绑定查询参数 struct需加上 如 form:"name"
	router.POST("/query", func(c *gin.Context) {

		var userInfo UserInfo
		err := c.ShouldBindQuery(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)

	})
	router.GET("/query", func(c *gin.Context) {

		var userInfo UserInfo
		err := c.ShouldBindQuery(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)

	})

	//struct需加上 uri:"name"
	router.POST("/uri/:name/:age/:sex", func(c *gin.Context) {

		var userInfo UserInfo
		err := c.ShouldBindUri(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)

	})
	router.GET("/uri/:name/:age/:sex", func(c *gin.Context) {

		var userInfo UserInfo
		err := c.ShouldBindUri(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)

	})

	//会根据请求头中的content-type去自动绑定
	//绑定form-data、x-www-form-urlencode、json yaml
	router.POST("/form", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBind(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "wrong"})
			return
		}
		c.JSON(200, userInfo)
	})
	router.GET("/form", func(c *gin.Context) {
		var userInfo UserInfo
		err := c.ShouldBind(&userInfo)
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{"msg": "你错了"})
			return
		}
		c.JSON(200, userInfo)
	})

	router.Run(":8080")
}
