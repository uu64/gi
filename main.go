package main

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/uu64/gi/lib/gi"
	"github.com/uu64/gi/lib/github"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			os.Exit(1)
		}
	}()
	repository := github.New()

	// TODO: Should be loaded from config
	owner := "github"
	repo := "gitignore"
	path := "/"
	ref := "master"
	cmd := gi.New(repository, owner, repo, path, ref)

	days := []string{}
	prompt := &survey.MultiSelect{
		Message: "What days do you prefer:",
		Options: cmd.ListGitIgnorePath(),
	}
	survey.AskOne(prompt, &days)

	os.Exit(0)
}