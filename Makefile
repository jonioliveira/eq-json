.DEFAULT_GOAL := help

# Environment variables used by commands called from make (cannot be make variables).
export PROJECT_NAME?=eq-json
export LC_ALL=en_US.UTF-8
export PROJECT_ROOT=$(shell pwd)
export PATH=$(shell (echo "$$(go env GOPATH 2> /dev/null)/bin:" || echo ""))$(shell echo $$PATH)
export REPOSITORY?=jonioliveira/eq-json
export VERSION?=dev-latest

DOCKER_LOCAL_IMAGE=$(REPOSITORY):dev-local
DOCKER_PRECOMMIT_LOCAL_IMAGE=$(REPOSITORY)-precommit:dev-local
DOCKER_PRECOMMIT_BUILD=docker build -f build/ci/pre-commit/Dockerfile --tag $(DOCKER_PRECOMMIT_LOCAL_IMAGE) .
DOCKER_PRECOMMIT_RUN=docker run -t -v $$PROJECT_ROOT:/pre-commit $(DOCKER_PRECOMMIT_LOCAL_IMAGE)
DOCKER_DEV_BUILD=docker build -f build/package/Dockerfile --target development --tag $(DOCKER_LOCAL_IMAGE) --build-arg VERSION .
DOCKER_RUN_BASE=docker run --rm -v $$PROJECT_ROOT:/opt/app/ -v $$PROJECT_ROOT/.cache/:/.cache/ -p $(PORT):8080 -e GOCACHE=/.cache/go-build -e GOLANGCI_LINT_CACHE=/.cache/golangci-lint
DOCKER_DEV_RUN=$(DOCKER_RUN_BASE) $(DOCKER_LOCAL_IMAGE)
DOCKER_DEV_RUN_IT=$(DOCKER_RUN_BASE) -it $(DOCKER_LOCAL_IMAGE)

## General

# target: help - Display available recipes.
.PHONY: help
help:
	@egrep "^# target:" [Mm]akefile

## GIT

# Get the latest tag of the commit history.
.PHONY: git-latest-tag
git-latest-tag:
	@git describe --abbrev=0 --tags

# Get the tag of the current commit.
.PHONY: git-commit-tag
git-commit-tag:
	@git describe --exact-match --tags HEAD

# Get the version tag of the current commit or default to 'dev-latest'.
.PHONY: git-version-tag
git-version-tag:
	@(echo ${VERSION})


## Shell

# Lint all go files.
# No need to call `gofmt` since `golint` will be called.
.PHONY: shell-go-lint
shell-go-lint:
	golangci-lint run

# Lint and fix (if possible) all go files.
.PHONY: shell-go-fix
shell-go-fix:
	go fmt ./...
	golangci-lint run --fix
	go mod tidy

# Build the go binary.
.PHONY: shell-go-build
shell-go-build:
	go build ${GO_LDFLAGS} -o bin/$(PROJECT_NAME) cmd/$(PROJECT_NAME)/main.go

# Run tests.
.PHONY: shell-go-test
shell-go-test:
	go test ${GO_LDFLAGS} $(GO_TEST_FLAGS) ./...

# Run the app.
.PHONY: shell-go-run
shell-go-run:
	go run ${GO_LDFLAGS} cmd/$(PROJECT_NAME)/main.go $(filter-out $@,$(MAKECMDGOALS))

# Clean the cache.
.PHONY: shell-clean-cache
shell-clean-cache:
	rm -Rf $$PROJECT_ROOT/.cache

## Docker

# target: docker-sh - Run a sh shell inside the container.
.PHONY: docker-sh
docker-sh:
	$(DOCKER_DEV_BUILD)
	$(DOCKER_DEV_RUN_IT) sh

# target: docker-precommit-build - Build the precommit image inside the container
.PHONY: docker-precommit-build
docker-precommit-build:
	$(DOCKER_PRECOMMIT_BUILD)

# target: docker-precommit-install - install the precommit hooks in the image inside the container
.PHONY: docker-precommit-install
docker-precommit-install:
	$(DOCKER_PRECOMMIT_RUN) install

# target: docker-precommit-autoupdate - autoupdate the precommit hooks in the image inside the container
.PHONY: docker-precommit-autoupdate
docker-precommit-autoupdate:
	$(DOCKER_PRECOMMIT_RUN) autoupdate

# target: docker-precommit-clean - clean the precommit hooks in the image inside the container
.PHONY: docker-precommit-clean
docker-precommit-clean:
	$(DOCKER_PRECOMMIT_RUN) clean

# target: docker-precommit-run - run the precommit hooks in the image inside the container
.PHONY: docker-precommit-run
docker-precommit-run:
	$(DOCKER_PRECOMMIT_RUN) run --all-files

# Run the linter inside the container.
.PHONY: docker-lint-app
docker-lint-app:
	$(DOCKER_DEV_BUILD)
	$(DOCKER_DEV_RUN) make shell-go-generate-lint

# Run the linter and fix (if possible) inside the container.
.PHONY: docker-fix-app
docker-fix-app:
	$(DOCKER_DEV_BUILD)
	$(DOCKER_DEV_RUN) make shell-go-fix-generate

# Build the app inside the container.
.PHONY: docker-build-app
docker-build-app:
	$(DOCKER_DEV_BUILD)
	$(DOCKER_DEV_RUN) make shell-go-build

# Run app tests inside the container.
.PHONY: docker-test-app
docker-test-app:
	$(DOCKER_DEV_BUILD)
	$(DOCKER_DEV_RUN) make GO_TEST_FLAGS="$(GO_TEST_FLAGS)" shell-go-generate-test

# Run the app inside the container.
.PHONY: docker-run-app-only
docker-run-app-only:
	$(DOCKER_DEV_BUILD)
	$(DOCKER_DEV_RUN_IT) make shell-go-run $(filter-out $@,$(MAKECMDGOALS))

# Build the app container image.
.PHONY: docker-build
docker-build-image:
	docker build -f build/package/Dockerfile --target production --tag $(REPOSITORY):${VERSION} .

# Push the app container image.
.PHONY: docker-push
docker-push:
	docker push ${REPOSITORY}:${VERSION}

# Delete the container image and its assets.
.PHONY: docker-clean
docker-clean:
	docker rmi -f $(DOCKER_LOCAL_IMAGE)

## Alias

# target: app-version - Get the current app version.
.PHONY: app-version
app-version: git-version-tag

# target: lint - Run the linter (inside the container).
.PHONY: lint
lint: docker-lint-app

# target: fix - Run the linter and fix issues if possible (runs inside the container). Good to be called on file save in an IDE.
.PHONY: fix
fix: docker-fix-app

# target: build - Build the app (inside the container).
.PHONY: build
build: shell-go-build

# target: test - Run app tests (inside the container).
.PHONY: test
test: docker-test-app

# target: run - Run the app {inside the container}.
.PHONY: run
run:
	go run ${GO_LDFLAGS} cmd/$(PROJECT_NAME)/main.go $(filter-out $@,$(MAKECMDGOALS))

# target: clean - Clean cache and local docker image.
.PHONY: clean
clean: shell-clean-cache docker-clean

pre-commit-install: docker-precommit-build docker-precommit-install docker-precommit-autoupdate

pre-commit: docker-precommit-run
