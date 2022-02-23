SHELL=/bin/bash

# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
ENV=$(ONROBOT)

compile:
	@$(GOBUILD) -o build/$(ENV)/meter cmd/main.go

compile-linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o build/$(ENV)/meter-linux cmd/main.go

run:
	./build/$(ENV)/meter -config=build/$(ENV)/config.json

clean: