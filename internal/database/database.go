package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	host     = "pg-for-go-mangosteen"
	port     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

var DB *sql.DB

func Connect() {
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
}

func CreateMigrate(filename string) {
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "config/migrations", "-seq", filename)
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("创建迁移文件%s成功", filename)
}
func CreateTables() {

}

func MigrateUp() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", pwd),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname),
	)
	if err != nil {
		log.Fatalln(err)
	}

	err = m.Up()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("迁移成功")
}

func MigrateDown() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", pwd),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname),
	)
	if err != nil {
		log.Fatalln(err)
	}

	err = m.Steps(-1)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("已会退一个版本！")
}

func Curd() {
}

func Close() {

}
