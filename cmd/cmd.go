package cmd

import (
	"log"
	"mangosteen/internal/database"
	"mangosteen/internal/email"
	"mangosteen/internal/jwt_helper"
	"mangosteen/internal/router"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	emailCmd := &cobra.Command{
		Use: "email",
		Run: func(cmd *cobra.Command, args []string) {
			// RunServer()
			email.Send()
		},
	}

	coverCmd := &cobra.Command{
		Use: "cover",
		Run: func(cmd *cobra.Command, args []string) {
			if err := os.Mkdir("coverage", os.ModePerm); err != nil {
				log.Println(err)
			}

			if err := exec.Command("MailHog").Start(); err != nil {
				log.Println(err)
			}

			if err := exec.Command("go", "test", "-coverprofile=coverage/coverage.out", "./...").Run(); err != nil {
				log.Fatalln(err)
			}

			if err := exec.Command("go", "tool", "cover", "-html=coverage/coverage.out", "-o", "coverage/index.html").Run(); err != nil {
				log.Fatalln(err)
			}

			http.ListenAndServe(":8001", http.FileServer(http.Dir("./coverage")))
		},
	}

	generateHMACKey := &cobra.Command{
		Use: "generateHMACKey",
		Run: func(cmd *cobra.Command, args []string) {
			bytes, err := jwt_helper.GenerateHMACKey()
			if err != nil {
				log.Fatalln("生成generateHMACKey失败")
			}

			keyPath := viper.GetString("jwt.hmac.key_path")

			if err := os.WriteFile(keyPath, bytes, os.ModePerm); err != nil {
				log.Fatalln("写入generateHMACKey失败", err)
			}

			log.Printf("生成generateHMACKey成功，并写入[%s]文件中", keyPath)
		},
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

	rootCmd.AddCommand(serverCmd, coverCmd, dbCmd, emailCmd, generateHMACKey)
	dbCmd.AddCommand(dbCreateCmd, dbMigrateCom, dbMigrateDownCom, dbMigrateCreate, dbCurd)
	rootCmd.Execute()
}

func RunServer() {
	r := router.New()
	r.Run(":8080")
}
