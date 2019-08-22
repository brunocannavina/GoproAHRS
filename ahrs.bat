@echo off

set scriptdir=%~dp0
set file=%1
set nombre=%~n1


IF [%1]==[] GOTO :eof
dir %file% /a:a/b/s > "%scriptdir%%nombre%.txt"
