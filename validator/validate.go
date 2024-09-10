package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// 常用验证器
// required： 必填字段，如：binding:"required"
// // 针对字符串的长度
// min 最小长度，如：binding:"min=5"
// max 最大长度，如：binding:"max=10"
// len 长度，如：binding:"len=6"
//
// // 针对数字的大小
// eq 等于，如：binding:"eq=3"
// ne 不等于，如：binding:"ne=12"
// gt 大于，如：binding:"gt=10"
// gte 大于等于，如：binding:"gte=10"
// lt 小于，如：binding:"lt=10"
// lte 小于等于，如：binding:"lte=10"
// // 针对同级字段的
// eqfield 等于其他字段的值，如：PassWord string `binding:"eqfield=Password"`
// nefield 不等于其他字段的值
//
// 枚举  只能是red 或green
//oneof=red green
//// 字符串
//contains=fengfeng  // 包含fengfeng的字符串
//excludes // 不包含
//startswith  // 字符串前缀
//endswith  // 字符串后缀
//
//// 数组
//dive  // dive后面的验证就是针对数组中的每一个元素
//
//// 网络验证
//ip
//ipv4
//ipv6
//uri
//url
//// uri 在于I(Identifier)是统一资源标示符，可以唯一标识一个资源。
//// url 在于Locater，是统一资源定位符，提供找到该资源的确切路径
//
//// 日期验证  1月2号下午3点4分5秒在2006年
//datetime=2006-01-02

type UserInfo struct {
	Username string `json:"username" binding:"required" msg:"用户名不能为空"`
	Password string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
	Email    string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
}

func GetValidMsg(err error, obj interface{}) string {
	// obj为结构体指针
	getObj := reflect.TypeOf(obj)
	// 断言为具体的类型，err是一个接口
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg") //错误信息不需要全部返回，当找到第一个错误的信息时，就可以结束
			}
		}
	}
	return err.Error()
}

// 如果用户名不等于fengfeng就校验失败
func signValid(fl validator.FieldLevel) bool {
	if name, ok := fl.Field().Interface().(string); ok {
		return name == "fengfeng"
	}

	return false
}

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		type UserInfo struct {
			Name string `json:"name" binding:"sign" msg:"用户名错误"`
			Age  int    `json:"age" binding:""`
		}
		var user UserInfo
		err := c.ShouldBindJSON(&user)
		if err != nil {
			// 显示自定义的错误信息
			msg := GetValidMsg(err, &user)
			c.JSON(200, gin.H{"msg": msg})
			return
		}
		c.JSON(200, user)
	})
	// 注册
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}
	router.Run(":80")
}
