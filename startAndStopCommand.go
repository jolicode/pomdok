package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"

	"github.com/mkideal/cli"
)

type startT struct {
	cli.Helper
}

var startCommand = &cli.Command{
	Name: "start",
	Desc: "start all apps linked to your symfony binary",
	Argv: func() interface{} { return new(startT) },
	Fn: func(ctx *cli.Context) error {
		printHeader()
		startOrStopCommand("local:server:start --daemon", "started ✔")

		return nil
	},
}

type stopT struct {
	cli.Helper
}

var stopCommand = &cli.Command{
	Name: "stop",
	Desc: "stop all apps linked to your symfony binary",
	Argv: func() interface{} { return new(stopT) },
	Fn: func(ctx *cli.Context) error {
		printHeader()
		startOrStopCommand("local:server:stop", "stopped 🛑")

		return nil
	},
}

func startOrStopCommand(command string, message string) {
	if false == symfonyProxyRunning() {
		runCommand("/usr/local/bin/symfony proxy:start")
		fmt.Print("Started Symfony proxy server 👮\n")
	}

	user, _ := user.Current()
	file, _ := ioutil.ReadFile(fmt.Sprintf("%s/.symfony/proxy.json", user.HomeDir))

	symfonyJsonData := SymfonyJsonProxy{}
	json.Unmarshal(file, &symfonyJsonData)

	for domain, path := range symfonyJsonData.Domains {
		runCommand(fmt.Sprintf("/usr/local/bin/symfony %s --dir=%s", command, path))
		fmt.Printf("%s %s\n", message, yellow(fmt.Sprintf("%s.%s", domain, symfonyJsonData.Tld)))
	}

	return
}

func symfonyProxyRunning() bool {
	response, err := http.Get("http://127.0.0.1:7080/")

	if nil != err || 200 != response.StatusCode {
		return false
	}

	return true
}
