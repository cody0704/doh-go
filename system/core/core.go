package core

import (
	"github.com/cody0704/doh-go/system/config"
)

type H map[string]interface{}

func NewH() H {
	config := config.GetConfig()

	h := H{}
	h["title"] = config.GetString("server.title")

	return h
}
