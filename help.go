package cli

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/fatih/color"
)

type Help struct {
	title    string
	usage    string
	commands string
	options  string
}

func (h *Help) print() {
	boldText := color.New(color.FgHiWhite, color.Bold).SprintFunc()

	fmt.Println(h.title + "\n")
	fmt.Println(boldText("Usage: ") + h.usage + "\n")
	if len(h.commands) > 0 {
		fmt.Println(boldText("Commands: \n") + h.commands)
	}

	if len(h.options) > 0 {
		fmt.Println(boldText("Options: \n") + h.options)
	}

}

func (cmd *Command) generateHelp() error {
	if Debug {
		infoLogger.Println("generating help for " + cmd.Name)
		defer infoLogger.Println("generating finished")
	}

	optionsBuf := new(bytes.Buffer)
	writer := tabwriter.NewWriter(optionsBuf, 1, 1, 5, ' ', 0)
	for idx, v := range cmd.Options {
		if idx != len(cmd.Options)-1 {
			fmt.Fprintf(writer, "    -%s, --%s\t%s\n", v.ShortName, v.Name, v.Description)
		} else {
			fmt.Fprintf(writer, "    -%s, --%s\t%s", v.ShortName, v.Name, v.Description)
		}
	}
	writer.Flush()

	cmd.Help = Help{
		title:   cmd.Description,
		usage:   cmd.Usage,
		options: optionsBuf.String(),
	}

	return nil
}

func (cmd *rootCommand) generateRootHelp(commands map[string]*Command) error {
	if Debug {
		infoLogger.Println("generating help for root")
		defer infoLogger.Println("generating finished")
	}

	buf := new(bytes.Buffer)
	writer := tabwriter.NewWriter(buf, 1, 1, 5, ' ', 0)

	idx := 0
	for _, v := range commands {
		if idx != len(commands)-1 {
			fmt.Fprintf(writer, "%s\t%s\n", v.Name, v.Description)
		} else {
			fmt.Fprintf(writer, "%s\t%s", v.Name, v.Description)
		}
		idx++
	}

	writer.Flush()

	cmd.Help = Help{
		title:    cmd.Description,
		usage:    cmd.Usage,
		commands: buf.String(),
	}

	return nil
}
