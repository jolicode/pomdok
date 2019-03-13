package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
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
		title := color.New(color.Bold).SprintFunc()
		data := [][]string{
			[]string{title("System"), sprintCheckSystem(), "Operating System you run"},
			[]string{title("PHP"), sprintCheckCliExists("php"), "PHP runtime"},
			[]string{title("nginx"), sprintCheckCliExists("nginx"), "Proxy server"},
			[]string{title("Symfony"), sprintCheckCliExists("symfony"), "Symfony CLI"},
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

func sprintCheckSystem() string {
	system := runtime.GOOS
	if system == "linux" || system == "darwin" {
		return color.GreenString(system)
	} else {
		return color.RedString(system)
	}
}

func sprintCheckCliExists(command string) string {
	out := checkCliExists(command)
	if out == "" {
		out = color.RedString("Not-found")
	} else {
		out = color.GreenString(out)
	}

	return out
}

func checkCliExists(command string) string {
	return execCommand(fmt.Sprintf("which %s", command))
}

func execCommand(command string) string {
	out, _ := exec.Command("sh", "-c", command).Output()

	return strings.TrimSuffix(string(out), "\n")
}
