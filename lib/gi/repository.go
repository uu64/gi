package gi

import "context"

// Repository is the place that stores gitignore files
type Repository interface {
	ListAllContents(ctx context.Context, owner, repo, ref, path string) []*Content
	GetFileContent(ctx context.Context, owner, repo, ref, path string) *string
}
