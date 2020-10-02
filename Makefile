.PHONY: test
test:
	clear
	go test -p 1 -count=1 -timeout 30s -coverprofile=./cover/all-profile.out -covermode=set -coverpkg=./... ./...; \
	go tool cover -html=./cover/all-profile.out -o ./cover/all-coverage.html

lint:
	golangci-lint run ./...

gen:
	go generate ./...
