package main

import (
	// "fmt"
	// "github.com/fvbock/endless"
	// "github.com/gin-gonic/gin"
	"github.com/chenyangguang/hundun/router"
	// "io"
	// "log"
	// "net/http"
	// "os"
	// "time"
)

func main() {
	router.Init()
	// add gin
	// gin.DisableConsoleColor()
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// router := gin.Default()
	// router.GET("/hi", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "hello world",
	// 	})
	// })

	// router.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// })

	// router.GET("/user/:name/*action", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	action := c.Param("action")
	// 	msg := name + " is " + action
	// 	c.String(http.StatusOK, msg)
	// })

	// router.POST("/form_post", func(c *gin.Context) {
	// 	msg := c.PostForm("msg")
	// 	name := c.PostForm("name")
	// 	nick := c.DefaultPostForm("nick", "anonymous")

	// 	c.JSON(200, gin.H{
	// 		"status": "posted",
	// 		"msg":    msg,
	// 		"name":   name,
	// 		"nick":   nick,
	// 	})
	// })

	// router.POST("/upload", func(c *gin.Context) {
	// 	id := c.Query("id")
	// 	page := c.DefaultQuery("page", "0")
	// 	name := c.PostForm("name")
	// 	msg := c.PostForm("msg")

	// 	file, _ := c.FormFile("file")

	// 	log.Println(file.Filename)

	// 	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!"), file.Filename)

	// 	fmt.Printf("id: %s; page %s; name %s; msg %s", id, page, name, msg)
	// })

	// // add multiple file
	// router.POST("/upload2", func(c *gin.Context) {
	// 	form, _ := c.MultipartForm()
	// 	files := form.File["files[]"]

	// 	for _, file := range files {
	// 		log.Println(file.Filename)
	// 	}

	// 	c.String(http.StatusOK, fmt.Sprintf("%d files  uploaded!"), len(files))
	// })

	// // goroutines inside a middlleware
	// router.GET("long_async", func(c *gin.Context) {
	// 	cCp := c.Copy()
	// 	go func() {
	// 		time.Sleep(5 * time.Second)
	// 		fmt.Println("done ! in path " + cCp.Request.URL.Path)
	// 		//			log.Println("Done ! in path " + cCp.Request.URL.Path)
	// 	}()
	// })

	// //	router.Run(":8080") // localhost:8080 default

	// //  无缝重启
	// endless.ListenAndServe(":8080", router)

	// v1 := router.Group("v1")
	// {
	// 	v1.POST("login", loginEndpoint)
	// 	v1.POST("submit", submitEndpoint)
	// 	v1.POST("read", readEndpoint)
	// }
	// v2:=router.Group("v2")
	// {
	// 	v2.POST("login", loginEndpoint)
	// 	v2.POST("submit", submitEndpoint)
	// 	v2.POST("read", readEndpoint)
	//	}
}
