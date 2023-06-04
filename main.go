package main

import (
	"fmt"
	"git_pull_all/utils"
)

func main() {
	bro := utils.IsGitRepo("./")
	fmt.Println(bro)
}
