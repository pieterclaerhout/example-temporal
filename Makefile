install-server:
	@curl -L https://github.com/temporalio/temporal/releases/latest/download/docker.tar.gz | tar -xz --strip-components 1 docker/docker-compose.yml

run-server:
	docker-compose up

run-start:
	go run start/main.go

run-worker:
	go run worker/main.go
