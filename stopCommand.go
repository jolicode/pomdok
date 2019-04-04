package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/mkideal/cli"
)

type stop struct {
	cli.Helper
}

var stopCommand = &cli.Command{
	Name: "stop",
	Desc: "Stop all apps linked to your symfony binary",
	Argv: func() interface{} { return new(stop) },
	Fn: func(ctx *cli.Context) error {
		printHeader()

		user, _ := user.Current()
		file, _ := ioutil.ReadFile(fmt.Sprintf("%s/.symfony/proxy.json", user.HomeDir))

		symfonyJsonData := SymfonyJsonProxy{}
		json.Unmarshal(file, &symfonyJsonData)

		for domain, path := range symfonyJsonData.Domains {
			runCommand(fmt.Sprintf("/usr/local/bin/symfony local:server:stop --dir=%s", path))
			fmt.Printf("%s stopped 🛑\n", yellow(fmt.Sprintf("%s.%s", domain, symfonyJsonData.Tld)))
		}

		return nil
	},
}
