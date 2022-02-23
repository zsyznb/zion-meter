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
	@echo group number $(group)
	@echo user number per group $(user)
	./build/$(ENV)/meter -config=build/$(ENV)/config.json -group=$(group) -user=$(user)

clean: