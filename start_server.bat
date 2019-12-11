set path=%path%;%cd%\dlls
taskkill /F /IM go-bioserver.exe
go-bioserver.exe
pause