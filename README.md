# Cgo practice

## Used tool
* Go 1.15
* GCC 8.1.0 / MinGW-W64, Unbuntu 18.04
* MSVC - No. https://github.com/golang/go/issues/20982

## Folders
* `hello` - C codes are embeded
* `hello-again` - C codes are separated

## Build
* hello - Just run `go build` or `go install`
* hello-again - Use make. See `Makefile`
    * On MinGW on Windows, use `mingw32-make.exe`
