package gi

import "context"

// Repository is the place that stores gitignore files
type Repository interface {
	ListAllFilePaths(ctx context.Context, owner, repo, ref, path string) *[]string
	GetFileContent(ctx context.Context, owner, repo, ref, path string) *string
}
