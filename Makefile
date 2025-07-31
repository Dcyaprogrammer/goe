.PHONY: build run test clean deps

# 构建项目
build:
	go build -o goe main.go

# 运行项目
run:
	go run main.go

# 运行测试
test:
	go test ./...

# 清理构建文件
clean:
	rm -f goe

# 下载依赖
deps:
	go mod tidy
	go mod download

# 格式化代码
fmt:
	go fmt ./...

# 代码检查
lint:
	golangci-lint run

# 安装依赖
install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# 开发模式运行（自动重载）
dev:
	air

# 数据库迁移
migrate:
	go run main.go migrate 