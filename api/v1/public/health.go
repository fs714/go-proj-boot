package public

import (
	"net/http"

	"github.com/fs714/go-proj-boot/pkg/utils/code"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": code.RespOk,
		"msg":    "",
		"data":   "ok",
	})
}
