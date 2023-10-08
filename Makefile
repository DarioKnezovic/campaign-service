# Docker image name
IMAGE_NAME := campaign-service

.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker run -it --rm --name $(IMAGE_NAME) -p 5060:5060 $(IMAGE_NAME)

.PHONY: hot-reload
hot-reload:
	docker run -it --rm --name $(IMAGE_NAME) -v $(PWD):/app -w /app -p 5060:5060 $(IMAGE_NAME) air

.PHONY: dev
dev:
	docker-compose up

proto-generate:
	cd proto && protoc --go_out=. --go-grpc_out=. user.proto