package main

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/uu64/gi/lib/gi"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			os.Exit(1)
		}
	}()
	cmd := gi.New()

	days := []string{}
	prompt := &survey.MultiSelect{
		Message: "What days do you prefer:",
		Options: cmd.ListGitIgnore(),
	}
	survey.AskOne(prompt, &days)

	os.Exit(0)
}
