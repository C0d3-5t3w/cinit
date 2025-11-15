package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var makefile = `
.PHONY: all init build clean run install uninstall
# Change this
APP=app

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
`

var license = `
MIT License

Copyright (c) 2025 C0D3-5T3W 

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.v
`

var mainfile = `
#include <stdio.h>

int main(void) {}

`

var cmakelists = `
cmake_minimum_required(VERSION 3.10)
project(app C) # Change app name

set(CMAKE_C_STANDARD 99)
set(CMAKE_C_STANDARD_REQUIRED ON)
set(CMAKE_C_EXTENSIONS OFF)

# Add compiler warnings
if(CMAKE_C_COMPILER_ID MATCHES "GNU|Clang")
	add_compile_options(-Wall -Wextra -Wpedantic)
elseif(MSVC)
	add_compile_options(/W4)
endif()

# Source files
file(GLOB_RECURSE SOURCES "src/*.c")
file(GLOB_RECURSE HEADERS "include/*.h")

# Include directories
include_directories(include)

# Executable
add_executable(${PROJECT_NAME} ${SOURCES} ${HEADERS})

# Add subdirectories if needed
# add_subdirectory(lib)

# Link libraries
# target_link_libraries(${PROJECT_NAME} PRIVATE your_library)
`

func genDirs() {

	err := os.MkdirAll("src", 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating src directory: %v\n", err)
		os.Exit(1)
	}
}

func genFiles() {

	if err := os.WriteFile("CMakeLists.txt", []byte(cmakelists), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating CMakeLists.txt: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile("Makefile", []byte(makefile), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Makefile: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile("LICENSE", []byte(license), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating LICENSE: %v\n", err)
		os.Exit(1)
	}

	mainPath := filepath.Join("src", "main.c")
	if err := os.WriteFile(mainPath, []byte(mainfile), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating main.c: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	genDirs()
	genFiles()
	fmt.Printf(`
		Project created successfully!
			- CMakeLists.txt
			- Makefile
			- LICENSE
			- src/main.c
	`)
}
