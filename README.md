# Cgo practice

## Tool
* Go 1.15
* GCC 8.1.0 / MinGW-W64

## Folders
* `hello` - C codes are embeded
* `hello-again` - C codes are separated

## Build
* hello - Just run `go build` or `go install`
* hello-again - Use make. See `Makefile`
    * On MinGW on Windows, use `mingw32-make.exe`
