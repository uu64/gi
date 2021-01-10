package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/jessevdk/go-flags"
	"github.com/uu64/gi/lib/config"
	"github.com/uu64/gi/lib/core"
	"github.com/uu64/gi/lib/github"
)

const (
	selectMsg         = "Select repository:"
	multiSelectMsg    = "Select gitignore templates:"
	inputMsg          = "Output path (Existing file will be overwritten):"
	loadingMsg        = "Loading..."
	downloadingMsg    = "Downloading..."
	cancelMsg         = "Canceled."
	successMsg        = "Complete."
	defaultOutputPath = ".gitignore"
	maxSelectLen      = 5
)

// Version will be set in build step.
var Version = "unset"

// Cmd is the object that has everything required to show CLI.
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
	return &Cmd{
		// The gi object will be set after user select a repository.
		gi:         nil,
		cfg:        cfg,
		spinner:    spinner.New(spinner.CharSets[14], 100*time.Millisecond),
		options:    new([]string),
		selected:   new([]string),
		outputPath: new(string),
	}
}

// Option is the object that represents command options.
type Option struct {
	Version bool `long:"version" description:"Show version"`
}

// Start starts the gi command.
func (cmd *Cmd) Start(args []string) {
	var opt Option
	var err error

	_, err = flags.ParseArgs(&opt, args)
	if err != nil {
		cmd.fail()
	}

	switch {
	case opt.Version:
		fmt.Println(Version)
		cmd.success()
	default:
		if err = cmd.showRepositoryOption(); err != nil {
			cmd.fail(fmt.Sprintf("%+v", err))
		}

		if err = cmd.loadOptions(); err != nil {
			cmd.fail(fmt.Sprintf("%+v", err))
		}

		if err = cmd.showGitIgnoreOption(); err != nil {
			cmd.fail(fmt.Sprintf("%+v", err))
		}

		if err = cmd.showOutputPathInput(); err != nil {
			cmd.fail(fmt.Sprintf("%+v", err))
		}

		if err = cmd.download(); err != nil {
			cmd.fail(fmt.Sprintf("%+v", err))
		}

		cmd.success(successMsg)
	}
}

func (cmd *Cmd) success(message ...string) {
	for _, s := range message {
		fmt.Println(s)
	}
	os.Exit(0)
}

func (cmd *Cmd) fail(message ...string) {
	for _, s := range message {
		fmt.Println(s)
	}
	os.Exit(1)
}

func (cmd *Cmd) canceled(message ...string) {
	if r := recover(); r != nil {
		for _, s := range message {
			fmt.Println(s)
		}
		// Command was cancelld on CTRL+C
		os.Exit(0)
	}
}

func (cmd *Cmd) startSpinner(message string) {
	cmd.spinner.Suffix = fmt.Sprintf(" %s", message)
	cmd.spinner.Start()
}

func (cmd *Cmd) stopSpinner() {
	cmd.spinner.Stop()
}

func (cmd *Cmd) loadOptions() error {
	if cmd.gi == nil {
		return fmt.Errorf("repository is not selected")
	}

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
	if *cmd.outputPath == "" {
		return fmt.Errorf("output path is not selected")
	}

	cmd.startSpinner(downloadingMsg)
	defer cmd.stopSpinner()

	wd, err := os.Getwd()
	if err != nil {
		cmd.fail(fmt.Sprintf("%+v", fmt.Errorf("failed to get working directory: %w", err)))
	}

	f, err := os.Create(filepath.Join(wd, *cmd.outputPath))
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	err = cmd.gi.Download(*cmd.selected, f)
	if err != nil {
		return err
	}

	return nil
}

func (cmd *Cmd) showRepositoryOption() error {
	options := []string{}
	for _, repo := range cmd.cfg.Repos {
		options = append(options,
			fmt.Sprintf("%s/%s (%s)", repo.Owner, repo.Name, repo.Branch))
	}

	selected := 0
	// If there is only 1 repository, selection prompt does not shown.
	if len(cmd.cfg.Repos) != 1 {
		prompt := &survey.Select{
			Message:  selectMsg,
			Options:  options,
			PageSize: cmd.cfg.Cli.PageSize,
		}

		defer cmd.canceled()
		err := survey.AskOne(prompt, &selected)
		if err != nil {
			return fmt.Errorf("failed to show a selection prompt: %w", err)
		}
	}

	repo := cmd.cfg.Repos[selected]
	cmd.gi = core.NewGi(
		github.NewRepository(repo.Owner, repo.Name, repo.Branch, cmd.cfg.Auth.Token))
	return nil
}

func (cmd *Cmd) showGitIgnoreOption() error {
	prompt := &survey.MultiSelect{
		Message:  multiSelectMsg,
		Options:  *cmd.options,
		PageSize: cmd.cfg.Cli.PageSize,
	}
	validator := func(ans interface{}) error {
		// val's Kind should be Slice
		val := reflect.ValueOf(ans)

		if val.Len() == 0 {
			return fmt.Errorf("value is required")
		}
		if val.Len() > maxSelectLen {
			return fmt.Errorf("you can select up to %d", maxSelectLen)
		}
		return nil
	}

	defer cmd.canceled(cancelMsg)
	err := survey.AskOne(prompt, cmd.selected, survey.WithValidator(validator))
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

	defer cmd.canceled(cancelMsg)
	err := survey.AskOne(prompt, cmd.outputPath)
	if err != nil {
		return fmt.Errorf("failed to show a text input: %w", err)
	}

	return nil
}
