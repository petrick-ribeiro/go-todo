build:
	@go build -o bin/gotodo

run: build
	@./bin/gotodo

up:
	@docker-compose up -d --build
down:
	@docker-compose down