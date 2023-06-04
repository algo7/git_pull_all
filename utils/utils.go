package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// IsGitRepo check if the given folder is a github repository
func IsGitRepo(path string) bool {

	// Read the directory content of the given path
	entry, err := os.ReadDir(path)

	if err != nil {
		return false
	}

	// Loop through the /folder names
	for _, v := range entry {

		// If the .git folder is found, read the config folder inside
		if v.Name() == ".git" {
			content, err := os.ReadFile(path + "/.git/config")

			if err != nil {
				return false
			}

			// show file content
			fmt.Println(string(content))

			return true
		}
	}
	return false
}

// commandExecuteSuccessfully execute a command and return true if it was successful
func commandExecuteSuccessfully(command string, args []string) bool {

	cmd := exec.Command(command, args...)

	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

// GitIsInstalled check if git is installed
func GitIsInstalled() {
	res := commandExecuteSuccessfully("git", []string{"version"})

	if res == false {
		fmt.Println("Git is not installed")
		os.Exit(1)
	}

	fmt.Println("Git is installed")
}
