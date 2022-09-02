package middlewares

import (
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
)

// Secure is Secure Middleware.
func httpsecure() (secureConf gin.HandlerFunc) {
	secureConf = secure.New(secure.Config{
		SSLRedirect: true,
		SSLHost:     "localhost:443",
		//ContentSecurityPolicy: "default-src 'self'",
	})

	return
}
