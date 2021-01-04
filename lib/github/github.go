package github

import (
	"context"
	"encoding/base64"
	"sort"

	"github.com/google/go-github/v33/github"
	"github.com/uu64/gi/lib/core"
)

// Github is implementation of the github repository.
type Github struct {
	client *github.Client
}

// NewGithub returns a new Github object.
func NewGithub() *Github {
	return &Github{
		client: github.NewClient(nil),
	}
}

// GetTree returns a slice of contents sorted by the path.
func (gh *Github) GetTree(ctx context.Context, owner, repo, ref string, recursive bool) ([]*core.TreeNode, error) {
	contents := []*core.TreeNode{}
	opts := new(github.RepositoryContentGetOptions)
	opts.Ref = ref

	tree, _, err := gh.client.Git.GetTree(ctx, owner, repo, ref, recursive)
	if err != nil {
		return nil, err
	}

	for _, entry := range tree.Entries {
		node := core.NewTreeNode(getNodeType(entry.GetType()), entry.GetPath(), entry.GetSHA())
		contents = append(contents, node)
	}

	sort.Slice(contents, func(i, j int) bool {
		return *contents[i].Path < *contents[j].Path
	})
	return contents, nil
}

// GetBlob returns the decoded content of the specified SHA.
func (gh *Github) GetBlob(ctx context.Context, owner, repo, sha string) (*string, error) {
	blob, _, err := gh.client.Git.GetBlob(ctx, owner, repo, sha)
	if err != nil {
		return nil, err
	}

	content := blob.GetContent()
	bytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return nil, err
	}

	decodedContent := string(bytes)
	return &decodedContent, nil
}

func getNodeType(treeEntryType string) core.NodeType {
	nodeTypeMap := map[string]core.NodeType{
		"blob":   core.NtBlob,
		"tree":   core.NtTree,
		"commit": core.NtSubmodule,
	}
	return nodeTypeMap[treeEntryType]
}
