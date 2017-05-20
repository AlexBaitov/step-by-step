build:
	env CGO_ENABLED=0 GOOS=linux go build

docker-build:
	docker build -t sbs .
