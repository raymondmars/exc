package main

import (
	"fmt"
	"os"
)

func main() {

	currentPath, crr := os.Getwd()
	if crr != nil {
		panic(crr)
	}
	if len(os.Args) > 1 {
		cmdAlias := os.Args[1]
		if cmdAlias[0] != '-' {
			if cmd, isExists := Db.FindByAlias(cmdAlias); isExists {
				cmd.execute()
			} else {
				fmt.Printf("Err: command alias: %s not found, You can add it use -a\n", cmdAlias)
			}
		} else {
			switch cmdAlias {
			case "-a", "--add":
				if len(os.Args) < 4 {
					fmt.Println("Err: Please provide command and alias")
					return
				}
				cmd := Commander{
					Path:    currentPath,
					Command: os.Args[2],
					Alias:   os.Args[3],
					Args:    os.Args[4:],
				}
				Db.Add(cmd)
				break
			case "-r", "--remove":
				if len(os.Args) < 3 {
					fmt.Println("Err: Please provide alias")
					return
				}
				Db.Remove(os.Args[2])
				break
			case "-l", "--list":
				showCurrentPathCommands(Db.GetCurrentAll())
				break
			case "-h", "--help":
				displayHelp()
				break
			case "-v", "--version":
				fmt.Printf("exec V%s\n", Version)
				break
			}
		}
	} else {
		displayHelp()
	}

}
