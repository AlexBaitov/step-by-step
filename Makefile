build:
	env GOOS=linux go build

docker-build:
	docker build -t sbs .
