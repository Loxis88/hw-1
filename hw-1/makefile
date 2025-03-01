# Переменные
APP_NAME := hw-1
BUILD_DIR := build
MAIN_PATH := ./cmd/main.go
BIN_PATH := $(BUILD_DIR)/$(APP_NAME)

# Цель по умолчанию
.PHONY: all
all: deps lint build

# Сборка приложения
.PHONY: build
build: lint
	@echo "Building application..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BIN_PATH) $(MAIN_PATH)
	@echo "Build completed: $(BIN_PATH)"

# Обновление и установка зависимостей
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy
	@go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
	@go install github.com/uudashr/gocognit/cmd/gocognit@latest
	@echo "Dependencies installed"

# Запуск приложения
.PHONY: run
run: build
	@echo "Running application..."
	@./$(BIN_PATH)

# Проверка цикломатической и когнитивной сложности
.PHONY: lint
lint:
	@echo "Running code complexity checks..."
	@gocyclo -over 10 .
	@gocognit -over 10 .
	@echo "Code complexity checks completed"

# Очистка
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean completed"
