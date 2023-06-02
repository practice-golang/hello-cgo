@echo off

del *.exe *.def *.dll *.exp *.lib *.a c.out company\*.obj company\*.o /q /s > nul 2>&1
del company\*.obj company\*.o company\*.exe /q /s > nul 2>&1
