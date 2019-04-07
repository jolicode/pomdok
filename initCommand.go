package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"

	"github.com/mkideal/cli"
	"gopkg.in/yaml.v2"
)

type initT struct {
	cli.Helper
	Config string `cli:"config" usage:"Your configuration file" dft:"pomdok.yaml"`
}

var initCommand = &cli.Command{
	Name: "init",
	Desc: "init your local symfony binary environment to work with a given project",
	Argv: func() interface{} { return new(initT) },
	Fn: func(ctx *cli.Context) error {
		printHeader()

		argv := ctx.Argv().(*initT)
		config := PomdokYamlConfig{}

		data, _ := ioutil.ReadFile(argv.Config)
		yaml.Unmarshal([]byte(data), &config)

		fileDomains := make(map[string]string)
		currentDirectory, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		for _, element := range config.Pomdok.Projects {
			fileDomains[element.Domain] = currentDirectory + element.Path
		}

		symfonyJsonData := SymfonyJsonProxy{
			Tld:     config.Pomdok.Tld,
			Port:    7080,
			Domains: fileDomains,
		}
		symfonyJson, _ := json.MarshalIndent(symfonyJsonData, "", "  ")

		user, _ := user.Current()
		ioutil.WriteFile(fmt.Sprintf("%s/.symfony/proxy.json", user.HomeDir), symfonyJson, 0644)
		fmt.Printf("Project setup done ✔\n")

		runCommand("/usr/local/bin/symfony local:server:ca:install")
		fmt.Printf("Local certificate authority installed 🔐\n")

		return nil
	},
}
