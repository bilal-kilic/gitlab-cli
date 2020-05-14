package utils

import (
	"errors"
	"os/exec"
	"strings"
)

func GetProjectName() (string, error) {
	command := exec.Command("git", "remote", "get-url", "origin")
	res, _ := command.Output()
	projectUrl := string(res)

	if projectUrl == "" {
		return "", errors.New("Not a git repository")
	}

	parts := strings.Split(projectUrl, "/")
	fileName := parts[len(parts)-1]
	return fileName[:strings.IndexByte(fileName, '.')], nil
}