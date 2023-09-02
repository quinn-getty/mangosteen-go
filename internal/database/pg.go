package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lib/pq"
)

const (
	host     = "pg-for-go-mangosteen"
	port     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)

func Connect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connect db success")
}

func CreateTables() {
	// 创建 users 表
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("create users table success")
}

func handleError(err error, successMsg string) {
	if err != nil {
		log.Println(err)
	} else if successMsg != "" {
		fmt.Println(successMsg)
	}
}

func Migrate() {
	_, err := DB.Exec(`ALTER TABLE users ADD COLUMN address VARCHAR(200)`)
	handleError(err, "users add address success")

	_, err = DB.Exec(`ALTER TABLE users ADD COLUMN phone VARCHAR(50)`)
	handleError(err, "users add phone success")

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		amount INT NOT NULL,
		happened_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`)
	handleError(err, "create items table success")

	// 给users 的email 增加唯一索引
	_, err = DB.Exec(`CREATE UNIQUE INDEX user_email_index ON users (email)`)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("给users 的 email 添加唯一索引成功")
	}
}

func Curd() {
	// 新增
	_, err := DB.Exec(`INSERT INTO users (email) values ('1@qq.com')`)
	if err != nil {
		switch err.(type) {
		case *pq.Error:
			pqErr := err.(*pq.Error)
			log.Println(pqErr.Code.Name())
			log.Println(pqErr.Message)
		default:
			log.Println(err)
		}

	} else {
		log.Println("create a user success")

	}

	// // 更新
	// _, err = DB.Exec(`Update users SET phone = 15504473441 where email = '1@qq.com'`)
	// handleError(err, "update success")

	//  // 条件查询
	// result, err := DB.Query(`SELECT phone FROM users where email = '1@qq.com'`)
	// log.Println(result)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for result.Next() {
	// 		var phone string
	// 		result.Scan(&phone)
	// 		log.Println(phone)
	// 	}
	// 	log.Println("Select success")
	// }

	// // 分页查询 （更安全）
	// stmt, err := DB.Prepare(`SELECT phone FROM users where email = $1 offset $2 limit $3`)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// result, err := stmt.Query("1@qq.com", 0, 3)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	for result.Next() {
	// 		var phone string
	// 		result.Scan(&phone)
	// 		log.Println(phone)
	// 	}
	// 	log.Println("Select success")
	// }

	// // 删除
	// _, err = DB.Exec(`DELETE FROM users WHERE email = '1@qq.com'`)
	// handleError(err, "delete success")
}

func Close() {
	err := DB.Close()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("close database success")

}
