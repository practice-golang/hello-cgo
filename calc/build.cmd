@REM cl as cgo not working
@REM go env -w CC=cl.exe

cd add
cl /nologo /c add.cpp
cd ..
cd random_float
cl /nologo /c random_float.cpp
cd ..
cd random_int
cl /nologo /c random_int.cpp
cd ..

link /nologo /DLL /OUT:calc.dll add\add.obj random_float\random_float.obj random_int\random_int.obj

gendef calc.dll
dlltool -dllname calc.dll --def calc.def --output-lib libcalc.a

go build
