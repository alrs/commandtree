package main

import (
	"commandtree"
	"fmt"
	"log"
)

func main() {
	commands := commandtree.NewRoot()

	commands["other"] = commandtree.NewCommand(
		"Does something else",
		"these long descriptions are super long\nand really tedious\nand go on.",
		func() error {
			fmt.Println("other")
			return nil
		})

	commands["other"].Subcommands["sub"] = commandtree.NewCommand(
		"a subcommand",
		"this is an esoteric subcommand",
		func() error {
			fmt.Println("sub")
			return nil
		})

	commands["other"].Subcommands["constructasub"] = commandtree.NewCommand(
		"created via a constructor function",
		"such a super long help message.\nYes.\n",
		func() error { return nil })

	commands["other"].Subcommands["sub"].Subcommands["waydeep"] = commandtree.NewCommand(
		"a deep subcommand",
		"there is nothing less likely than a subcommand this deep",
		func() error {
			fmt.Println("waydeep")
			return nil
		})

	err := commands.Do()
	if err != nil {
		log.Fatal(err)
	}
}
