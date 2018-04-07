package router

import (
	"github.com/chenyangguang/hundun/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init() {
	r := gin.Default()
	r.LoadHTMLGlob("tpls/*")
	v1 := r.Group("/v1")
	{
		v1.GET("/hi", handlers.HiPage)
		v1.GET("/user/:name", func(c *gin.Context) {
			uname := c.Param("uname")
			c.String(http.StatusOK, "You are the first ,%s", uname)
		})
		v1.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "HOME TITLE",
				"text":  "WELCOME TO THE EARTH!",
			})
		})
	}

	r.Run(":8080")
}
