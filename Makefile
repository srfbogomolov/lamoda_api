up:
		docker compose -f docker-compose.yml up -d $(c)
stop:
		docker compose -f docker-compose.yml stop $(c)
down:
		docker-compose -f docker-compose.yml down