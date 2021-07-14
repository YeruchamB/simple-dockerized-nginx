# Compile and run the Go code
go_run:
	go mod tidy
	go run main.go

docker_build:
	docker build -t webserver .

docker_run:
	docker container rm -f web # Kill and remove any existing web container running
	docker run -it --rm -d -p 8080:80 --name web webserver

all: go_run docker_build docker_run
