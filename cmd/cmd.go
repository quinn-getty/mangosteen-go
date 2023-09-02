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

	dbMigrateCom := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			database.Migrate()
		},
	}

	dbCurd := &cobra.Command{
		Use: "curd",
		Run: func(cmd *cobra.Command, args []string) {
			database.Curd()
		},
	}

	rootCmd.AddCommand(serverCmd, dbCmd)

	dbCmd.AddCommand(dbCreateCmd, dbMigrateCom, dbCurd)

	database.Connect()
	defer database.Close()
	rootCmd.Execute()
}

func RunServer() {
	r := router.New()
	r.Run(":8080")
}
