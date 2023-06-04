package utils

import (
	"fmt"
	"os"
)

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
