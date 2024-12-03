# Пути
CONFIG_FILE := config.yml
MIGRATION_PATH := internal/db/migrations

# Чтение конфигурации из YAML-файла
DB_HOST := $(shell yq e .dbconfig.host $(CONFIG_FILE))
DB_PORT := $(shell yq e .dbconfig.port $(CONFIG_FILE))
DB_USER := $(shell yq e .dbconfig.user $(CONFIG_FILE))
DB_PASSWORD := $(shell yq e .dbconfig.password $(CONFIG_FILE))
DB_NAME := $(shell yq e .dbconfig.dbname $(CONFIG_FILE))
DB_SSLMODE := $(shell yq e .dbconfig.sslmode $(CONFIG_FILE))

# Полный URL для подключения к базе данных
DB_URL := postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

# Команда для создания миграции
migrate-create:
	migrate create -ext=sql -dir=internal/database/migrations -seq init
# Команда для выполнения миграций "up"
migrate-up:
	migrate -path=$(MIGRATION_PATH) -database "$(DB_URL)" -verbose up

# Команда для выполнения миграций "down"
migrate-down:
	migrate -path=$(MIGRATION_PATH) -database "$(DB_URL)" -verbose down

# Команда для выполнения миграций "up" на один шаг
migrate-up-step:
	migrate -path=$(MIGRATION_PATH) -database "$(DB_URL)" -verbose up 1

# Команда для выполнения миграций "down" на один шаг
migrate-down-step:
	migrate -path=$(MIGRATION_PATH) -database "$(DB_URL)" -verbose down 1

# Команда для просмотра состояния миграций
migrate-status:
	migrate -path=$(MIGRATION_PATH) -database "$(DB_URL)" -verbose version
