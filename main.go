package main

import (
	"fmt"
	"os"

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
	vcs := github.NewGithub()

	// TODO: Should be loaded from config
	owner := "github"
	repo := "gitignore"
	ref := "master"
	cmd := gi.NewGi(vcs, owner, repo, ref)

	// TODO: error handling
	gitignores, err := cmd.ListGitIgnorePath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	selected := []string{}
	tui.ShowGitIgnoreOption(&gitignores, &selected)

	outputPath := ""
	tui.ShowOutputPathInput(&outputPath)

	cmd.Download(outputPath, selected)

	os.Exit(0)
}
