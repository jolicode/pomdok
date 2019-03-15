package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var underline = color.New(color.Underline).SprintFunc()

func checkBinaryExists(command string) (bool, string) {
	out := execCommand(fmt.Sprintf("which %s", command))
	return out != "", out
}

func execCommand(command string) string {
	out, _ := exec.Command("sh", "-c", command).Output()

	return strings.TrimSuffix(string(out), "\n")
}
