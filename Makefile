up:
	docker compose up -d

down:
	docker compose down

down-volumes:
	docker compose down --volumes --remove-orphans

build:
	docker compose build

logs:
	docker compose logs -f
	
clean-build:
	docker compose build --no-cache