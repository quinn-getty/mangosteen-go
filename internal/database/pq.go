package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "pg-for-go-mangosteen"
	port     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	DB = db

	// connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// DB = db
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("connect db success")
}

type User struct {
	ID        int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateTables() {
	u1 := User{Email: "1.@qq.com"}
	err := DB.Migrator().CreateTable(&u1)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("创建User 表成功")
	}
}

func handleError(err error, successMsg string) {
	if err != nil {
		log.Println(err)
	} else if successMsg != "" {
		fmt.Println(successMsg)
	}
}

func Migrate() {

}

func Curd() {

}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
	log.Println("close database success")

}
