
DEV_TAG ?= dev

DOCKER_HOSTNAME ?= registry.github.com/vipindasvg/cryptoserver
# dev images
DEV_BACKEND_IMAGE ?= $(DOCKER_HOSTNAME)/backend:$(DEV_TAG)

.ONESHELL:
build_backend_binary: main.go
	go mod tidy
	@go build -ldflags "-s -w" -o build/backend

#dev
build_dev_backend_image: build_backend_binary
	docker build -f Dockerfile -t $(DEV_BACKEND_IMAGE) .


.ONESHELL:
run_dev_compose:
	docker-compose up --build --force-recreate
