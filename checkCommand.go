package main

import (
	"fmt"
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
	Desc: "check compatibility and needed binaries",
	Argv: func() interface{} { return new(check) },
	Fn: func(ctx *cli.Context) error {
		systemOk, systemString := sprintCheckSystem()
		phpOk, phpString := sprintCheckCliExists("php")
		symfonyOk, symfonyString := sprintCheckCliExists("symfony")

		data := [][]string{
			[]string{bold("System"), systemString, "Operating System you run"},
			[]string{bold("PHP"), phpString, "PHP runtime"},
			[]string{bold("Symfony"), symfonyString, "Symfony CLI"},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Value", "Description"})
		for _, v := range data {
			table.Append(v)
		}

		printHeader()
		fmt.Printf("Check needed binaries, %s means everything is ok, \n", green("green status"))
		fmt.Printf("%s means you have to install corresponding binary.\n", red("red status"))
		table.Render()

		if systemOk == false || phpOk == false || symfonyOk == false {
			fmt.Printf("We saw atleast one missing binary, you can run %s to fix it.\n", yellow("pomdok install"))
		} else {
			fmt.Printf("Everything is fine, you can start using %s 🎉.\n", yellow("pomdok"))
		}

		return nil
	},
}

func sprintCheckSystem() (bool, string) {
	system := runtime.GOOS

	if system == "linux" || system == "darwin" {
		return true, green(system)
	}
	return false, red(system)
}

func sprintCheckCliExists(command string) (bool, string) {
	exists, out := checkBinaryExists(command)
	if exists {
		out = green(out)
	} else {
		out = red("Not-found")
	}

	return exists, out
}
