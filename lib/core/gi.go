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
	repository  Repository
	pathHashMap map[string]*string
}

// NewGi returns a new Gi object.
func NewGi(repo Repository) *Gi {
	gi := Gi{
		repository:  repo,
		pathHashMap: make(map[string]*string),
	}
	return &gi
}

// ListGitIgnorePath returns the list that contains the filepath of gitignore.
func (gi *Gi) ListGitIgnorePath() ([]string, error) {
	gitignores := []string{}
	gitignoreExt := ".gitignore"

	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	contents, err := gi.repository.GetTree(ctx, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get a list of remote objects: %w", err)
	}

	for _, content := range contents {
		path := *content.Path
		if path != gitignoreExt && strings.HasSuffix(path, gitignoreExt) {
			gi.pathHashMap[path] = content.SHA
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
		sha := gi.pathHashMap[item]
		content, err := gi.repository.GetBlob(ctx, *sha)
		if err != nil {
			return fmt.Errorf("failed to get a blob: %w", err)
		}
		writer.WriteString(fmt.Sprintf("# %s\n", item))
		writer.Write(content)
		writer.WriteString(fmt.Sprintf("\n"))
	}

	err := writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to write a file: %w", err)
	}
	return nil
}
