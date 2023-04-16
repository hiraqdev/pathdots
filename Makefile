tidy:
	go mod tidy

test: tidy
	go test -race -v ./core/...  