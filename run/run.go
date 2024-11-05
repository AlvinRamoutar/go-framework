package run

import (
	"context"
	"os/exec"
	"strings"
)

type Run struct {
	Config *RunConfig
	Cmds   []Cmd
}

type Cmd struct {
	Name string
	Command string
	Args []string
	Context context.Context
	Stdout strings.Builder
	Stderr strings.Builder
	Status string
}

func (c *Cmd) Run() {
	if c.Context == nil {
		runnableCmd := exec.Command(c.Command, c.Args...)
		var asdf strings.Builder
		runnableCmd = &asdf
	} else {
		
	}
}