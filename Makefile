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

run-start-transfer: build-start
	./dist/start transfer

run-start-greeting: build-start
	./dist/start greeting

run-start-cron: build-start
	./dist/start cron

run-worker-transfer: build-worker
	./dist/worker transfer

run-worker-greeting: build-worker
	./dist/worker greeting

run-worker-cron: build-worker
	./dist/worker cron
