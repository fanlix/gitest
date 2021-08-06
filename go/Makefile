
TARGET=hello
DIST=../dist

all: test build

test:
	go test

build:
	mkdir -p ${DIST}
	go build -o ${DIST}/${TARGET}

clean:
	rm -f ${DIST}/${TARGET}
