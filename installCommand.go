package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
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
				fmt.Print("You need to have " + underline("linux") + " or " + underline("darwin") + " (MacOS) in order to be compatible with this binary.\n")
			}
		}

		return nil
	},
}

func checkIfRoot() bool {
	out := execCommand("whoami")
	if out == "root" {
		return true
	}

	fmt.Print("You need to be " + underline("root") + " or use " + underline("sudo") + " in order to use this command.\n")

	return false
}

func linuxInstall() {
	exists, _ := checkBinaryExists("php")
	if exists == false {
		fmt.Print("Starting " + color.YellowString("php") + " installation\n")
		execCommand("apt install php -y")

		exists, _ = checkBinaryExists("php")
		if exists == false {
			fmt.Printf("%s installation error ...\n", color.YellowString("php"))
			os.Exit(1)
		}

		fmt.Printf("%s installed\n", color.YellowString("php"))
		fmt.Printf("With this command we only installed %s binary but no extensions,\n", color.YellowString("php"))
		fmt.Printf("if you do need extensions you'll have to install them by yourself.\n")
		fmt.Print("\n")
	}

	exists, _ = checkBinaryExists("nginx")
	if exists == false {
		fmt.Printf("Starting %s installation.\n", color.YellowString("nginx"))
		execCommand("apt install nginx -y")

		exists, _ = checkBinaryExists("nginx")
		if exists == false {
			fmt.Printf("%s installation error ...\n", color.YellowString("nginx"))
			os.Exit(1)
		}

		fmt.Printf("%s installed in path: /etc/nginx\n", color.YellowString("nginx"))
		execCommand("rm /etc/nginx/sites-enabled/*")
		fmt.Print("Removed default enabled configuration to not bind port 80\n")
		execCommand("/etc/init.d/nginx restart")
		fmt.Printf("Restarted %s to update configuration\n", color.YellowString("nginx"))
		fmt.Print("\n")
	}

	exists, _ = checkBinaryExists("symfony")
	if exists == false {
		fmt.Print("Starting " + underline("symfony") + " installation\n")
		execCommand("wget https://get.symfony.com/cli/installer -O - | bash")
		execCommand("mv $HOME/.symfony/bin/symfony /usr/local/bin/")

		exists, _ = checkBinaryExists("symfony")
		if exists == false {
			fmt.Printf("%s installation error ...\n", color.YellowString("symfony"))
			os.Exit(1)
		}

		fmt.Printf("%s installed\n", color.YellowString("symfony"))
		fmt.Print("\n")
	}

	return
}

func darwinInstall() {
	return
}
