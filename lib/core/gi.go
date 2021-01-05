package core

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"strings"
)

// Repository is the place that stores gitignore files
type Repository interface {
	// GetTree returns contents sorted by the path.
	GetTree(ctx context.Context, recursive bool) ([]*TreeNode, error)
	// GetBlob returns the decoded content of the specified SHA.
	GetBlob(ctx context.Context, sha string) ([]byte, error)
}

// Gi is the object to handle data of the remote repository.
type Gi struct {
	repository Repository
	owner      string
	repo       string
	ref        string
}

// NewGi returns a new Gi object.
func NewGi(repo Repository) *Gi {
	gi := Gi{
		repository: repo,
	}
	return &gi
}

const gitignoreExt = ".gitignore"

var pathHashMap = make(map[string]*string)

// ListGitIgnorePath returns the list that contains the filepath of gitignore.
func (gi *Gi) ListGitIgnorePath() ([]string, error) {
	gitignores := []string{}

	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()
	contents, err := gi.repository.GetTree(ctx, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get a list of remote objects: %w", err)
	}

	for _, content := range contents {
		path := *content.Path
		if path != gitignoreExt && strings.HasSuffix(path, gitignoreExt) {
			pathHashMap[path] = content.SHA
			gitignores = append(gitignores, path)
		}
	}
	return gitignores, nil
}

// Download get the content of selected files and writes them merged.
func (gi *Gi) Download(selected []string, w io.Writer) error {
	writer := bufio.NewWriter(w)

	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	for _, item := range selected {
		sha := pathHashMap[item]
		content, err := gi.repository.GetBlob(ctx, *sha)
		if err != nil {
			return fmt.Errorf("failed to get a blob: %w", err)
		}
		writer.WriteString(fmt.Sprintf("# %s\n", item))
		writer.Write(content)
	}

	err := writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to write a file: %w", err)
	}
	return nil
}
