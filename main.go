package main

import (
	"gin-stegosaurus/system/config"
	"gin-stegosaurus/system/route"
)

func main() {
	config.Init()
	config := config.GetConfig()
	r := route.NewRouter()
	go r.Run(":80")

	r.RunTLS(":"+config.GetString("server.port"), "./app/certs/157.245.57.91.crt", "./app/certs/157.245.57.91.key")
}
