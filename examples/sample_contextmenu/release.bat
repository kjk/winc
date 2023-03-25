rsrc -manifest app.manifest -ico=app.ico,application_lightning.ico,application_edit.ico,application_error.ico -o rsrc_windows_amd64.syso
go build -ldflags="-H windowsgui"
