package api

import (
	"github.com/fs714/go-proj-boot/api/v1/auth"
	"github.com/fs714/go-proj-boot/api/v1/public"
	"github.com/fs714/go-proj-boot/pkg/utils/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Config.Common.RunMode)
	gin.DisableConsoleColor()
	r := gin.New()
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/api/v1/health"},
	}))
	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithDecompressFn(gzip.DefaultDecompressHandle)))
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	if config.Config.Common.Profiling {
		pprof.Register(r)
	}

	// no authentication
	publicGroup := r.Group("")
	{
		public.InitRoute(publicGroup)
	}

	// authentication
	privateGroup := r.Group("")
	{
		auth.InitRoute(privateGroup)
	}

	return r
}
