.PHONY: test
test:
	clear
	go test -count=1 -timeout 30s -cover ./...

race:
	go test -count=1 -race -timeout 30s ./...

cover:
	go test -count=1 -timeout 10s -coverprofile=cover-profile.out -covermode=set -coverpkg=./... ./...

cover-html: cover
	go tool cover -html=cover-profile.out -o cover-coverage.html

lint:
	golangci-lint run ./...

gen:
	go generate ./...
