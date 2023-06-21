package cli

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

var lock sync.Mutex

type rootCommand struct {
	Name        string
	Usage       string
	Description string
	Options     []*Option
	Help        Help
	Debug       bool
}

type Command struct {
	Name        string
	Usage       string
	Description string
	Arguments   int
	Help        Help
	Options     []*Option
	Run         func([][]string)
}

type Option struct {
	Name        string
	ShortName   string
	Description string
	Arguments   int
}

func (root *Root) NewCommand(newCommand *Command) error {
	errorMsg := func(err error) error {
		return fmt.Errorf("cannot create command: '%s'. %s", newCommand.Name, err)
	}

	if err := validateCommand(*newCommand); err != nil {
		return errorMsg(err)
	}

	eg, _ := errgroup.WithContext(context.Background())
	for _, v := range newCommand.Options {
		current := v
		eg.Go(func() error {
			if err := validateOption(*current); err != nil {
				return err
			}
			return nil
		})

	}
	if err := eg.Wait(); err != nil {
		return errorMsg(err)
	}

	newCommand.Usage = fmt.Sprintf("%s %s", root.rootCommand.Name, newCommand.Usage)
	if err := newCommand.generateHelp(); err != nil {
		return errorMsg(err)
	}

	lock.Lock()
	defer lock.Unlock()
	root.Commands[newCommand.Name] = newCommand
	return nil
}
