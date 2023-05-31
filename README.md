# Cgo practice

## Used tool
* Go 1.20
* GCC / MinGW-W64 (12.2.0 posix-seh), Unbuntu 18.04 (8.1.0)
* ~~MSVC 2017 Express~~ MSVC 2022 Community

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
    * MSVC - See [calc/build.cmd](calc/build.cmd)
```dos
./msvc_env_2022.ps1
cd calc
./build.cmd
```
* structs
    * MSVC - See [structs/build.cmd](structs/build.cmd)
    * Go is weak very much for c++ struct/class/type
```powershell
./msvc_env_2022.ps1
cd structs
./build.cmd
```
