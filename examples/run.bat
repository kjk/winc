del *.syso
del *.exe
go-winres make
@rem go build -ldflags="-H windowsgui" -o app.exe
go build -o app.exe
.\app.exe
