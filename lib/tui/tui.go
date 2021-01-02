package tui

import "github.com/AlecAivazis/survey/v2"

const (
	multiSelectMsg    = "Select gitignore templates:"
	inputMsg          = "Output path:"
	defaultOutputPath = "./.gitignore"
)

func ShowGitIgnoreOption(gitignoreList, selected *[]string, pagesize int) error {
	prompt := &survey.MultiSelect{
		Message:  multiSelectMsg,
		Options:  *gitignoreList,
		PageSize: pagesize,
	}

	err := survey.AskOne(prompt, selected)
	if err != nil {
		return err
	}

	return nil
}

func ShowOutputPathInput(input *string) error {
	prompt := &survey.Input{
		Message: inputMsg,
		Default: defaultOutputPath,
	}

	err := survey.AskOne(prompt, input)
	if err != nil {
		return err
	}

	return nil
}
