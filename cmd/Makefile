GO_FILES = $(shell go list ./...)

test:
	go test $(GO_FILES) -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out
