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

func printHeader() {
	fmt.Printf("%s version %s\n\n", color.GreenString(app.Name), color.YellowString(app.Version))
}

type rootT struct {
	cli.Helper
}

var rootCommand = &cli.Command{
	Desc: "this is root command",
	// Argv is a factory function of argument object
	// ctx.Argv() is if Command.Argv == nil or Command.Argv() is nil
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		ctx.String("Hello, root command\n")
		return nil
	},
}
