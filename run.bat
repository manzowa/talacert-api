@echo off
chcp 65001> nul
echo.
:: DEBUT APPLICATION
mode con cols=75 lines=25
color 0A
setlocal 
title APPLICATION SECU_CLI
cls
echo   ╔═══════════════════════════════════════════════════╗
echo   ║                     API TALACERT                  ║
echo   ╚═══════════════════════════════════════════════════╝
echo.   
:: go clean -cache
:: go clean -modcache                    
go run cmd\main.go
