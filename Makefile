TOP := $(CURDIR)

# This is how we want to name the binary
BINARY=cache-proxy

VERSION := `cat $(TOP)/version`

# Setup the -ldflag option for go build
LDFLAGS=-ldflags "-X 'main.Version=${VERSION}'"

# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY}

# Cleans our project: deletes binaries
clean:
	rm -rf ./${BINARY}*
	@echo "Cleanup complete!"

# Test the project
test: clean build
	docker run --rm -it -w /go/src/mycode -v $(CURDIR):/go/src/mycode -u 1000:1000 golang:latest go test -v

.PHONY: clean install
