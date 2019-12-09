set path=%path%;%cd%\dlls
taskkill /F /IM go-bioserver.exe
start go-bioserver.exe