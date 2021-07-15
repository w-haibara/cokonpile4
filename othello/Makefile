out: *.go go.mod
	go fmt ./...
	go build -o out main.go

.PHONY: init
init:
	go mod init out
	go mod tidy

.PHONY: run
run:
	go fmt ./...
	go run ./...
