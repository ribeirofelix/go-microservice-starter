default: docker

build:
	go build -o dist/server

run: build
	./dist/server

test: test-all

test-short:
	ENV=test go test -short -v $(GOPACKAGES)

test-all: build-test
	docker-compose run test

build-test:
	docker-compose build test

docker: kill app
	docker-compose up web

app: .
	docker-compose build web

kill:
	-docker-compose kill