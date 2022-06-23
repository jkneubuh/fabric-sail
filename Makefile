
.PHONY: all
all: build

.PHONY: build
build:
	CGO_ENABLED=0 go build -o build/bin/sail -v .

.PHONY: test
test:
	go test -v ./...

