package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

const shortListTemplate = `alias			command			args			
------------------------------------------------------{{range .}}
{{.Alias}}			{{.Command}}			{{StringsJoin .Args " "}}{{end}}
`

const detailListTemplate = `path			alias			command			args			
--------------------------------------------------------------------------------{{range .}}
{{.Path}}			{{.Alias}}			{{.Command}}			{{StringsJoin .Args " "}}{{end}}
`

func showCommands(cmds []Commander, templateStr string) {
	var tml = template.Must(template.New("commandlist").Funcs(template.FuncMap{"StringsJoin": strings.Join}).Parse(templateStr))

	if err := tml.Execute(os.Stdout, cmds); err != nil {
		log.Fatal(err)
	}
}
func showCurrentPathCommands(cmds []Commander) {
	showCommands(cmds, shortListTemplate)
}

func showAllPathCommands(cmds []Commander) {
	showCommands(cmds, detailListTemplate)
}
