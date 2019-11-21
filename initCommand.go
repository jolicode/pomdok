package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"strings"
	"syscall"

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
		config, baseDirectory, _ := loadPomdokConfig(argv.Config)
		if config.Pomdok.Tld == "" {
			fmt.Printf("Configuration file error 🙊. Maybe you should give a %s to your domains 🧐\n", yellow("tld"))
			return nil
		}
		if config.Pomdok.Projects == nil {
			fmt.Printf("Configuration file error 🙊. Maybe you should add %s 🧐\n", yellow("projects"))
			return nil
		}

		fileDomains := make(map[string]string)
		filePorts := make(map[string]int)
		for _, element := range config.Pomdok.Projects {
			if element.Domain == "" {
				fmt.Printf("Configuration file error 🙊. One of the project has empty/no %s 🧐\n", yellow("domain"))
				return nil
			}
			if element.Path == "" {
				fmt.Printf("Configuration file error 🙊. One of the project has empty/no %s 🧐\n", yellow("path"))
				return nil
			}

			fullPath := baseDirectory + element.Path
			if _, err := os.Stat(fullPath); os.IsNotExist(err) {
				fmt.Printf("Configuration file error 🙊. %s path is not found 🧐\n", bold(fullPath))
				return nil
			}

			if _, ok := fileDomains[element.Domain]; ok {
				fmt.Printf("Configuration file error 🙊. Domain %s is used more than one time 🧐\n", yellow(element.Domain))
				return nil
			}

			fileDomains[element.Domain] = fullPath
			filePorts[element.Domain] = element.Port
		}

		symfonyJSONData := SymfonyJSONProxy{
			Tld:     config.Pomdok.Tld,
			Port:    7080,
			Domains: fileDomains,
			Ports:   filePorts,
		}
		symfonyJSON, _ := json.MarshalIndent(symfonyJSONData, "", "  ")

		currentUser, _ := user.Current()

		info, err := os.Stat(fmt.Sprintf("%s/.symfony", currentUser.HomeDir))
		if os.IsNotExist(err) {
			fmt.Printf("Symfony Binary not installed 🙊. Please use %s to see what you should do 🧐\n", yellow("symfony check"))
			return nil
		}

		symfonyDirUserUID := fmt.Sprint((info.Sys().(*syscall.Stat_t)).Uid)
		symfonyDirUser, _ := user.LookupId(symfonyDirUserUID)
		if symfonyDirUser.Username != currentUser.Username {
			fmt.Printf("Permission error 🙊. Directory ~/.symfony is owned by %s, please use: 'sudo chown -R %s ~/.symfony' 🧐\n", yellow(symfonyDirUser.Username), currentUser.Username)
			return nil
		}

		ioutil.WriteFile(fmt.Sprintf("%s/.symfony/proxy.json", currentUser.HomeDir), symfonyJSON, 0644)
		fmt.Printf("Project setup done ✔\n")

		return nil
	},
}

func loadPomdokConfig(fileName string) (PomdokYamlConfig, string, error) {
	config := PomdokYamlConfig{}

	file := findFileUp(fileName, 0)
	if file == "" {
		return config, "", errors.New("No file found")
	}

	data, _ := ioutil.ReadFile(file)
	yaml.Unmarshal([]byte(data), &config)

	return config, path.Dir(file), nil
}

func findFileUp(file string, level int) string {
	temp := file
	if level > 0 {
		temp = strings.Repeat("../", level) + file
	}

	currentDirectory, _ := os.Getwd()
	temp = path.Clean(currentDirectory + "/" + temp)

	if temp == "/" {
		fmt.Print("Configuration file does not exists 🙊. Maybe you should create or rename your configuration file ? 🧐\n")
		return ""
	}

	if _, err := os.Stat(temp); os.IsNotExist(err) {
		return findFileUp(file, level+1) // not found, go up
	}

	return temp
}
