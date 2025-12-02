package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var makefile = `.PHONY: all init build clean run install uninstall

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

`

var license = `MIT License

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
SOFTWARE.v`

var mainfile = `#include "root.h" // global header (includes libc.h which includes stdio.h, stdlib.h, etc)

#ifdef __cplusplus
extern "C" {
#endif

int main(void) {
  PH();
}

#ifdef __cplusplus
}
#endif // __cplusplus
`

var rootheader = `#pragma once
#ifndef ROOT_H
#define ROOT_H

#include "libc.h" // standard C library header

#ifdef __cplusplus
extern "C" {
#endif

void PH(void); // placeholder function

#ifdef __cplusplus
}
#endif // __cplusplus

#endif // ROOT_H
`

var cmakelists = `cmake_minimum_required(VERSION 4.0.1)
project(app C) # Change app name

set(CMAKE_C_STANDARD 23)
set(CMAKE_C_STANDARD_REQUIRED ON)
set(CMAKE_C_EXTENSIONS ON)

# Add compiler warnings
if(CMAKE_C_COMPILER_ID MATCHES "cc|gnu|gcc|clang|appleclang") # cc|gnu|gcc|clang|appleclang
	add_compile_options(-Wall -Wextra -Wpedantic)
elseif(MSVC)
	add_compile_options(/W4)
endif()

# Find raylib
# find_package(raylib REQUIRED)

# Source files
file(GLOB_RECURSE SOURCES "src/*.c")
file(GLOB_RECURSE HEADERS "src/*.h")

# Executable
add_executable(${PROJECT_NAME} ${SOURCES} ${HEADERS})

# Link raylib
# target_link_libraries(${PROJECT_NAME} raylib)

# Include directories
target_include_directories(${PROJECT_NAME} PRIVATE src)`

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

	rootHeaderPath := filepath.Join("src", "root.h")
	if err := os.WriteFile(rootHeaderPath, []byte(rootheader), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating root.h: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	genDirs()
	genFiles()
	fmt.Printf(`Project created successfully!\n
- CMakeLists.txt\n
- Makefile\n
- LICENSE\n
- src/main.c\n
- src/root.h\n
\n`)
}
