package main

import (
	"github.com/mkideal/cli"
)

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
