package auth

import (
	"net/http"

	"github.com/fs714/go-proj-boot/pkg/utils/code"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": code.RespOk,
		"msg":    "",
		"data":   "",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": code.RespOk,
		"msg":    "",
		"data":   "",
	})
}

func ChangePassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": code.RespOk,
		"msg":    "",
		"data":   "",
	})
}

func ResetPassword(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": code.RespOk,
		"msg":    "",
		"data":   "",
	})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": code.RespOk,
		"msg":    "",
		"data":   "",
	})
}
