package github

import (
	"context"
	"encoding/base64"
	"net/http"
	"sort"

	"github.com/google/go-github/v33/github"
	"github.com/uu64/gi/lib/core"
	"golang.org/x/oauth2"
)

// Repository is implementation of the github repository.
type Repository struct {
	client *github.Client
	owner  string
	name   string
	branch string
}

// NewRepository returns a new Github object.
func NewRepository(owner, name, branch, token string) *Repository {
	hc := new(http.Client)

	if token != "" {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		hc = oauth2.NewClient(ctx, ts)
	}

	return &Repository{
		client: github.NewClient(hc),
		owner:  owner,
		name:   name,
		branch: branch,
	}
}

// GetTree returns contents sorted by the path.
func (gh *Repository) GetTree(ctx context.Context, recursive bool) ([]*core.TreeNode, error) {
	contents := []*core.TreeNode{}

	branch, _, err := gh.client.Repositories.GetBranch(ctx, gh.owner, gh.name, gh.branch)
	if err != nil {
		return nil, err
	}

	treeSHA := branch.Commit.Commit.Tree.SHA
	tree, _, err := gh.client.Git.GetTree(ctx, gh.owner, gh.name, *treeSHA, recursive)
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
func (gh *Repository) GetBlob(ctx context.Context, sha string) ([]byte, error) {
	blob, _, err := gh.client.Git.GetBlob(ctx, gh.owner, gh.name, sha)
	if err != nil {
		return nil, err
	}

	content := blob.GetContent()
	bytes, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func getNodeType(treeEntryType string) core.NodeType {
	nodeTypeMap := map[string]core.NodeType{
		"blob":   core.NtBlob,
		"tree":   core.NtTree,
		"commit": core.NtSubmodule,
	}
	return nodeTypeMap[treeEntryType]
}
