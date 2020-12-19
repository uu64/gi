package github

import (
	"context"
	"sync"

	"github.com/google/go-github/v33/github"
)

var once sync.Once
var client *github.Client

func getClient() *github.Client {
	once.Do(func() {
		client = github.NewClient(nil)
	})
	return client
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
		println(content.GetName)
	}
}
