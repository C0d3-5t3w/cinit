.PHONY: all init build clean run install uninstall

default_target: all
.PHONY : default_target

APP=app # Change this

all: clean init build run

init:
	mkdir -p bin
build:
	cd bin && cmake .. && make
clean:
	rm -rf bin
run:
	cd bin && ./${APP}
install:
	cp bin/${APP} /usr/local/bin/${APP}
uninstall:
	rm /usr/local/bin/${APP}

