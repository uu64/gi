package gi

import (
	"context"
	"strings"
)

const gitignoreExt = ".gitignore"

// Gi is the object containing everything required to run gi.
type Gi struct {
	vcs   VCS
	owner string
	repo  string
	path  string
	ref   string
}

// New returns a Gi object.
func New(vcs VCS, owner, repo, path, ref string) *Gi {
	gi := Gi{
		vcs:   vcs,
		owner: owner,
		repo:  repo,
		path:  path,
		ref:   ref,
	}
	return &gi
}

// ListGitIgnorePath returns the list that contains the filepath of gitignore.
func (gi *Gi) ListGitIgnorePath() ([]string, error) {
	gitignores := []string{}

	ctx := context.Background()
	paths, err := gi.vcs.ListAllFilePaths(ctx, gi.owner, gi.repo, gi.ref, gi.path)
	if err != nil {
		return nil, err
	}

	for _, path := range *paths {
		if strings.HasSuffix(path, gitignoreExt) {
			gitignores = append(gitignores, path)
		}
	}
	return gitignores, nil
}

// Download returns the list that contains the decoded content of gitignore.
func (gi *Gi) Download(paths []string) []*string {
	contents := []*string{}

	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	for _, path := range paths {
		content := gi.vcs.GetFileContent(ctx, gi.owner, gi.repo, gi.ref, path)
		contents = append(contents, content)
	}

	return contents
}
