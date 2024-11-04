GOOSE_DBSTRING ?= "root:123456@tcp(127.0.0.1:3308)/GO_ECOMMERCE"
GOOSE_MIGRATION_DIR ?= sql/schema
GOOSE_DRIVER ?= mysql


# name app
APP_NAME = server

# run:
# 	go run ./cmd/$(APP_NAME)/

# build:
# 	go build -o bin/migration_app ./cmd/$(APP_NAME)/

# dev:
# 	 docker-compose up && go run ./cmd/$(APP_NAME)

# kill:
# 	docker-compose kill

# up:
# 	docker-compose up -d

# down:
# 	docker-compose down

start:
	 docker-compose up && go run ./cmd/$(APP_NAME)

docker_build:
	docker-compose up -d --build
	docker-compose ps

docker_stop:
	docker-compose down

docker_up:
	docker-compose up

dev:
	go run ./cmd/$(APP_NAME)

upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up

downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE MIGRATION_DIR) down

resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE MIGRATION_DIR) reset

sqlgen:
	sqlc generate

.PHONY: start dev downse upse resetse docker_build docker_stop docker_up