build:
	env CGO_ENABLED=0 GOOS=linux go build

docker-build:
	docker build -t sbs .

docker-run:
	docker run --rm -p8000:8000 sbs:latest

build-and-run-in-docker: build
build-and-run-in-docker: docker-build
build-and-run-in-docker: docker-run
