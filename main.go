package main

import (
	"github.com/cody0704/doh-go/system/config"
	"github.com/cody0704/doh-go/system/route"
)

func main() {
	config.Init()
	config := config.GetConfig()
	r := route.NewRouter()
	// go r.Run(":80")

	r.RunTLS(":"+config.GetString("server.port"), "./app/certs/127.0.0.1.crt", "./app/certs/127.0.0.1.key")
}
