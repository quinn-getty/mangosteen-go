# 运行

go build ./mangosteen

# 测试

go test ./test/...

# 自动安装依赖

go mod tidy

# 创建 数据库容器

docker run -d --name pg-for-go-mangosteen -e POSTGRES_USER=mangosteen -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=mangosteen_dev -e PGDATA=/var/lib/postgresql/data/pgdata -v pg-go-mangosteen-data:/var/lib/postgresql/data --network=network1 -p=5432:5432 postgres:14

# 安装使用 sqlc

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

## 解析生成包

sqlc generate

# 迁移表

## 安装全局 golang-migrate

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

## 生成迁移文件

migrate create -ext sql -dir config/migrations -seq create_users_table
或者
go build .; ./mangosteen db create:migration create_users_table

## 执行迁移

### 降级 一次

go build .; ./mangosteen db migrate:down
或者
migrate -database "postgres://mangosteen:123456@pg-for-go-mangosteen:5432/mangosteen_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" down 1

### 升级 全量

go build .; ./mangosteen db migrate
或者
migrate -database "postgres://mangosteen:123456@pg-for-go-mangosteen:5432/mangosteen_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" up

## 更新 user 表字段

migrate create -ext sql -dir config/migrations -seq add_phone_for_users
