package main

import (
	"github.com/mkideal/cli"
)

type start struct {
	cli.Helper
	Output string `cli:"o,output" usage:"Output directory for temporary files (nginx configuration, SSL certificates, ...)" dft:"/tmp/"`
	Docker string `cli:"d,docker"`
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
