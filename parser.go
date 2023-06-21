package cli

import (
	"fmt"
	"os"
	"strings"
)

func (root *Root) Parser() {
	if Debug {
		infoLogger.Println("running parser")
		defer infoLogger.Println("parser finished")
	}

	arguments := os.Args

	if len(arguments) < 2 {
		root.rootCommand.Help.print()
		return
	}

	commandName := arguments[1]
	commandArgs := arguments[2:]
	if commandName == "-h" || commandName == "--help" {
		root.rootCommand.Help.print()
		return
	}

	command, err := root.findCommand(commandName)
	if err != nil {
		errorLogger.Fatalln(err)
	}

	for _, v := range commandArgs {
		if v == "-h" || v == "--help" {
			command.Help.print()
			return
		}
	}

	var sortedCommandArgs [][]string
	var currentGroup []string
	for _, v := range commandArgs {
		if strings.HasPrefix(v, "-") || strings.HasPrefix(v, "--") {
			if currentGroup != nil {
				sortedCommandArgs = append(sortedCommandArgs, currentGroup)
			}
			currentGroup = nil
		}
		currentGroup = append(currentGroup, v)
	}
	sortedCommandArgs = append(sortedCommandArgs, currentGroup)

	fmt.Println(checkArgumentsCommand(command, sortedCommandArgs[0]))

	if len(commandArgs) > 0 {
		for _, v := range sortedCommandArgs {
			if !strings.Contains(v[0], "-") || !strings.Contains(v[0], "--") {
				continue
			}

			if err := checkOption(v[0], command); err != nil {
				errorLogger.Fatalln(err)
			}
		}
		command.Run(sortedCommandArgs)
	} else {
		command.Run(nil)
	}

}

func (root *Root) findCommand(name string) (*Command, error) {
	availableCommands := root.Commands
	for k, v := range availableCommands {
		if k == name {
			return v, nil
		}
	}

	return nil, fmt.Errorf("unrecognized command '%s'", name)
}

func checkOption(name string, command *Command) error {
	formattedName := strings.TrimPrefix(name, "--")
	formattedName = strings.TrimPrefix(formattedName, "-")

	for _, v := range command.Options {
		if v.Name == formattedName || v.ShortName == formattedName {
			return nil
		}
	}

	return fmt.Errorf("unrecognized option '%s'", name)
}

func checkArgumentsCommand(command *Command, args []string) bool {
	return command.Arguments == len(args)
}

func checkArgumentsOption(name string, command *Command, args []string) bool {
	formattedName := strings.TrimPrefix(name, "-")
	formattedName = strings.TrimPrefix(formattedName, "--")

	for _, v := range command.Options {
		if (v.Name == formattedName || v.ShortName == formattedName) && v.Arguments == len(args) {
			return true
		}
	}

	return false
}
