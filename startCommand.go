package main

import (
	"github.com/mkideal/cli"
)

type start struct {
	cli.Helper
}

var startCommand = &cli.Command{
	Name: "start",
	Desc: "this is a child command",
	Argv: func() interface{} { return new(start) },
	Fn: func(ctx *cli.Context) error {
		ctx.String("Hello, start command\n")
		return nil
	},
}
