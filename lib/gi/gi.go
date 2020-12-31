package gi

import (
	"context"
	"strings"
)

const gitignoreExt = ".gitignore"
// var pathHashMap map[string]*string

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
			// pathHashMap[*content.Path]
			gitignores = append(gitignores, *content.Path)
		}
	}
	return gitignores, nil
}

// Download returns the list that contains the decoded content of gitignore.
func (gi *Gi) Download(shas []string) []*string {
	contents := []*string{}

	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	for _, sha := range shas {
		content, _ := gi.vcs.GetBlob(ctx, gi.owner, gi.repo, sha)
		contents = append(contents, content)
	}

	return contents
}
