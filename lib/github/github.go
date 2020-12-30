package github

import (
	"context"
	"sort"

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

// ListAllFilePaths returns a pointer to a slice containing the paths of all files sorted by ascii code.
// The slice does not include the paths of objects other than files (ex: directories, submodules...).
func (gh *Github) ListAllFilePaths(ctx context.Context, owner, repo, ref, path string) *[]string {
	contents := []string{}
	opts := new(github.RepositoryContentGetOptions)
	opts.Ref = ref

	// TODO: error handling
	tree, _, _ := gh.client.Git.GetTree(ctx, owner, repo, ref, true)
	for _, entry := range tree.Entries {
		ct := getContentType(entry.GetType())
		if ct == gi.CtFile {
			contents = append(contents, entry.GetPath())
		}
	}

	sort.Slice(contents, func(i, j int) bool {
		return contents[i] < contents[j]
	})
	return &contents
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
