.PHONY: build tag push 

GIT_TAG := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.0.1")
VERSION := $(GIT_TAG)-$(shell git rev-parse --short HEAD)
USERNAME := chyiyaqing

build:
	docker build -t order:latest .

tag: build
	docker tag order:latest $(USERNAME)/order:v$(VERSION)

push: tag
	docker push $(USERNAME)/order:v$(VERSION)

all: push
