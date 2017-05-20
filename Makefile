build:
	env CGO_ENABLED=0 GOOS=linux go build

docker-build:
	docker build -t sbs .

docker-run:
	docker run --rm -p8000:8000 sbs:latest
