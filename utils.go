package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func checkCliExists(command string) string {
	return execCommand(fmt.Sprintf("which %s", command))
}

func execCommand(command string) string {
	out, _ := exec.Command("sh", "-c", command).Output()

	return strings.TrimSuffix(string(out), "\n")
}
