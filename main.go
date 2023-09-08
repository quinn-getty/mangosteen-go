package main

import (
	"log"
	"mangosteen/cmd"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("env.config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	cmd.Run()
}
