package main

import (
	"flag"

	"github.com/cody0704/doh-go/system/route"
)

var addr, cert, key string

func main() {
	flag.StringVar(&addr, "addr", ":443", "listen address")
	flag.StringVar(&cert, "cert", "./app/certs/192.168.0.1.crt", "cert")
	flag.StringVar(&key, "key", "./app/certs/192.168.0.1.key", "key")
	flag.Parse()

	r := route.NewRouter()

	r.RunTLS(addr, cert, key)
}
