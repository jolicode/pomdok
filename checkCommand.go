package main

import (
	"os"
	"runtime"

	"github.com/mkideal/cli"
	"github.com/olekukonko/tablewriter"
)

type check struct {
	cli.Helper
}

var checkCommand = &cli.Command{
	Name: "check",
	Desc: "this is a check command",
	Argv: func() interface{} { return new(start) },
	Fn: func(ctx *cli.Context) error {
		data := [][]string{
			[]string{"System", runtime.GOOS, "Operating System you run"},
			[]string{"B", "The Very very Bad Man", "288"},
			[]string{"C", "The Ugly", "120"},
			[]string{"D", "The Gopher", "800"},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Value", "Description"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render()

		return nil
	},
}
