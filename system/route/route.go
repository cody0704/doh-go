package route

import (
	"gin-stegosaurus/system/controllers"
	"gin-stegosaurus/system/middlewares"

	"github.com/gin-gonic/gin"
)

// NewRouter is Entrypoint
func NewRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	// this is output gin log to write log file
	// f, _ := os.Create("./log/gin-stegosaurus.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// LoadMiddleware
	middlewares.Loggin(router)

	routePath(router)

	return router
}

// This is route path
func routePath(r *gin.Engine) {
	xin := new(controllers.XinController)
	r.GET("/dns-query", xin.DNSQuery)
	r.POST("/dns-query", xin.DNSQuery)

	r.NoRoute(xin.NotFound)
}
