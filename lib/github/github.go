package github

import (
	"context"

	"github.com/google/go-github/v33/github"
	"github.com/uu64/gi/lib/gi"
)

// Github is implementation of the github repository.
type Github struct {
	client *github.Client
}

// New returns a Github object.
func New() *Github {
	return &Github{
		client: github.NewClient(nil),
	}
}

// ListAllContents returns the list that contains the Content object.
func (gh *Github) ListAllContents(ctx context.Context, owner, repo, ref, path string) []*gi.Content {
	contents := []*gi.Content{}
	opts := new(github.RepositoryContentGetOptions)
	opts.Ref = ref

	// TODO: error handling
	tree, _, _ := gh.client.Git.GetTree(ctx, owner, repo, ref, true)
	for _, entry := range tree.Entries {
		file := gi.Content{
			Type: getContentType(entry.GetType()),
			Path: entry.GetPath(),
		}
		contents = append(contents, &file)
	}
	return contents
}

// GetFileContent returns the decoded content of the specified file.
func (gh *Github) GetFileContent(ctx context.Context, owner, repo, ref, path string) *string {
	opts := new(github.RepositoryContentGetOptions)
	opts.Ref = ref

	// TODO: error handling
	content, _, _, _ := gh.client.Repositories.GetContents(ctx, owner, repo, path, opts)
	data, _ := content.GetContent()

	return &data
}

func getContentType(treeEntryType string) gi.ContentType {
	switch treeEntryType {
	case "blob":
		return gi.CtFile
	case "tree":
		return gi.CtDirectory
	case "commit":
		return gi.CtSubmodule
	case "symlink":
		return gi.CtSymLink
	default:
		// If TreeEntry is unknown object, this returns -1.
		return -1
	}
}
