build: ## Build the container
	docker-compose -f docker-compose.yml build

run: ## Run container
	make build 
	docker-compose -f docker-compose.yml up -d

stop: ## Stopcontainer
	docker-compose down -v

test: ## testing
	docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down	