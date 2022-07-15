package auth

import "github.com/gin-gonic/gin"

func InitRoute(Router *gin.RouterGroup) gin.IRoutes {
	baseRoute := Router.Group("/api/v1/auth")
	{
		baseRoute.POST("/register", Register)
		baseRoute.POST("/password/change", ChangePassword)
		baseRoute.POST("/password/reset", ResetPassword)
		baseRoute.POST("/logout", Logout)
	}

	return baseRoute
}
