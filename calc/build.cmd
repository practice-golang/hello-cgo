@REM cl as cgo not working
@REM go env -w CC=cl.exe

cd add
cl /nologo /c add.cpp
cd ..

link /nologo /DLL /OUT:adder.dll add\add.obj

gendef adder.dll
dlltool -dllname adder.dll --def adder.def --output-lib libadder.a

go build
