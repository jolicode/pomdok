package main

import (
	"github.com/mkideal/cli"
)

type setup struct {
	cli.Helper
}

/**
 * Configuration should be as following:
 * ```
 * pomdok:
 *   tld: 'test'
 *   projects:
 *     - host: 'api.project'
 *       path: '/project/api'
 *     - host: 'www.project'
 *       path: '/project/www'
 * ```
 **/

var setupCommand = &cli.Command{
	Name: "setup",
	Desc: "Setup your local symfony binary environment to work with a given project",
	Argv: func() interface{} { return new(start) },
	Fn: func(ctx *cli.Context) error {
		printHeader()

		return nil
	},
}
