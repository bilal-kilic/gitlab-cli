package utils

import (
	"os/exec"
	"strings"
)

func GetProjectName() string {
	command := exec.Command("git", "remote", "get-url", "origin")
	res, _ := command.Output()
	projectUrl := string(res)

	parts := strings.Split(projectUrl, "/")
	fileName := parts[len(parts)-1]
	return fileName[:strings.IndexByte(fileName, '.')]
}