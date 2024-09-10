package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

func main() {
	router := gin.Default()

	//加载templates文件夹下所有文件模板
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	router.GET("/user", func(c *gin.Context) {
		//指定默认值
		//http://localhost:8080/user 才会打印出来默认的值
		name := c.DefaultQuery("name", "枯藤")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	//JSON
	router.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	// 结构体转json
	router.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		type Msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg := Msg{"fengfeng", "hey", 21}
		// 注意 msg.Name 变成了 "user" 字段
		// 以下方式都会输出 :   {"user": "hanru", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	//返回xml
	router.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"user": "hanru", "message": "hey", "status": http.StatusOK})
	})

	//返回yaml
	router.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"user": "hanru", "message": "hey", "status": http.StatusOK})
	})

	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/tem", func(c *gin.Context) {
		//渲染模板 ,test.html中有这个 title接收 : {{  .title  }
		c.HTML(http.StatusOK, "test.html", gin.H{"title": "我是测试", "ce": "123456"})
	})

	router.GET("/redirect", func(c *gin.Context) {
		//支持内部和外部的重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	router.GET("/query", query)
	router.GET("/param/:user_id/:book_id", param)

	//可以接收 multipart/form-data; 和application/x-www-form-urlencoded
	// application/x-www-form-urlencoded:轻型表单，只支持普通文本,被encodeURI编码过的querystring,如：name=%E5%BC%A0%E4%B8%89&age=18
	//multipart/form-data:每一个part部分，都支持不同的Content-Type，比如图片、音频、视频等。
	router.POST("/form", form)
	//POST raw
	router.POST("/raw", raw)

	//请求头相关
	router.GET("/head", func(c *gin.Context) {
		// 首字母大小写不区分  单词与单词之间用 - 连接
		// 用于获取一个请求头
		fmt.Println("head:", c.GetHeader("User-Agent"))
		//fmt.Println(c.GetHeader("user-agent"))
		//fmt.Println(c.GetHeader("user-Agent"))
		//fmt.Println(c.GetHeader("user-AGent"))

		// Header 是一个普通的 map[string][]string
		fmt.Println("head1:", c.Request.Header)
		// 如果是使用 Get方法或者是 .GetHeader,那么可以不用区分大小写，并且返回第一个value
		fmt.Println(".Get(User-Agent):", c.Request.Header.Get("User-Agent"))
		fmt.Println("c.Request.Header[User-Agent]:", c.Request.Header["User-Agent"])
		// 如果是用map的取值方式，请注意大小写问题
		fmt.Println("c.Request.Header[User-Agent]1:", c.Request.Header["user-agent"])

		// 自定义的请求头，用Get方法也是免大小写
		fmt.Println("Token:", c.Request.Header.Get("Token"))
		fmt.Println("token:", c.Request.Header.Get("token"))
		c.JSON(200, gin.H{"msg": "成功"})
	})
	// 设置响应头
	router.GET("/res", func(c *gin.Context) {
		c.Header("Token", "jhgeu%hsg845jUIF83jh")
		c.Header("Content-Type", "application/text; charset=utf-8")
		c.JSON(0, gin.H{"data": "看看响应头"})
	})
	router.Run(":8080")
}

func query(c *gin.Context) {
	fmt.Println(c.Query("user"))
	fmt.Println(c.DefaultQuery("addr", "四川省"))
	fmt.Println(c.GetQuery("user"))
	fmt.Println(c.QueryArray("user")) // 拿到多个相同的查询参数
}

func param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
	fmt.Println(c.Param("book_id"))
}

func form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "四川省")) // 如果用户没传，就使用默认值
	forms, err := c.MultipartForm()               // 接收所有的form参数，包括文件
	fmt.Println(forms, err)
}

func raw(c *gin.Context) {
	body, _ := c.GetRawData()
	contentType := c.GetHeader("Content-Type")
	switch contentType {
	case "application/json":

		// json解析到结构体
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		err := json.Unmarshal(body, &user)
		if err != nil {
			fmt.Println("Unmarshal error:", err.Error())
		}
		fmt.Println(user)
	}
}
