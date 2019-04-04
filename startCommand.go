package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/mkideal/cli"
)

type start struct {
	cli.Helper
}

var startCommand = &cli.Command{
	Name: "start",
	Desc: "start all apps linked to your symfony binary",
	Argv: func() interface{} { return new(start) },
	Fn: func(ctx *cli.Context) error {
		printHeader()

		user, _ := user.Current()
		file, _ := ioutil.ReadFile(fmt.Sprintf("%s/.symfony/proxy.json", user.HomeDir))

		symfonyJsonData := SymfonyJsonProxy{}
		json.Unmarshal(file, &symfonyJsonData)

		for domain, path := range symfonyJsonData.Domains {
			runCommand(fmt.Sprintf("/usr/local/bin/symfony local:server:start --dir=%s --daemon", path))
			fmt.Printf("%s started ✔\n", yellow(fmt.Sprintf("%s.%s", domain, symfonyJsonData.Tld)))
		}

		return nil
	},
}
