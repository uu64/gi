package github

import (
	"context"
	"strings"

	"github.com/google/go-github/v33/github"
	"github.com/uu64/gi/lib/gi"
)

type Github struct {
	client *github.Client
}

func New() *Github {
	return &Github{
		client: github.NewClient(nil),
	}
}

func (gh *Github) GetAllFiles(ctx context.Context, owner, repo, ref, path string) []gi.File {
	contents := []gi.File{}
	opts := new(github.RepositoryContentGetOptions)
	opts.Ref = ref

	// TODO: error handling
	tree, _, _ := gh.client.Git.GetTree(ctx, owner, repo, ref, true)
	for _, entry := range tree.Entries {
		if strings.Compare(*entry.Type, "blob") == 0 {
			file := gi.File{
				Path: entry.GetPath(),
			}
			contents = append(contents, file)
		}
	}
	return contents
}

func (gh *Github) GetFileContent(ctx context.Context, owner, repo, ref, path string) *string {
	opts := new(github.RepositoryContentGetOptions)
	opts.Ref = ref

	// TODO: error handling
	content, _, _, _ := gh.client.Repositories.GetContents(ctx, owner, repo, path, opts)
	data, _ := content.GetContent()

	return &data
}
