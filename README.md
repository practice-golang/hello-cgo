# Cgo practice

## Used tool
* Go 1.17
* GCC / MinGW-W64 (10.3.0 posix-seh), Unbuntu 18.04 (8.1.0)
* MSVC 2017 Express - Not work

## Folders
* `hello` - C codes are embeded
* `hello-again` - C codes are separated

## Build
* hello - Just run `go build` or `go install`
* hello-again
    * MinGW - Use make. See `Makefile`
```powershell
cd hello-again
mingw32-make.exe
```
    * MSVC - See `build.cmd`
```dos
msvc_env.cmd
cd calc
build.cmd
```