@echo off

echo Processing...

set scriptdir=%~dp0
set nombre=%~n1
set nomext=%~nx1

IF [%1]==[] GOTO :eof
IF EXIST "%scriptdir%\%~n1-ahrs.csv" GOTO :eof

echo Wait please...

START "" /WAIT /MIN "%scriptdir%\bin\ahrs" -i "%nomext%"
