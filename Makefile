LDFLAGS := -ldflags '-w -s -extldflags "-static"'
BUILD_CMD := CGO_ENABLED=0 go build -ldflags '-w -s -extldflags "-static"'

install-server:
	@curl -L https://github.com/temporalio/temporal/releases/latest/download/docker.tar.gz | tar -xz --strip-components 1 docker/docker-compose.yml

run-server:
	docker-compose up

build-start:
	$(BUILD_CMD) -o dist/start start/main.go

build-worker:
	$(BUILD_CMD) -o dist/worker worker/main.go

build: build-start build-worker

run-start-withdraw: build-start
	./dist/start withdraw

run-start-greeting: build-start
	./dist/start greeting

run-worker-withdraw: build-worker
	./dist/worker withdraw

run-worker-greeting: build-worker
	./dist/worker greeting
