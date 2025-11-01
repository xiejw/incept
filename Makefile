GO_MOD = github.com/xiejw/incept

run:
	go run cmd/main.go

fmt:
	go fmt ${GO_MOD}/...
