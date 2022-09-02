package route

import (
	"github.com/cody0704/doh-go/system/controllers"

	"github.com/gin-gonic/gin"
)

// NewRouter is Entrypoint
func NewRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	routePath(router)

	return router
}

// This is route path
func routePath(r *gin.Engine) {
	handle := new(controllers.Controller)
	r.GET("/dns-query", handle.DNSQuery)
	r.POST("/dns-query", handle.DNSQuery)

	r.NoRoute(handle.NotFound)
}
