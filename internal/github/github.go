package github

import (
	"context"

	"github.com/google/go-github/v33/github"
)

var client *github.Client

func init() {
	client = github.NewClient(nil)
}

func GetGitIgnores() {
	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	// TODO: Should be loaded from config
	owner := "github"
	repo := "gitignore"
	path := "/"

	opts := new(github.RepositoryContentGetOptions)
	opts.Ref = "master"

	_, contents, _, _ := client.Repositories.GetContents(ctx, owner, repo, path, opts)
	for _, content := range contents {
		println(content.GetName())
	}
}
