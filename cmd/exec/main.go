package main

import (
	//"fmt"

	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {

	// parseArgs()

	app := cli.NewApp()
	app.Name = "exec"
	app.Usage = "save your command,args and execute it by alias"
	app.Version = Version
	app.Commands = []cli.Command{}
	app.Flags = []cli.Flag{
		cli.Commander{
			Name:  "add, a",
			Usage: "-a command alias args to add a command",
			Action: func(c *cli.Context) error {
				fmt.Println(c.Args().Get(0))
			},
		},
		cli.StringFlag{
			Name:  "remove, r",
			Usage: "-r alias to remove a command",
		},
	}

	// app.Commands = []cli.Command{
	// 	{
	// 		Name:    "add",
	// 		Aliases: []string{"a"},
	// 		Usage:   "add a command to cache",
	// 		Action: func(c *cli.Context) error {

	// 			return nil
	// 		},
	// 	},
	// 	{
	// 		Name:    "remove",
	// 		Aliases: []string{"r"},
	// 		Usage:   "remove a command from cache",
	// 		Action: func(c *cli.Context) error {
	// 			return nil
	// 		},
	// 	},
	// }

	sort.Sort(cli.FlagsByName(app.Flags))
	// sort.Sort(cli.CommandsByName(app.Commands))

	if len(os.Args) > 1 && os.Args[1] != "help" && os.Args[1] != "h" {
		cmdAlias := os.Args[1]
		if cmdAlias[0] != '-' {
			if cmd, isExists := Db.FindByAlias(cmdAlias); isExists {
				cmd.execute()
			} else {
				fmt.Printf("Err: command alias: %s not found, You can add it use -a\n", cmdAlias)
			}
		}
	} else {
		app.Run(os.Args)
		// fmt.Println("Err: Please provide command alias.")
	}

	//currentPath, crr := os.Getwd()
	//if crr != nil {
	//	panic(crr)
	//}
	//if len(os.Args) > 1 {
	//	cmdAlias := os.Args[1]
	//	if cmdAlias[0] != '-' {
	//		if cmd, isExists := Db.FindByAlias(cmdAlias); isExists {
	//			cmd.execute()
	//		} else {
	//			fmt.Printf("Err: command alias: %s not found, You can add it use -a\n", cmdAlias)
	//		}
	//	} else {
	//		switch cmdAlias {
	//		case "-a":
	//			if len(os.Args) < 4 {
	//				fmt.Println("Err: Please provide command and alias")
	//				return
	//			}
	//			cmd := Commander{
	//				path:    currentPath,
	//				command: os.Args[2],
	//				alias:   os.Args[3],
	//				args:    os.Args[4:],
	//			}
	//			Db.Add(cmd)
	//			break
	//		case "-r":
	//			if len(os.Args) < 3 {
	//				fmt.Println("Err: Please provide alias")
	//				return
	//			}
	//			Db.Remove(os.Args[2])
	//			break
	//		}
	//	}
	//} else {
	//	fmt.Println("Err: Please provide command alias.")
	//}

}

func parseArgs() {
	if len(os.Args) == 1 {
		displayLogo()
	} else if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "h" || os.Args[1] == "help" {
			displayLogo()
		}
	}
}
