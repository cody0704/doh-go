package config

import (
	"github.com/spf13/viper"

	"log"
)

var config *viper.Viper

func Init() {
	var err error
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName("config")
	v.AddConfigPath("./app/config/")
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	config = v
}

func GetConfig() *viper.Viper {
	return config
}
