package gi

import (
	"context"
	"strings"
)

const gitignoreExt = "gitignore"

type Gi struct {
	vcs   VCS
	owner string
	repo  string
	path  string
	ref   string
}

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

func (gi *Gi) ListGitIgnore() []string {
	gitignores := []string{}

	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	files := gi.vcs.GetAllFiles(ctx, gi.owner, gi.repo, gi.ref, gi.path)

	for _, file := range files {
		if strings.HasSuffix(file.Path, gitignoreExt) {
			gitignores = append(gitignores, file.Path)
		}
	}
	return gitignores
}

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
