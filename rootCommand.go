package main

import (
	"fmt"

	"github.com/mkideal/cli"
)

type appT struct {
	Name    string
	Version string
}

var app = appT{
	"pomdok",
	"v1.3.1",
}

func sprintHeader() string {
	return fmt.Sprintf("%s version %s", green(app.Name), yellow(app.Version))
}

func printHeader() {
	fmt.Print(sprintHeader() + "\n\n")
}

type rootT struct {
	cli.Helper
}

var rootCommand = &cli.Command{
	Desc: sprintHeader(),
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		printHeader()

		fmt.Print("Usage: pomdok <command>\n")
		fmt.Printf("More information on usage with %s command.\n", yellow("help"))

		return nil
	},
}
