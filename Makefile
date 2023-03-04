service/up:
		docker compose -f docker-compose.yaml up -d $(c)

service/stop:
		docker compose -f docker-compose.yaml stop $(c)

service/down:
		docker compose -f docker-compose.yaml down

migrations/up:
		goose --dir ./migrations postgres 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up

migrations/down:
		goose --dir ./migrations postgres 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' down