package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
		runCommand("symfony proxy:stop")

		return nil
	},
}

func startOrStopCommand(command string, message string) {
	if false == symfonyProxyRunning() {
		runCommand("symfony proxy:start")
		fmt.Print("Started Symfony proxy server 👮\n")
	}

	user, _ := user.Current()
	symfonyProxyConfigPah := fmt.Sprintf("%s/.symfony/proxy.json", user.HomeDir)
	if _, err := os.Stat(symfonyProxyConfigPah); os.IsNotExist(err) {
		fmt.Printf("Symfony proxy configuration does not exists 🙊. Maybe you should run %s before %s ? 🧐\n", yellow("init"), yellow("start"))
		return
	}
	file, _ := ioutil.ReadFile(symfonyProxyConfigPah)

	symfonyJSONData := SymfonyJSONProxy{}
	json.Unmarshal(file, &symfonyJSONData)

	for domain, path := range symfonyJSONData.Domains {
		forcedPort := symfonyJSONData.Ports[domain]
		formattedCommand := fmt.Sprintf("symfony %s --dir=%s", command, path)

		if "local:server:start --daemon" == command && 0 != forcedPort {
			formattedCommand = fmt.Sprintf("symfony %s --port=%d --dir=%s", command, forcedPort, path)
		}

		runCommand(formattedCommand)
		fmt.Printf("%s %s\n", message, yellow(fmt.Sprintf("%s.%s", domain, symfonyJSONData.Tld)))
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
