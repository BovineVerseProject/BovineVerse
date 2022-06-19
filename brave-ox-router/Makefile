
.PHONY: build clean help doc

all: build

BASE_PATH=$(subst $(prefix),,$(shell pwd))

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ${BASE_PATH}/bin/brave-ox-web ${BASE_PATH}/web
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o ${BASE_PATH}/bin/brave-ox-scanner ${BASE_PATH}/scanner

run-scanner:
	@go run ${BASE_PATH}/scanner/main.go

run-web:
	@go run ${BASE_PATH}/web/

doc:
	@cd ${BASE_PATH}/doc && apidoc  -i ./ -o apidoc

clean:
	@if [ -d ${BASE_PATH}/bin ] ; then rm -f ${BASE_PATH}/bin/* ; fi
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"

