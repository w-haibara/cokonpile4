out: *.go go.mod
	gofmt -w *.go
	go build -o out main.go

.PHONY: init
init:
	go mod init sample
	go mod tidy

.PHONY: test
test:
	gofmt -w *.go
	go test
