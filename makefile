BUILD := $(shell git describe)

default:out/example

clean:
	rm -rf ./cmd/example/buildVersion.go  out

test:
	go get -d -t ./
	go vet && go test

out/example:implementation.go cmd/example/main.go
	mkdir -p out
	echo "package main \n\nvar buildVersion = \"$(BUILD)\" " > ./cmd/example/buildVersion.go
	go build -o out/example ./cmd/example
