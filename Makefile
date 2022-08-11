GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
CURRENT_DIR := $(shell pwd)
GIT_TAG := $(shell git tag -l --sort=-creatordate | head -n 1)

.PHONY: all
all: help

.PHONY: lint
lint:
	golangci-lint run --disable-all -E golint,goimports,misspell

.PHONY: prepare
prepare: tidy lint

.PHONY: tidy
tidy:
	go mod tidy -v

.PHONY: plugin-template
plugin-template:
	docker run \
		--rm \
		-v ${CURRENT_DIR}/.krew.yaml:/tmp/template-file.yaml \
		rajatjindal/krew-release-bot:v0.0.43 \
		krew-release-bot template \
		--tag "${GIT_TAG}" \
		--template-file /tmp/template-file.yaml

.PHONY: help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@echo "  ${YELLOW}lint             ${RESET} Run linters via golangci-lint"
	@echo "  ${YELLOW}tidy             ${RESET} Run tidy for go module to remove unused dependencies"
	@echo "  ${YELLOW}prepare          ${RESET} Run all available checks and generators"
