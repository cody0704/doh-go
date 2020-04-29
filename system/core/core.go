package core

import (
	"gin-stegosaurus/system/config"
)

type H map[string]interface{}

func NewH() H {
	config := config.GetConfig()

	h := H{}
	h["title"] = config.GetString("server.title")

	return h
}
