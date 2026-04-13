package main

import (
	githubTypes "github.com/NateMartes/action-to-gitlab-ci/pkg/types/github"
	"os"
	"fmt"
)

func main() {
	fileBytes, _ := os.ReadFile(os.Args[1])
	workflow, err := githubTypes.UnmarshalGithubWorkflow(fileBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(workflow);
}