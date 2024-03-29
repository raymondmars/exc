package main

import (
	"fmt"
)

//to display help doc
func displayHelp() {

	var helpDoc = `Exc is a tool for saving you command,args in different directory and execute it by alias

Usage:
    exc [alias]  - to execute you command
    exc [command] [arguments...] - to do other operations
 
Version:
    %s
 
Commands:
    --add value, -a value        to add a command named with it's alias to cache (The value should be: alias command [args])
    --remove value, -r value     to remove a command from cache (The value should be: alias)
    --list, -l                   to show command list, it will depend on you current directory
    --list-all, -la              to show all command list in all directory
    --help, -h                   to show help
    --version, -v                to print the version

Examples:
    exc -a gs git status        set command "git status" as "gs", after it, you can execute it use "exc gs"   
    exc -r gs                   remove gs alias from cache   
`
	fmt.Printf(helpDoc, Version)
	fmt.Println("\nSee more detail: https://github.com/0RaymondJiang0/exc")
}
