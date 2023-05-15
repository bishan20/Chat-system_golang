DB_URL = postgresql://root:secret@localhost:5432/chat_system?sslmode=disable

sqlc:
	sqlc generate

migcrt:
	goose -dir ./db/migration/ create initial_schema sql

gooseup:
	goose -dir ./db/migration/ -v postgres "$(DB_URL)" up

goosedown:
	goose -dir ./db/migration/ -v postgres "$(DB_URL)" down

images:
	sudo docker images

container:
	sudo docker ps

container-all:
	sudo docker ps -a

rm:
	sudo docker rm chat-system_golang_api_1

rmi:
	sudo docker rmi chat-system_golang_api

docker:
	sudo docker-compose up

postgres:
	sudo docker start chat-system_golang_postgres_1

go:
	go run main.go


