test:
	go test github.com/fritzkeyzer/go-utils/...

build:
	go work sync
	go work use .
	go env GOWORK
	go build github.com/fritzkeyzer/go-utils/...