@rem rsrc -manifest app.manifest -ico=app.ico,add.ico,application_lightning.ico,application_edit.ico,application_error.ico -o rsrc.syso

del rsrc_windows_amd64.syso
go-winres make
go build -ldflags="-H windowsgui"
