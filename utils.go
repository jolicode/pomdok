package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

var underline = color.New(color.Underline).SprintFunc()
var bold = color.New(color.Bold).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()

func checkBinaryExists(command string) (bool, string) {
	path, err := exec.LookPath(command)
	return err == nil, path
}

func outputCommand(command string) string {
	out, _ := exec.Command("sh", "-c", command).Output()

	return strings.TrimSuffix(string(out), "\n")
}

func runCommand(command string) {
	exec.Command("sh", "-c", command).Run()

	return
}

func getSymfonyCliConfigPath() string {
	currentUser, _ := user.Current()
	pathForSymfonyCli5 := fmt.Sprintf("%s/.symfony5", currentUser.HomeDir)
	pathForSymfonyCliOld := fmt.Sprintf("%s/.symfony", currentUser.HomeDir)
	path := ""

	if _, err := os.Stat(pathForSymfonyCli5); err == nil {
		path = pathForSymfonyCli5
	} else if errors.Is(err, os.ErrNotExist) {
		if _, err := os.Stat(pathForSymfonyCliOld); err == nil {
			path = pathForSymfonyCliOld
		} else if errors.Is(err, os.ErrNotExist) {
			path = ""
		} else {
			exitWithMessage(fmt.Sprintf("Unexpected error while checking for the existence of %s. Error: %v\n", pathForSymfonyCliOld, err))
		}
	} else {
		exitWithMessage(fmt.Sprintf("Unexpected error while checking for the existence of %s. Error: %v\n", pathForSymfonyCli5, err))
	}

	if path != "" {
		info, _ := os.Stat(path)
		symfonyDirUserUID := fmt.Sprint((info.Sys().(*syscall.Stat_t)).Uid)
		symfonyDirUser, _ := user.LookupId(symfonyDirUserUID)
		if symfonyDirUser.Username != currentUser.Username {
			exitWithMessage(fmt.Sprintf("Permission error 🙊. Directory %s is owned by %s, please use: 'sudo chown -R %s %s' 🧐\n", path, yellow(symfonyDirUser.Username), currentUser.Username, path))
		}
	}

	return path
}

func exitWithMessage(message string) {
	fmt.Fprint(os.Stderr, message)
	os.Exit(1)
}
