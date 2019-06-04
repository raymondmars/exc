package main

import (
	"github.com/fatih/color"
)

func displayLogo() {
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
	color.Yellow(logo, Version)
}
