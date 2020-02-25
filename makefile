BUILD := $(shell git describe)

default:out/example

clean:
	rm -rf out

test:
	go vet && go test

out/example:implementation.go cmd/example/main.go
	mkdir -p out
	echo "package main \n\nvar buildVersion = \"$(BUILD)\" \n" > ./cmd/example/buildVersion.go
	go build -o out/example ./cmd/example
