# Cgo practice

## Used tool
* Go 1.17
* GCC / MinGW-W64 (10.3.0 posix-seh), Unbuntu 18.04 (8.1.0)
* MSVC 2017 Express

## Folders
* `hello` - C codes are embeded
* `hello-again` - C codes are separated
* `calc` - Use MSVC

## Build
* hello - Just run `go build` or `go install`
* hello-again
    * MinGW - Use make. See `hello-again/Makefile`
```powershell
cd hello-again
mingw32-make.exe
```
* calc
    * MSVC - See `calc/build.cmd`
```dos
msvc_env.cmd
cd calc
build.cmd
```
