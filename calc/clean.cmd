@echo off

del struct.exe *.def *.dll *.exp *.lib *.a c.out add\*.obj add\*.o random_float\*.obj random_float\*.o /q /s > nul 2>&1
