package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(rootCommand,
		cli.Tree(helpCommand),
		cli.Tree(startCommand),
		cli.Tree(stopCommand),
		cli.Tree(checkCommand),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
