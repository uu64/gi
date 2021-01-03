package tui

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

const (
	multiSelectMsg    = "Select gitignore templates:"
	inputMsg          = "Output path:"
	defaultOutputPath = "./.gitignore"
)

// ShowGitIgnoreOption shows a multi-selection prompt to select gitignores.
func ShowGitIgnoreOption(gitignoreList, selected *[]string, pagesize int) error {
	prompt := &survey.MultiSelect{
		Message:  multiSelectMsg,
		Options:  *gitignoreList,
		PageSize: pagesize,
	}

	err := survey.AskOne(prompt, selected)
	if err != nil {
		return fmt.Errorf("failed to show a multi-selection prompt: %w", err)
	}

	return nil
}

// ShowOutputPathInput show a text input to input the output path.
func ShowOutputPathInput(input *string) error {
	prompt := &survey.Input{
		Message: inputMsg,
		Default: defaultOutputPath,
	}

	err := survey.AskOne(prompt, input)
	if err != nil {
		return fmt.Errorf("failed to show a text input: %w", err)
	}

	return nil
}
