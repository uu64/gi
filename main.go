package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/uu64/gi/lib/gi"
	"github.com/uu64/gi/lib/github"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	vcs := github.NewGithub()

	// TODO: Should be loaded from config
	owner := "github"
	repo := "gitignore"
	ref := "master"
	cmd := gi.NewGi(vcs, owner, repo, ref)

	// TODO: error handling
	gitignores, err := cmd.ListGitIgnorePath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	selected := []string{}
	multiSelection := &survey.MultiSelect{
		Message: "Select gitignore templates:",
		Options: gitignores,
	}
	survey.AskOne(multiSelection, &selected)

	outputPath := ""
	input := &survey.Input{
		Message: "Output path:",
		Default: "./.gitignore",
	}
	survey.AskOne(input, &outputPath)
	fmt.Println(outputPath)

	contents := cmd.Download(selected)
	for _, content := range contents {
		fmt.Println(*content)
	}

	os.Exit(0)
}
