service/up:
		docker compose -f docker-compose.yml up -d $(c)

service/stop:
		docker compose -f docker-compose.yml stop $(c)

service/down:
		docker-compose -f docker-compose.yml down

migration/up:
		goose postgres 'postgres://user:password@localhost:5432/dev?sslmode=disable' up

migration/down:
		goose postgres 'postgres://user:password@localhost:5432/dev?sslmode=disable' down