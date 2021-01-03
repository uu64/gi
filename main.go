package main

import (
	"fmt"
	"os"

	"github.com/uu64/gi/lib/config"
	"github.com/uu64/gi/lib/gi"
	"github.com/uu64/gi/lib/github"
	"github.com/uu64/gi/lib/tui"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	cfg := config.Get()
	vcs := github.NewGithub()
	cmd := gi.NewGi(vcs, cfg.Remote.Owner, cfg.Remote.Repository, cfg.Remote.Ref)

	gitignores, err := cmd.ListGitIgnorePath()
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}

	selected := []string{}
	tui.ShowGitIgnoreOption(&gitignores, &selected, cfg.Tui.PageSize)

	outputPath := ""
	tui.ShowOutputPathInput(&outputPath)

	cmd.Download(outputPath, selected)

	os.Exit(0)
}
