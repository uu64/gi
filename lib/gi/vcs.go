package gi

import "context"

// VCS is the place that stores gitignore files
type VCS interface {
	ListAllFilePaths(ctx context.Context, owner, repo, ref, path string) (*[]string, error)
	GetFileContent(ctx context.Context, owner, repo, ref, path string) *string
}
