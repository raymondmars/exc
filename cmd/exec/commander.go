package main

import (
	"fmt"
	"os/exec"
)

type Commander struct {
	path    string
	alias   string
	command string
	args    []string
}

func (c *Commander) execute() {
	cmd := exec.Command(c.command, c.args...)
	out, _ := cmd.CombinedOutput()
	fmt.Printf(string(out))
}
