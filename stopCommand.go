package main

import (
	"github.com/mkideal/cli"
)

type stop struct {
	cli.Helper
	Output string `cli:"o,output" usage:"Output directory for temporary files (nginx configuration, SSL certificates, ...)" dft:"/tmp/"`
	Docker string `cli:"d,docker"`
}

var stopCommand = &cli.Command{
	Name: "stop",
	Desc: "this is a stop command",
	Argv: func() interface{} { return new(start) },
	Fn: func(ctx *cli.Context) error {
		ctx.String("Hello, start command\n")
		return nil
	},
}
