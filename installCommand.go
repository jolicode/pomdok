package main

import (
	"github.com/mkideal/cli"
)

type install struct {
	cli.Helper
}

var installCommand = &cli.Command{
	Name: "install",
	Desc: "install command",
	Argv: func() interface{} { return new(start) },
	Fn: func(ctx *cli.Context) error {
		ctx.String("Hello, start command\n")
		return nil
	},
}
