package main

import (
	"github.com/mkideal/cli"
)

type check struct {
	cli.Helper
}

var checkCommand = &cli.Command{
	Name: "check",
	Desc: "this is a check command",
	Argv: func() interface{} { return new(start) },
	Fn: func(ctx *cli.Context) error {
		ctx.String("Hello, start command\n")
		return nil
	},
}
