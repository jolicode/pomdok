package main

import (
	"github.com/mkideal/cli"
)

type stop struct {
	cli.Helper
}

var stopCommand = &cli.Command{
	Name: "stop",
	Desc: "this is a stop command",
	Argv: func() interface{} { return new(stop) },
	Fn: func(ctx *cli.Context) error {
		ctx.String("Hello, start command\n")
		return nil
	},
}
