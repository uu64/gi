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

// New returns a Gi object.
// TODO: refactor
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
		return nil, err
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
func (gi *Gi) Download(outputPath string, selected []string) []*string {
	contents := []*string{}

	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	for _, item := range selected {
		sha := pathHashMap[item]
		content, _ := gi.vcs.GetBlob(ctx, gi.owner, gi.repo, *sha)
		contents = append(contents, content)
	}
	gi.write(outputPath, contents)

	return contents
}

func (gi *Gi) write(path string, contents []*string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	writer := bufio.NewWriter(file)

	for _, content := range contents {
		fmt.Println(*content)
		// TODO: error handling
		_, err := writer.WriteString(*content)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	writer.Flush()
	return nil
}
