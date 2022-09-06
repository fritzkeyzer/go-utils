test:
	go test ./...

build:
	go build ./...

# requires: gomarkdoc
# go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
doc:
	#gomarkdoc --output '{{.Dir}}/readme.md' ./...
	(cd logpage && gomarkdoc -o readme.md)
	(cd env && gomarkdoc -o readme.md)
	(cd pretty && gomarkdoc -o readme.md)
	(cd stacks && gomarkdoc -o readme.md)
	(cd stringutil && gomarkdoc -o readme.md)