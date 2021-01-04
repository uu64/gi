package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/uu64/gi/lib/config"
	"github.com/uu64/gi/lib/core"
	"github.com/uu64/gi/lib/github"
)

const (
	multiSelectMsg = "Select gitignore templates:"
	inputMsg       = "Input the output path (Existing file will be overwritten):"
	loadingMsg     = "Loading..."
	downloadingMsg = "Downloading..."
	cancelMsg      = "Canceled"
	// TODO: fix path format
	defaultOutputPath = "./.gitignore"
)

var s *spinner.Spinner

func init() {
	s = spinner.New(spinner.CharSets[14], 100*time.Millisecond)
}

// Start starts the gi command.
func Start(cfg *config.Config) {
	vcs := github.NewGithub()
	cmd := core.NewGi(vcs, cfg.Remote.Owner, cfg.Remote.Repository, cfg.Remote.Ref)

	var options []string
	var err error

	wait(loadingMsg, func() {
		if options, err = cmd.ListGitIgnorePath(); err != nil {
			fail(err)
		}
	})

	selected := []string{}
	showGitIgnoreOption(&options, &selected, cfg.Tui.PageSize)

	outputPath := ""
	showOutputPathInput(&outputPath)

	wait(downloadingMsg, func() {
		if err = cmd.Download(outputPath, selected); err != nil {
			fail(err)
		}
	})

	os.Exit(0)
}

func fail(err error) {
	fmt.Printf("%+v", err)
	os.Exit(1)
}

func cancelled() {
	if r := recover(); r != nil {
		// Command was cancelld on CTRL+C
		fmt.Println(cancelMsg)
		os.Exit(0)
	}
}

func wait(message string, fn func()) {
	s.Suffix = fmt.Sprintf(" %s", message)
	s.Start()

	fn()

	s.Stop()
}

func showGitIgnoreOption(gitignoreList, selected *[]string, pagesize int) error {
	prompt := &survey.MultiSelect{
		Message:  multiSelectMsg,
		Options:  *gitignoreList,
		PageSize: pagesize,
	}

	defer cancelled()
	if err := survey.AskOne(prompt, selected); err != nil {
		return fmt.Errorf("failed to show a multi-selection prompt: %w", err)
	}

	return nil
}

func showOutputPathInput(input *string) error {
	prompt := &survey.Input{
		Message: inputMsg,
		Default: defaultOutputPath,
	}

	defer cancelled()
	if err := survey.AskOne(prompt, input); err != nil {
		return fmt.Errorf("failed to show a text input: %w", err)
	}

	return nil
}
