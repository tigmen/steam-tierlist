.PHONY: build
bin:
	mkdir bin

build: bin 
	go mod tidy
	go build -o ./bin/out.exe ./cmd/app/app.go

.DEFAULT_GOAL := build

test: testdata
	go test -v -coverprofile=./testdata/coverage.out ./...

testdata:
	mkdir testdata

cover: testdata/coverage.out
	go tool cover -html=testdata/coverage.out