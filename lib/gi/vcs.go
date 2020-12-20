package gi

import "context"

type VCS interface {
	GetAllFiles(ctx context.Context, owner, repo, ref, path string) []File
	GetFileContent(ctx context.Context, owner, repo, ref, path string) *string
}
