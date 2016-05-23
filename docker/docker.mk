DOCKER_REGISTRY_HOST := docker.io
DOCKER_REGISTRY_PORT := 5000
DOCKER_REGISTRY := $(DOCKER_REGISTRY_HOST):$(DOCKER_REGISTRY_PORT)

DOCKER_IMAGE := cedriclam/$(notdir $(CURDIR))
VERSION := $(shell git describe --tags --always --dirty)
TAG ?= $(VERSION)

SHELL := /bin/bash
IMAGE_REPO := cedriclam

.PHONY: all
all: build test

.PHONY: build
build: Dockerfile
	[[ -n "$$SF_ENVIRONMENT" ]] && NO_CACHE=true || NO_CACHE=false; \
	docker build --no-cache=$$NO_CACHE -t $(DOCKER_IMAGE):$(VERSION) .

.PHONY: test
test:
	echo NO TEST

.PHONY: publish
publish:
	docker tag -f $(DOCKER_IMAGE):$(VERSION) $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):$(TAG)
	docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):$(TAG)

.PHONY: publish-latest
publish-latest:
	docker tag -f $(DOCKER_IMAGE):$(VERSION) $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):latest
	docker push $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):latest

.PHONY: clean
clean:
	-docker rmi $(DOCKER_IMAGE):$(VERSION) $(DOCKER_REGISTRY)/$(DOCKER_IMAGE):$(VERSION) || true
