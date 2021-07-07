@echo off
setlocal

set GOOS=js
set GOARCH=wasm
go build -o main.wasm main.go

set GOOS=
set GOARCH=
go run server.go

endlocal
