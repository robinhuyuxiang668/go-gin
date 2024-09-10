package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// 单位是字节， << 是左移预算符号，等价于 8 * 2^20
	// gin对文件上传大小的默认值是32MB
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/upload", singleFileUpload)
	router.POST("/uploads", multiFilesUpload)
	router.POST("/download", download)
	router.GET("/download", download)
	router.Run(":8080")
}

func singleFileUpload(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	log.Println(file.Filename, file.Size)

	dst := "./file/" + file.Filename
	// 上传文件至指定的完整文件路径   文件路径，注意要从项目根路径开始写
	c.SaveUploadedFile(file, dst)

	//第二种保存文件方式
	//file, _ := c.FormFile("file")
	//log.Println(file.Filename)
	//// 读取文件中的数据，返回文件对象
	//fileRead, _ := file.Open()
	////data, _ := io.ReadAll(fileRead)
	////fmt.Println(string(data))
	//
	//dst := "./" + file.Filename
	//// 创建一个文件
	//out, err := os.Create(dst)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer out.Close()
	//// 拷贝文件对象到out中
	//io.Copy(out, fileRead)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

}

func multiFilesUpload(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"] // 注意这里名字不要对不上了

	for _, file := range files {
		log.Println(file.Filename)
		// 上传文件至指定目录
		c.SaveUploadedFile(file, "./file/"+file.Filename)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func download(c *gin.Context) {
	//c.Header("Content-Type", "application/octet-stream")              // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
	//c.Header("Content-Disposition", "attachment; filename="+"牛逼.png") // 用来指定下载下来的文件名
	//c.Header("Content-Transfer-Encoding", "binary")                   // 表示传输过程中的编码形式，乱码问题可能就是因为它
	//c.File("./1.png")
	//log.Println("download end")

	//如果是前后端模式下，后端就只需要响应一个文件数据
	//文件名和其他信息就写在请求头中
	c.Header("fileName", "NB.png")
	c.Header("msg", "文件下载成功")
	c.File("./1.png")
	log.Println("download end")
}
