.PHONY: help build

TMP_DIR := $(shell mktemp -d)
HOMEDIR := $(shell echo $$HOME)
DEST = "${HOMEDIR}/bin"
BINARY = "gen-name"

build: ## Build binary from golang source
	@go build -o ${TMP_DIR}/${BINARY}
	@cp ${TMP_DIR}/${BINARY} ${DEST}/${BINARY}
	@chmod +x ${DEST}/${BINARY}
	@rm -rf ${TMP_DIR}

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help