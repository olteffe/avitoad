.PHONY: clean test security build run swag

APP_NAME = avitoad
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/internal/database/migrations
DATABASE_URL = postgres://postgres:password@avitoad/postgres?sslmode=disable

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

linter:
	golangci-lint run

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

docker.build:
	docker build -t avitoad .

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.run: docker.network docker.postgres swag docker.net_http migrate.up

docker.stop:
	docker stop avitoad avitoad-postgres

docker.net_http: docker.build
	docker run --rm -d \
		--name avitoad \
		--network dev-network \
		-p 5000:5000 \
		avitoad

docker.postgres:
	docker run --rm -d \
		--name avitoad-postgres \
		--network dev-network \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=postgres \
		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

swag:
	swag init -g cmd/avitoad/main.go