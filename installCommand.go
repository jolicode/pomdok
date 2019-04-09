package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	"github.com/mkideal/cli"
)

type install struct {
	cli.Helper
}

var installCommand = &cli.Command{
	Name: "install",
	Desc: "install needed binaries, run check command before to be sure you need it",
	Argv: func() interface{} { return new(install) },
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
	user, _ := user.Current()
	if "root" != user.Username {
		fmt.Printf("You need to use %s in order to use this command.\n", underline("sudo"))
		return false
	}

	if "root" == user.Username && "" == os.Getenv("SUDO_USER") {
		fmt.Printf("Please do not run this command as %s but with %s.\n", underline("root"), underline("sudo"))
		return false
	}

	return true
}

func linuxInstall() {
	phpInstall("apt install php -y")
	symfonyCliInstall()

	return
}

func darwinInstall() {
	phpInstall("brew install php72 nss")
	symfonyCliInstall()

	return
}

func phpInstall(command string) {
	exists, _ := checkBinaryExists("php")
	if exists == false {
		fmt.Printf("Starting %s installation 🏃\n", yellow("php"))
		runCommand(command)

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
		runCommand("wget https://get.symfony.com/cli/installer -O - | bash")
		runCommand("mv $HOME/.symfony/bin/symfony /usr/local/bin/")

		exists, _ = checkBinaryExists("symfony")
		if exists == false {
			fmt.Printf("%s installation error ... 😭\n", yellow("symfony"))
			os.Exit(1)
		}

		currentUser, _ := user.Current()
		username := currentUser.Username
		if runtime.GOOS == "darwin" {
			runCommand(fmt.Sprintf("sudo chown -R %s:staff ~/.symfony", username))
		} else {
			runCommand(fmt.Sprintf("sudo chown -R %s:%s ~/.symfony", username, username))
		}

		fmt.Printf("%s installed ✔\n", yellow("symfony"))
		fmt.Print("\n")
	}
}
