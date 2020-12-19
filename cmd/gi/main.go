package main

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/uu64/gi/internal/github"
)

func main() {
	days := []string{}
	prompt := &survey.MultiSelect{
		Message: "What days do you prefer:",
		Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	}
	survey.AskOne(prompt, &days)

	github.GetGitIgnores()
}
