DIR=$(shell pwd)
PROJECT_BIN=${DIR}/bin
GOLANGCI_LINT = ${PROJECT_BIN}/golangci-lint

.PHONY: .install-linter
.install-linter:
	### INSTALL GOLANGCI-LINT ###
	[ -f $(PROJECT_BIN)/golangci-lint ] || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(PROJECT_BIN) v1.50.1

.PHONY: lint
lint: .install-linter
	### RUN GOLANGCI-LINT ###
	$(GOLANGCI_LINT) run ./...


## run: run project
.PHONY: run
run:
	$(PROJECT_BIN)/main

.PHONY: go-build
go-build:
	go build -o $(PROJECT_BIN) ./main.go

.PHONY: docker-run
docker-run:
	make go-build && sudo docker-compose up --build funnytgbot 
