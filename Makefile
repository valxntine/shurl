run:
	docker compose up --build --no-recreate -d client api

run-api:
	docker compose up -d api

down:
	docker compose down

clean:
	docker system prune
