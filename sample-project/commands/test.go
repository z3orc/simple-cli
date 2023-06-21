package commands

import (
	"fmt"

	cli "github.com/z3orc/simple-cli"
)

var Test *cli.Command = &cli.Command{Name: "run",
	Usage:       "run",
	Description: "runs the test program",
	Arguments:   0,
	Options:     []*cli.Option{testOption},
	Run: func(args [][]string) {
		if len(args) > 0 {
			fmt.Println("used an option")
		}
		fmt.Println("Running test")
		fmt.Println("Test finished")
	}}

var testOption *cli.Option = &cli.Option{
	Name:        "opt",
	ShortName:   "o",
	Description: "a demo option",
	Arguments:   1,
}
