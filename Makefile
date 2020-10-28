GO_VERSION=1.15.3

.PHONY: git-init
git-init:
	git config core.hooksPath .githooks

.PHONY: build-docker-plugin
build-docker-plugin:
	@docker run --rm -v ${PWD}:/project -w /project golang:${GO_VERSION}-alpine3.12 \
 	go build \
	-o ./bin/docker.1s \
 	./docker/

.PHONY: vendor
vendor:
	@docker run --rm -v ${PWD}:/project -w /project golang:${GO_VERSION}-alpine3.12 go mod tidy
	@docker run --rm -v ${PWD}:/project -w /project golang:${GO_VERSION}-alpine3.12 go mod vendor
