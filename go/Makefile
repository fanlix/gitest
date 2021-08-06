
TARGET=hello
DIST=../dist

all: test build dist

test:
	go test

build:
	go build -o ${TARGET}

dist:
	mkdir -p ${DIST}
	cp ${TARGET} ${DIST}/

clean:
	rm -f ${TARGET}
