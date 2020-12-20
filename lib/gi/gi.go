package gi

import (
	"context"
	"strings"

	"github.com/uu64/gi/lib/github"
)

type VCS interface {
	GetAllContentPaths(ctx context.Context, owner, repo, ref, path string) []string
}

type Gi struct {
	vcs VCS
}

func New() *Gi {
	gi := Gi{
		vcs: github.New(),
	}
	return &gi
}

func (gi *Gi) ListGitIgnore() []string {
	gitignores := []string{}
	// TODO: Should be reconsidered if it is empty
	ctx := context.Background()

	// TODO: Should be loaded from config
	owner := "github"
	repo := "gitignore"
	path := "/"
	branch := "master"

	paths := gi.vcs.GetAllContentPaths(ctx, owner, repo, branch, path)

	for _, path := range paths {
		if strings.HasSuffix(path, "gitignore") {
			gitignores = append(gitignores, path)
		}
	}
	return gitignores
}

func Download(query []string) {

}
