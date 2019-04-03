package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(rootCommand,
		cli.Tree(cli.HelpCommand("display help information")),
		cli.Tree(startCommand),
		cli.Tree(stopCommand),
		cli.Tree(setupCommand),
		cli.Tree(checkCommand),
		cli.Tree(installCommand),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
