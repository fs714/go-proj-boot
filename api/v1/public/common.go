package public

import (
	"github.com/fs714/go-proj-boot/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func InitRoute(Router *gin.RouterGroup) gin.IRoutes {
	baseRoute := Router.Group("")
	{
		baseRoute.GET("health", Health)
		baseRoute.POST("/api/v1/auth/login", auth.Login)
	}

	return baseRoute
}
