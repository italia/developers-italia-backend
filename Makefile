include .env

.PHONY: run build lint test up stop prod-up prod-stop

default: build

run:
	go run main.go version.go

build:
	docker build -t italia/${NAME}:${VERSION} \
	    --build-arg NAME=${NAME} \
	    --build-arg PROJECT=${PROJECT} \
	    --build-arg VERSION=${VERSION} \
	    ./

lint:
	gometalinter --install
	gometalinter --exclude=vendor --exclude=middleware ./...

test:
	go test -race "${PROJECT}"/...

up:
	docker-compose up -d

stop:
	docker-compose stop

prod-up:
	docker-compose --file=docker-compose-prod.yml up -d

prod-stop:
	docker-compose --file=docker-compose-prod.yml stop
