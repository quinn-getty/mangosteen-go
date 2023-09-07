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

	dbMigrateCreate := &cobra.Command{
		Use: "create:migration",
		Run: func(cmd *cobra.Command, args []string) {
			database.CreateMigrate(args[0])
		},
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
			database.MigrateUp()
		},
	}

	dbMigrateDownCom := &cobra.Command{
		Use: "migrate:down",
		Run: func(cmd *cobra.Command, args []string) {
			database.MigrateDown()
		},
	}

	dbCurd := &cobra.Command{
		Use: "curd",
		Run: func(cmd *cobra.Command, args []string) {
			database.Curd()
		},
	}
	database.Connect()
	defer database.Close()

	rootCmd.AddCommand(serverCmd, dbCmd)
	dbCmd.AddCommand(dbCreateCmd, dbMigrateCom, dbMigrateDownCom, dbMigrateCreate, dbCurd)
	rootCmd.Execute()
}

func RunServer() {
	r := router.New()
	r.Run(":8080")
}
