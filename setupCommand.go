package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mkideal/cli"
	"gopkg.in/yaml.v2"
)

type setup struct {
	cli.Helper
	Config string `cli:"config" usage:"Your configuration file" dft:"pomdok.yaml"`
}

/**
 * Configuration should be as following:
 * ```
 * pomdok:
 *   tld: 'test'
 *   projects:
 *     - domain: 'api.project'
 *       path: '/project/api'
 *     - domain: 'www.project'
 *       path: '/project/www'
 * ```
 **/

type Config struct {
	Pomdok struct {
		Tld      string
		Projects []struct {
			Domain string `yaml:"domain"`
			Path   string `yaml:"path"`
		}
	}
}

var setupCommand = &cli.Command{
	Name: "setup",
	Desc: "Setup your local symfony binary environment to work with a given project",
	Argv: func() interface{} { return new(setup) },
	Fn: func(ctx *cli.Context) error {
		printHeader()

		argv := ctx.Argv().(*setup)
		config := Config{}

		data, _ := ioutil.ReadFile(argv.Config)
		yaml.Unmarshal([]byte(data), &config)
		fmt.Printf("--- t:\n%v\n\n", config.Pomdok.Tld)

		return nil
	},
}
