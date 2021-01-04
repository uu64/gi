package core

import "context"

// VCS is the place that stores gitignore files
type VCS interface {
	// GetTree returns a slice of contents sorted by the path.
	GetTree(ctx context.Context, owner, repo, ref string, recursive bool) ([]*TreeNode, error)
	// GetBlob returns the decoded content of the specified SHA.
	GetBlob(ctx context.Context, owner, repo, sha string) (*string, error)
}
