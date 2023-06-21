package cli

import (
	"log"
	"sync"

	"github.com/fatih/color"
)

type Cli struct {
	Name        string
	Description string
	Usage       string
	Commands    []*Command
	Debug       bool
}

type Root struct {
	rootCommand *rootCommand
	Commands    map[string]*Command
	Debug       bool
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

var Debug bool = true

func New(v Cli) *Root {
	Debug = v.Debug

	if Debug {
		greenText := color.New(color.FgHiGreen).SprintFunc()
		infoLogger.Println("loading cli")
		defer infoLogger.Println(greenText("cli loaded and ready"))
	}

	rootCommand := &rootCommand{
		Name:        v.Name,
		Usage:       v.Usage,
		Description: v.Description,
	}

	if err := validateCommand(*rootCommand); err != nil {
		warnLogger.Println(err)
	}

	newRoot := &Root{
		rootCommand: rootCommand,
		Commands:    make(map[string]*Command),
	}
	newRoot.loadCommands(v.Commands)

	if err := newRoot.rootCommand.generateRootHelp(newRoot.Commands); err != nil {
		warnLogger.Println(err)
	}

	return newRoot
}

func (root *Root) loadCommands(newCommands []*Command) error {
	var wg sync.WaitGroup
	for _, v := range newCommands {
		wg.Add(1)
		current := v
		go func() {
			err := root.NewCommand(current)
			if err != nil {
				warnLogger.Println(err)
			}

			if Debug {
				infoLogger.Println("loading command: " + current.Name)

				for _, o := range current.Options {
					infoLogger.Println("loading option: " + o.Name + " for " + current.Name)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return nil
}
