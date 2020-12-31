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

	t.Run("can list all file paths sorted by ascii code", func(t *testing.T) {
		ctx := context.Background()
		owner := "uu64"
		repo := "ghapi-test"
		ref := "main"
		path := "/"

		paths, err := gh.ListAllFilePaths(ctx, owner, repo, ref, path)
		sort.Slice(expected, func(i, j int) bool {
			return expected[i] < expected[j]
		})

		if assert.NoError(t, err) {
			assert.Equal(t, expected, *paths)
		}
	})

	t.Run("get an error when non-existent repository is specified", func(t *testing.T) {
		ctx := context.Background()
		owner := "uu64"
		repo := "non-existent"
		ref := "main"
		path := "/"

		_, err := gh.ListAllFilePaths(ctx, owner, repo, ref, path)
		assert.Error(t, err)
	})
}
