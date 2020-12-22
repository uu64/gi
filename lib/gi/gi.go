package gi

import (
	"context"
	"strings"
)

const gitignoreExt = "gitignore"

// Gi is the object containing everything required to run gi.
type Gi struct {
	repository Repository
	owner      string
	repo       string
	path       string
	ref        string
}

// New returns a Gi object.
func New(repository Repository, owner, repo, path, ref string) *Gi {
	gi := Gi{
		repository: repository,
		owner:      owner,
		repo:       repo,
		path:       path,
		ref:        ref,
	}
	return &gi
}

// ListGitIgnorePath returns the list that contains the filepath of gitignore.
func (gi *Gi) ListGitIgnorePath() []string {
	gitignores := []string{}

	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	files := gi.repository.ListAllContents(ctx, gi.owner, gi.repo, gi.ref, gi.path)

	for _, file := range files {
		if strings.HasSuffix(file.Path, gitignoreExt) {
			gitignores = append(gitignores, file.Path)
		}
	}
	return gitignores
}

// Download returns the list that contains the decoded content of gitignore.
func (gi *Gi) Download(paths []string) []*string {
	contents := []*string{}

	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	for _, path := range paths {
		content := gi.repository.GetFileContent(ctx, gi.owner, gi.repo, gi.ref, path)
		contents = append(contents, content)
	}

	return contents
}
