package main

import (
	"log"
	"mangosteen/cmd"
	"mangosteen/config"

	"github.com/spf13/viper"
)

func main() {
	config.LoadConfig()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	cmd.Run()
}
