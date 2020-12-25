install-server:
	@curl -L https://github.com/temporalio/temporal/releases/latest/download/docker.tar.gz | tar -xz --strip-components 1 docker/docker-compose.yml

run-server:
	docker-compose up

run-start-withdraw:
	go run start/main.go withdraw

run-start-greeting:
	go run start/main.go greeting

run-worker-withdraw:
	go run worker/main.go withdraw

run-worker-greeting:
	go run worker/main.go greeting
