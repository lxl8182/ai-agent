.PHONY: build run clean test help

# 变量定义
BINARY_NAME=ai-agent-server
BUILD_DIR=build
CMD_PATH=cmd/server/main.go

# 默认目标
help:
	@echo "Available commands:"
	@echo "  make build    - Build the project"
	@echo "  make run      - Run the project"
	@echo "  make test     - Run tests"
	@echo "  make clean    - Clean build artifacts"
	@echo "  make help     - Show this help"

# 构建项目
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_PATH)
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# 运行项目
run:
	@echo "Running..."
	go run $(CMD_PATH)

# 运行测试
test:
	@echo "Running tests..."
	go test -v ./...

# 清理构建文件
clean:
	@echo "Cleaning..."
	rm -rf $(BUILD_DIR)
	go clean
	@echo "Clean completed"

# 格式化代码
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 检查代码质量
lint:
	@echo "Running linter..."
	golangci-lint run

# 安装依赖
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy
