package route

import (
	"github.com/cody0704/doh-go/system/controllers"
	"github.com/cody0704/doh-go/system/middlewares"

	"github.com/gin-gonic/gin"
)

// NewRouter is Entrypoint
func NewRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	// this is output gin log to write log file
	//	config := config.GetConfig()
	//	filename := config.GetString("log.path")
	//	permissions := config.GetUint32("log.permissions")
	//	f, _ := os.OpenFile("./log/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(permissions))
	//	gin.DefaultWriter = io.MultiWriter(f)

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
