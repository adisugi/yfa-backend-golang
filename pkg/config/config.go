package config

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig() {
	viper.SetConfigName("App")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configuration")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("config error : ", err.Error())
	}
}
