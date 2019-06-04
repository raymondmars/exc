package main

import (
	"github.com/fatih/color"
)

func displayHelp() {
	var logo = `
███████╗██╗  ██╗███████╗ ██████╗
██╔════╝╚██╗██╔╝██╔════╝██╔════╝
█████╗   ╚███╔╝ █████╗  ██║     
██╔══╝   ██╔██╗ ██╔══╝  ██║     
███████╗██╔╝ ██╗███████╗╚██████╗
╚══════╝╚═╝  ╚═╝╚══════╝ ╚═════╝
EXEC V%s
https://github.com/0RaymondJiang0/exec
`
	var helpDoc = `
NAME:
    exec - save your command,args and execute it by alias

USAGE:
    exec [command alias]  - to execute you command
    exec [global options] [arguments...] - to do other operations
 
VERSION:
    %s
 
GLOBAL OPTIONS:
    --add value, -a value     -a command alias args to add a command
		--remove value, -r value  -r alias to remove a command
    --list, -l  							-l to show command list
    --help, -h                show help
    --version, -v             print the version
`
	color.Yellow(logo, Version)
	color.White(helpDoc, Version)
}
