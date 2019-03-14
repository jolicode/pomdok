package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mkideal/cli"
)

type appT struct {
	Name    string
	Version string
}

var app = appT{
	"pomdok",
	"0.1.0-beta",
}

func sprintHeader() string {
	return fmt.Sprintf("%s version %s", color.GreenString(app.Name), color.YellowString(app.Version))
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
		ctx.String("Hello, root command\n")
		return nil
	},
}
