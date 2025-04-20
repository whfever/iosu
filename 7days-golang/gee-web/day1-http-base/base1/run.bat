@echo off
title Gee Web Server
setlocal enabledelayedexpansion

:: Get current path
set "current_path=%CD%"
set "golang_path=C:\Users\Administrator\Documents\iosu\7days-golang"
set "example_path=C:\Users\Administrator\Documents\example"

:: Check which path we're in and set prompt accordingly
if "!current_path:~0,%golang_path:~0,1%!"=="!golang_path:~0,%golang_path:~0,1%!" (
    set "display_path=!current_path:%golang_path%=golang!"
    prompt $G!display_path!$G
) else if "!current_path:~0,%example_path:~0,1%!"=="!example_path:~0,%example_path:~0,1%!" (
    set "display_path=!current_path:%example_path%=doc\example!"
    prompt $G!display_path!$G
) else (
    prompt $G$P$G
)

cls
echo Starting Gee web server...
echo Server will be available at http://localhost:9999
echo Press Ctrl+C to stop the server
go run main.go