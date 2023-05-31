@REM cl as cgo not working
@REM go env -w CC=cl.exe

cd employee
cl /nologo /c employee.c
cd ..

link /nologo /DLL /OUT:structs.dll employee\employee.obj

gendef structs.dll
dlltool -dllname structs.dll --def structs.def --output-lib libstructs.a

go build
