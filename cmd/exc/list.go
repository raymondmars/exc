package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// to print commands only in current directory
func printCommands(cmds []Commander) {
	const format = "%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Alias", "Command", "Args")
	fmt.Fprintf(tw, format, "-----", "-------", "----")
	for _, t := range cmds {
		fmt.Fprintf(tw, format, t.Alias, t.Command, strings.Join(t.Args, " "))
	}
	tw.Flush() // calculate column widths and print table
}

// to print all commands in all directory
func printAllCommands(cmds []Commander) {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Path", "Alias", "Command", "Args")
	fmt.Fprintf(tw, format, "-----", "-----", "-------", "----")
	for _, t := range cmds {
		fmt.Fprintf(tw, format, t.Path, t.Alias, t.Command, strings.Join(t.Args, " "))
	}
	tw.Flush()
}
