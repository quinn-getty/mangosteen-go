# 运行

go build ./mangosteen

# 测试

go test ./test/...

# 自动安装依赖

go mod tidy

# 创建 数据库容器

docker run -d --name pg-for-go-mangosteen -e POSTGRES_USER=mangosteen -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=mangosteen_dev -e PGDATA=/var/lib/postgresql/data/pgdata -v pg-go-mangosteen-data:/var/lib/postgresql/data --network=network1 -p=5432:5432 postgres:14
