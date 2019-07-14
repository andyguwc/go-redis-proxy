build: ## Build the container
	docker-compose build

run: ## Run container
	docker-compose up -d

stop: ## Stopcontainer
	docker-compose down -v

test: ## testing
	docker-compose up -d
	go test -v ./...
	docker-compose down