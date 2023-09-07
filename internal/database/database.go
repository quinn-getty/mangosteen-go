package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"mangosteen/config/queries"
	"math/rand"
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
var DBCtx = context.Background()

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("数据库连接成功")
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
	// 增加
	q := queries.New(DB)
	email := fmt.Sprintf("%d@qq.com", rand.Int())
	user, err := q.CreateUser(DBCtx, email)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(user)
	}

	// 更新
	err = q.UpdateUser(DBCtx, queries.UpdateUserParams{
		ID:      user.ID,
		Email:   user.Email,
		Phone:   user.Phone,
		Address: "中国四川成都",
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println("更新成功")
	}

	// 查询列表
	userList, err := q.ListUsers(DBCtx, queries.ListUsersParams{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		log.Println(err)
	} else {
		log.Println(userList)
	}

	err = q.DeleteUser(DBCtx, userList[0].ID)
	if err != nil {
		log.Println(err)
	}
	log.Println("删除成功")

}

func Close() {

}
