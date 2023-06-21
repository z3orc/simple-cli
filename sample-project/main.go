package main

import (
	cli "github.com/z3orc/simple-cli"
	"github.com/z3orc/simple-cli/sample-project/commands"
)

func main() {
	cli.New(cli.Cli{
		Name:        "test",
		Description: "this is a sample project",
		Usage:       "test <command>",
		Commands:    []*cli.Command{commands.Test},
	}).Parser()
}
