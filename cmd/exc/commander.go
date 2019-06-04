package main

import (
	"fmt"
	"os/exec"
)

type Commander struct {
	Path    string
	Alias   string
	Command string
	Args    []string
}

func (c *Commander) execute() {
	cmd := exec.Command(c.Command, c.Args...)
	out, _ := cmd.CombinedOutput()
	fmt.Printf(string(out))
}
