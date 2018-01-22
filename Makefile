TOP := $(CURDIR)

# This is how we want to name the binary
BINARY=cache-proxy

VERSION := `cat $(TOP)/version`

# These are the values we want to pass for VERSION and BUILD
BUILD_DATE := `date +%FT%T%z`
BUILD_TAG := "$(shell git tag --contains | tail -n 1)"
COMMIT_SHORT := $(shell git log -1 --pretty=format:%h)
COMMIT_LONG := $(shell git log -1 --pretty=format:%H)


# Setup the -ldflag option for go build
LDFLAGS=-ldflags "-X 'main.Version=${VERSION}' -X 'main.BuildTag=${BUILD_TAG}' -X 'main.BuildDate=${BUILD_DATE}' -X 'main.CommitShort=${COMMIT_SHORT}' -X 'main.CommitLong=${COMMIT_LONG}'"

# Builds the project
build:
	go build ${LDFLAGS} -o ${BINARY}

# Install our project: copies binary
install:
	go install ${LDFLAGS}

# Cleans our project: deletes binaries
clean:
	rm -rf ./${BINARY}*
	@echo "Cleanup complete!"

# Test the project
test: clean build
	go test -v

.PHONY: clean install
