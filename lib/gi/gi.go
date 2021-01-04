package gi

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

const gitignoreExt = ".gitignore"

var pathHashMap = make(map[string]*string)

// Gi is the object containing everything required to run gi.
type Gi struct {
	vcs   VCS
	owner string
	repo  string
	ref   string
}

// NewGi returns a new Gi object.
func NewGi(vcs VCS, owner, repo, ref string) *Gi {
	gi := Gi{
		vcs:   vcs,
		owner: owner,
		repo:  repo,
		ref:   ref,
	}
	return &gi
}

// ListGitIgnorePath returns the list that contains the filepath of gitignore.
func (gi *Gi) ListGitIgnorePath() ([]string, error) {
	gitignores := []string{}

	ctx := context.Background()
	contents, err := gi.vcs.GetTree(ctx, gi.owner, gi.repo, gi.ref, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get a list of remote objects: %w", err)
	}

	for _, content := range contents {
		if strings.HasSuffix(*content.Path, gitignoreExt) {
			pathHashMap[*content.Path] = content.SHA
			gitignores = append(gitignores, *content.Path)
		}
	}
	return gitignores, nil
}

// Download returns the list that contains the decoded content of gitignore.
func (gi *Gi) Download(outputPath string, selected []string) error {
	contents := []*string{}

	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	for _, item := range selected {
		sha := pathHashMap[item]
		content, err := gi.vcs.GetBlob(ctx, gi.owner, gi.repo, *sha)
		if err != nil {
			return fmt.Errorf("failed to get a blob: %w", err)
		}
		contents = append(contents, content)
	}

	err := gi.write(outputPath, contents)
	if err != nil {
		return fmt.Errorf("failed to write a file: %w", err)
	}

	return nil
}

func (gi *Gi) write(path string, contents []*string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, content := range contents {
		_, err := writer.WriteString(*content)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}
