package commandtree

import (
	"bytes"
	"fmt"
	"os"
)

type Commands map[string]Command

type Command struct {
	DescriptionText string
	HelpText        string
	Subcommands     Commands
	exec            func() error
}

func NewCommand(desc string, help string, exec func() error) Command {
	return Command{
		DescriptionText: desc,
		HelpText:        help,
		Subcommands:     make(Commands),
		exec:            exec,
	}
}

func NewRoot() Commands {
	commands := Commands{}
	commands["help"] = NewCommand(
		"'help' followed by a subcommand (and any of its subcommands) prints help.",
		"",
		func() error {
			if len(os.Args) < 3 {
				fmt.Println(commands.HelpSubcommandText())
				os.Exit(0)
			} else {
				helpPtr, ok := commands[os.Args[2]]
				if !ok {
					return fmt.Errorf("subcommand %s not found", os.Args[2])
				}
				for _, h := range os.Args[3:] {
					if helpPtr.Subcommands != nil {
						helpPtr = helpPtr.Subcommands[h]
					}
				}
				fmt.Println(helpPtr.DescriptionText + "\n")
				fmt.Println(helpPtr.HelpText + "\n")
				if len(helpPtr.Subcommands) > 0 {
					fmt.Println(helpPtr.Subcommands.HelpSubcommandText())
				}
			}
			return nil
		})
	return commands
}

func (c Commands) HelpSubcommandText() string {
	b := bytes.Buffer{}
	b.WriteString("subcommands: \n\n")
	for k, v := range c {
		b.WriteString(fmt.Sprintf("%15s: %s\n", k, v.DescriptionText))
	}
	return b.String()
}

func (cmds Commands) Do() error {
	switch len(os.Args) {
	case 1:
		return fmt.Errorf("no commands selected")
	case 2:
		sub := os.Args[1]
		_, ok := cmds[sub]
		if !ok {
			return fmt.Errorf("%s is an invalid command", sub)
		}
		return cmds[sub].Execute()
	}

	var err error
	if os.Args[1] == "help" {
		err = cmds["help"].Execute()
	} else {
		cmdPtr := cmds[os.Args[1]]
		for _, c := range os.Args[2:] {
			if cmdPtr.Subcommands != nil {
				cmdPtr = cmdPtr.Subcommands[c]
			}
		}
		err = cmdPtr.Execute()
	}
	return err
}

func (c Command) Execute() error {
	return c.exec()
}
