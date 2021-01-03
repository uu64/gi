package tui

import (
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
)

const (
	multiSelectMsg    = "Select gitignore templates:"
	inputMsg          = "Input the output path (Existing file will be overwritten):"
	defaultOutputPath = "./.gitignore"
)

var s *spinner.Spinner

func init() {
	s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
}

func cancelled() {
	if r := recover(); r != nil {
		// Command was cancelld on CTRL+C
		fmt.Println("Cancelled.")
		os.Exit(1)
	}
}

// ShowGitIgnoreOption shows a multi-selection prompt to select gitignores.
func ShowGitIgnoreOption(gitignoreList, selected *[]string, pagesize int) error {
	prompt := &survey.MultiSelect{
		Message:  multiSelectMsg,
		Options:  *gitignoreList,
		PageSize: pagesize,
	}

	defer cancelled()
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

	defer cancelled()
	err := survey.AskOne(prompt, input)
	if err != nil {
		return fmt.Errorf("failed to show a text input: %w", err)
	}

	return nil
}

// StartSpinner starts the indicator.
func StartSpinner(message string) {
	s.Suffix = message
	s.Start()
}

// StopSpinner stops the indicator.
func StopSpinner() {
	s.Stop()
}
