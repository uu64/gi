package main

import (
	"fmt"
	"os"

	"github.com/uu64/gi/lib/config"
	"github.com/uu64/gi/lib/gi"
	"github.com/uu64/gi/lib/github"
	"github.com/uu64/gi/lib/tui"
)

func fail(err error) {
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}

func main() {
	cfg := config.Get()
	vcs := github.NewGithub()
	cmd := gi.NewGi(vcs, cfg.Remote.Owner, cfg.Remote.Repository, cfg.Remote.Ref)

	gitignores, err := cmd.ListGitIgnorePath()
	if err != nil {
		fail(err)
	}

	selected := []string{}
	err = tui.ShowGitIgnoreOption(&gitignores, &selected, cfg.Tui.PageSize)

	outputPath := ""
	err = tui.ShowOutputPathInput(&outputPath)

	cmd.Download(outputPath, selected)

	os.Exit(0)
}
