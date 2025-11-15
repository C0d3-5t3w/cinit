
.PHONY: all init build clean run
APP=cinit
all: clean init build
init:
	mkdir -p bin
build:
	go build -o bin/${APP} main.go
clean:
	rm -rf bin
run:
	cd bin && ./${APP}
install:
	cp bin/${APP} /usr/local/bin/${APP}
uninstall:
	rm /usr/local/bin/${APP}

