
TARGET=hello
DIST=../dist

all: test build dist

test:
	go test

build:
	CGO_ENABLED=0 go build -ldflags "-w -s" -o ${TARGET}

dist:
	mkdir -p ${DIST}
	cp ${TARGET} ${DIST}/

clean:
	rm -f ${TARGET}
