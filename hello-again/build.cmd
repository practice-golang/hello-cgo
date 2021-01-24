@REM cl as cgo Not work ^-^p

go env -w CC=cl.exe

cd c
cl /nologo /c say.c >> ..\c.out
cl /nologo /c sum.c >> ..\c.out
cd ..

lib /nologo c/say.obj c/sum.obj /OUT:lib/mylib.lib
@REM move mylib.lib lib/mylib.lib

go build ./go/hello-again.go
