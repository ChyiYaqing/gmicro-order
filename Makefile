.PHONY: build tag push 

GIT_TAG := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "0.0.1")
VERSION := $(GIT_TAG)-$(shell git rev-parse --short HEAD)
USERNAME := chyiyaqing
SERVICE_NAME := order

build:
	docker build -t $(SERVICE_NAME):latest .

tag: build
	docker tag $(SERVICE_NAME):latest $(USERNAME)/$(SERVICE_NAME):v$(VERSION)

push: tag
	docker push $(USERNAME)/$(SERVICE_NAME):v$(VERSION)

all: push
