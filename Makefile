.PHONY: help build

TMP_DIR := $(shell mktemp -d)
HOMEDIR := $(shell echo $$HOME)
DEST = "${HOMEDIR}/bin"
BINARY = "gen-name"

build: ## Build binary from golang source
	@echo "🌀 Building binary ..."
	@go build -o ${TMP_DIR}/${BINARY}
	@echo "🌀 Moving binary ..."
	@cp ${TMP_DIR}/${BINARY} ${DEST}/${BINARY}
	@echo "🌀 Change mod +x ..."
	@chmod +x ${DEST}/${BINARY}
	@echo "🌀 Cleaning workdir ..."
	@rm -rf ${TMP_DIR}
	@echo "Done ✅"

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help