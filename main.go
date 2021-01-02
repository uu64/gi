package main

import (
	"fmt"
	"os"

	cmdConfig "github.com/uu64/gi/lib/config"
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
	rc := cmdConfig.GetRemoteConfig()
	cmd := gi.NewGi(vcs, rc.Owner, rc.Repository, rc.Ref)

	// TODO: error handling
	gitignores, err := cmd.ListGitIgnorePath()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tc := cmdConfig.GetTuiConfig()
	selected := []string{}
	tui.ShowGitIgnoreOption(&gitignores, &selected, tc.PageSize)

	outputPath := ""
	tui.ShowOutputPathInput(&outputPath)

	cmd.Download(outputPath, selected)

	os.Exit(0)
}
