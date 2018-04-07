package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HiPage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "first code",
	})
}
