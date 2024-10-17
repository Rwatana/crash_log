up:
	docker compose up -d

down:
	docker compose down -v --remove-orphans

demo:
	docker-compose build --no-cache
	docker-compose up -d

