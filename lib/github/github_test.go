package github

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAllContents(t *testing.T) {
	expected := []string{
		"README.md",
		".gitmodules",
		"LICENSE",
		"testdocument.txt",
		"docs/testdocument.txt",
	}
	gh := New()

	t.Run("can list all file paths sorted by ascii code.", func(t *testing.T) {
		ctx := context.Background()
		owner := "uu64"
		repo := "ghapi-test"
		ref := "main"
		path := "/"

		paths := gh.ListAllFilePaths(ctx, owner, repo, ref, path)
		sort.Slice(expected, func(i, j int) bool {
			return expected[i] < expected[j]
		})
		assert.Equal(t, expected, *paths)
	})
}
