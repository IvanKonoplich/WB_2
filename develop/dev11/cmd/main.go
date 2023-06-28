package main

import (
	"WB2/develop/dev11/internal/app"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func main() {
	getConfig()
	log.Fatal(app.InitApp(viper.GetString("port")))
}

func getConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("develop/dev11/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
