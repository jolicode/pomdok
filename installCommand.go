package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/mkideal/cli"
)

type install struct {
	cli.Helper
}

var installCommand = &cli.Command{
	Name: "install",
	Desc: "Install needed binaries, run check command before to be sure you need it",
	Argv: func() interface{} { return new(start) },
	Fn: func(ctx *cli.Context) error {
		printHeader()

		if checkIfRoot() {
			if runtime.GOOS == "linux" {
				linuxInstall()
			} else if runtime.GOOS == "darwin" {
				darwinInstall()
			} else {
				fmt.Printf("You need to have %s or %s (MacOS) in order to be compatible with this binary.\n", underline("linux"), underline("darwin"))
			}

			fmt.Printf("Run %s command to verify if everything was installed.\n", yellow("pomdok check"))
		}

		return nil
	},
}

func checkIfRoot() bool {
	out := execCommand("whoami")
	if out == "root" {
		return true
	}

	fmt.Printf("You need to be %s or use %s in order to use this command.\n", underline("root"), underline("sudo"))

	return false
}

func linuxInstall() {
	phpInstall("apt install php -y")
	symfonyCliInstall()

	return
}

func darwinInstall() {
	phpInstall("brew install php72")
	symfonyCliInstall()

	return
}

func phpInstall(command string) {
	exists, _ := checkBinaryExists("php")
	if exists == false {
		fmt.Printf("Starting %s installation 🏃\n", yellow("php"))
		execCommand(command)

		exists, _ = checkBinaryExists("php")
		if exists == false {
			fmt.Printf("%s installation error ... 😭\n", yellow("php"))
			os.Exit(1)
		}

		fmt.Printf("%s installed ✔\n", yellow("php"))
		fmt.Printf("With this command we only installed %s binary but no extensions,\n", yellow("php"))
		fmt.Printf("if you do need extensions you'll have to install them by yourself.\n")
		fmt.Print("\n")
	}
}

func symfonyCliInstall() {
	exists, _ := checkBinaryExists("symfony")
	if exists == false {
		fmt.Printf("Starting %s installation 🏃\n", yellow("symfony"))
		execCommand("wget https://get.symfony.com/cli/installer -O - | bash")
		execCommand("mv $HOME/.symfony/bin/symfony /usr/local/bin/")

		exists, _ = checkBinaryExists("symfony")
		if exists == false {
			fmt.Printf("%s installation error ... 😭\n", yellow("symfony"))
			os.Exit(1)
		}

		fmt.Printf("%s installed ✔\n", yellow("symfony"))
		fmt.Print("\n")
	}
}
