.PHONY: build
bin:
	mkdir bin

build: bin 
	go mod tidy
	go build -o ./bin/out.exe ./cmd/app/app.go

.DEFAULT_GOAL := build
