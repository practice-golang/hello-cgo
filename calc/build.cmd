@REM cl as cgo not working
@REM go env -w CC=cl.exe

cd add
cl /nologo /c add.cpp >> ..\c.out
cd ..

link /nologo /DLL /OUT:mylib.dll add\add.obj

gendef mylib.dll
dlltool -d mylib.def -l libmylib.a mylib.dll

go build
