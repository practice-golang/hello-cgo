@REM cl as cgo Not work ^-^p

cd c
cl /nologo /c say.c >> ..\msvc.out
cl /nologo /c sum.c >> ..\msvc.out
lib /nologo say.obj sum.obj /OUT:mylib.lib
