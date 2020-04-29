package middlewares

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

// Loggin Middleware
func Loggin(r *gin.Engine) {
	r.Use(httpsecure())
}

// Secure is Secure Middleware.
func httpsecure() (secureConf gin.HandlerFunc) {
	secureConf = secure.New(secure.Config{
		SSLRedirect: true,
		SSLHost:     "localhost:443",
		//ContentSecurityPolicy: "default-src 'self'",
	})

	return
}
