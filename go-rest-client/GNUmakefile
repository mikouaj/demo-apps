GOCMD=go
TEST?=$$(go list ./... |grep -v 'vendor')
BINARY=go-rest-client

default: build test

build:
	CGO_ENABLED=0 ${GOCMD} build -o ${BINARY}

test:
	echo $(TEST) | \
		xargs -t ${GOCMD} test -v

clean:
	${GOCMD} clean

.PHONY: test clean