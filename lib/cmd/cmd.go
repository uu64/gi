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
	inputMsg       = "Output path (Existing file will be overwritten):"
	loadingMsg     = "Loading..."
	downloadingMsg = "Downloading..."
	cancelMsg      = "Canceled."
	successMsg     = "Complete."
	// TODO: fix path format
	defaultOutputPath = "./.gitignore"
)

// Cmd is the object that has everything required to show tui.
type Cmd struct {
	gi         *core.Gi
	cfg        *config.Config
	spinner    *spinner.Spinner
	options    *[]string
	selected   *[]string
	outputPath *string
}

// NewCmd returns a new Cmd object.
func NewCmd() *Cmd {
	cfg := config.Get()
	vcs := github.NewGithub()
	return &Cmd{
		gi:         core.NewGi(vcs, cfg.Remote.Owner, cfg.Remote.Repository, cfg.Remote.Ref),
		cfg:        cfg,
		spinner:    spinner.New(spinner.CharSets[14], 100*time.Millisecond),
		options:    new([]string),
		selected:   new([]string),
		outputPath: new(string),
	}
}

func (cmd *Cmd) startSpinner(message string) {
	cmd.spinner.Suffix = fmt.Sprintf(" %s", message)
	cmd.spinner.Start()
}

func (cmd *Cmd) stopSpinner() {
	cmd.spinner.Stop()
}

func (cmd *Cmd) fail(err error) {
	fmt.Printf("%+v", err)
	os.Exit(1)
}

func (cmd *Cmd) canceled() {
	if r := recover(); r != nil {
		// Command was cancelld on CTRL+C
		fmt.Println(cancelMsg)
		os.Exit(0)
	}
}

func (cmd *Cmd) success() {
	fmt.Println(successMsg)
	os.Exit(0)
}

// Start starts the gi command.
func (cmd *Cmd) Start() {
	var err error

	if err = cmd.loadOptions(); err != nil {
		cmd.fail(err)
	}

	if err = cmd.showGitIgnoreOption(); err != nil {
		cmd.fail(err)
	}

	if err = cmd.showOutputPathInput(); err != nil {
		cmd.fail(err)
	}

	if err = cmd.download(); err != nil {
		cmd.fail(err)
	}

	cmd.success()
}

func (cmd *Cmd) loadOptions() error {
	cmd.startSpinner(loadingMsg)
	defer cmd.stopSpinner()

	res, err := cmd.gi.ListGitIgnorePath()
	if err != nil {
		return err
	}

	cmd.options = &res
	return nil
}

func (cmd *Cmd) download() error {
	cmd.startSpinner(downloadingMsg)
	defer cmd.stopSpinner()

	return cmd.gi.Download(*cmd.outputPath, *cmd.selected)
}

func (cmd *Cmd) showGitIgnoreOption() error {
	prompt := &survey.MultiSelect{
		Message:  multiSelectMsg,
		Options:  *cmd.options,
		PageSize: cmd.cfg.Tui.PageSize,
	}

	defer cmd.canceled()
	err := survey.AskOne(prompt, cmd.selected, survey.WithValidator(survey.Required))
	if err != nil {
		return fmt.Errorf("failed to show a multi-selection prompt: %w", err)
	}

	return nil
}

func (cmd *Cmd) showOutputPathInput() error {
	prompt := &survey.Input{
		Message: inputMsg,
		Default: defaultOutputPath,
	}

	defer cmd.canceled()
	err := survey.AskOne(prompt, cmd.outputPath)
	if err != nil {
		return fmt.Errorf("failed to show a text input: %w", err)
	}

	return nil
}
