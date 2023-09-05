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

var DB *gorm.DB

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
	Email     string `gorm:"uniqueIndex"`
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Item struct {
	ID         int
	UserID     int
	Amount     int
	HappenedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

var modules = []any{&User{}, &Item{}}

func CreateTables() {
	for _, module := range modules {
		err := DB.Migrator().CreateTable(module)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("创建 表成功")
		}

	}
}

func Migrate() {
	// 给users 增加 address
	DB.AutoMigrate(modules...)
}

func Curd() {
	// // 新增
	// user := User{Email: "7@qq.com", Phone: ""}
	// tx := DB.Create(&user)
	// log.Println(tx.RowsAffected)
	// log.Println(user)

	// // 查询
	// u2 := User{Phone: "15504473441"}
	// _ = DB.Find(&u2, 1)
	// log.Println(u2)

	// // 修改
	// u2.Phone = "15500000000"
	// tx := DB.Save(&u2)
	// if tx.Error != nil {
	// 	log.Println(tx.Error)
	// } else {
	// 	log.Println(tx.RowsAffected)
	// 	log.Println(u2)
	// }

	// //  查询多个数据
	// users := []User{}
	// DB.Offset(0).Limit(10).Find(&users)
	// log.Println(users)

	// // 删除
	// tx := DB.Delete(&User{ID: 1})
	// if tx.Error != nil {
	// 	log.Println(tx.Error)
	// } else {
	// 	log.Println(tx.RowsAffected)
	// }

}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
	log.Println("close database success")

}
