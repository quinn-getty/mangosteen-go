package cmd

import (
	"mangosteen/internal/database"
	"mangosteen/internal/router"

	"github.com/spf13/cobra"
)

func Run() {
	rootCmd := &cobra.Command{
		Use: "mangosteen",
	}

	serverCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			RunServer()
		},
	}

	dbCmd := &cobra.Command{
		Use: "db",
	}

	dbCreateCmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			database.CreateTables()
		},
	}

	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(dbCreateCmd)
	database.Connect()
	defer database.Close()
	rootCmd.Execute()
}

func RunServer() {
	r := router.New()
	r.Run(":8080")
}
