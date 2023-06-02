@REM cl as cgo not working
@REM go env -w CC=cl.exe

cd company
cl /nologo /EHsc /c company.cpp
cl /nologo /EHsc /DNOT_LIBRARY company.cpp
move company.exe ..\company.exe
cd ..

link /nologo /DLL /OUT:structs.dll company\company.obj

gendef structs.dll
dlltool -dllname structs.dll --def structs.def --output-lib libstructs.a

go build
