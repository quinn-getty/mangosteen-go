package cmd

import (
	"mangosteen/internal/database"
	"mangosteen/internal/router"
)

func RunServer() {
	database.Connect()
	database.CreateTables()
	defer database.Close()
	r := router.New()
	r.Run(":8080")
}
