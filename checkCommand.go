package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/mkideal/cli"
	"github.com/olekukonko/tablewriter"
)

type check struct {
	cli.Helper
}

var checkCommand = &cli.Command{
	Name: "check",
	Desc: "Check compatibility and needed binaries",
	Argv: func() interface{} { return new(start) },
	Fn: func(ctx *cli.Context) error {
		title := color.New(color.Bold).SprintFunc()

		systemOk, systemString := sprintCheckSystem()
		phpOk, phpString := sprintCheckCliExists("php")
		nginxOk, nginxString := sprintCheckCliExists("nginx")
		symfonyOk, symfonyString := sprintCheckCliExists("symfony")

		data := [][]string{
			[]string{title("System"), systemString, "Operating System you run"},
			[]string{title("PHP"), phpString, "PHP runtime"},
			[]string{title("nginx"), nginxString, "Proxy server"},
			[]string{title("Symfony"), symfonyString, "Symfony CLI"},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Name", "Value", "Description"})
		for _, v := range data {
			table.Append(v)
		}

		printHeader()
		fmt.Printf("Check needed binaries, %s means everything is ok, \n", color.GreenString("green status"))
		fmt.Printf("%s means you have to install corresponding binary.\n", color.RedString("red status"))
		table.Render()

		if systemOk == false || phpOk == false || nginxOk == false || symfonyOk == false {
			fmt.Printf("We saw atleast one missing binary, you can run %s to fix it.\n", color.YellowString("pomdok install"))
		} else {
			fmt.Printf("Everything is fine, you can start using %s 🎉.\n", color.YellowString("pomdok"))
		}

		return nil
	},
}

func sprintCheckSystem() (bool, string) {
	system := runtime.GOOS

	if system == "linux" || system == "darwin" {
		return true, color.GreenString(system)
	}
	return false, color.RedString(system)
}

func sprintCheckCliExists(command string) (bool, string) {
	exists, out := checkBinaryExists(command)
	if exists {
		out = color.GreenString(out)
	} else {
		out = color.RedString("Not-found")
	}

	return exists, out
}
