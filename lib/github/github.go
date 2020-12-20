package github

import (
	"context"

	"github.com/google/go-github/v33/github"
)

type Github struct {
	client *github.Client
}

func New() *Github {
	return &Github{
		client: github.NewClient(nil),
	}
}

func (vcs *Github) GetAllContentPaths(ctx context.Context, owner, repo, ref, path string) []string {
	contents := []string{}
	opts := new(github.RepositoryContentGetOptions)
	opts.Ref = ref

	tree, _, _ := vcs.client.Git.GetTree(ctx, owner, repo, ref, true)
	for _, entry := range tree.Entries {
		contents = append(contents, entry.GetPath())
	}
	return contents
}
