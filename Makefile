test:
	go test ./...

test-pretty:
	go test ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''| sed ''/ok/s//$(printf "\033[32mok\033[0m")/''

build:
	go build ./...


doc:
	# requires: gomarkdoc
    # go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	gomarkdoc --output '{{.Dir}}/readme.md' ./...