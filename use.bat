@echo off
for /f %%i in ('dir /B /A:D') do (
    echo go work use %%i
    go work use %%i
)