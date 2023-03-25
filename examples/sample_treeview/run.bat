del *.syso
del *.exe
go-winres make
go build -ldflags="-H windowsgui" -o app.exe
.\app.exe
