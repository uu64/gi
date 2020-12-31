package github

import (
	"context"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uu64/gi/lib/gi"
)

func TestNewGithub(t *testing.T) {
	t.Run("can get a new Github object", func(t *testing.T) {
		// TODO: add test
		t.Skip()
	})
}

func TestGetTree(t *testing.T) {
	gh := NewGithub()

	t.Run("can get contents sorted by path (not recursively)", func(t *testing.T) {
		expected := []*gi.TreeNode{
			gi.NewTreeNode(
				gi.NtBlob,
				"README.md",
				"6799dc110eaa15880d42c0b447013c30422dc527"),
			gi.NewTreeNode(
				gi.NtBlob,
				".gitmodules",
				"1a41f17d2c803f43f316adb1cf07e6b9cbfbda3d"),
			gi.NewTreeNode(
				gi.NtBlob,
				"LICENSE",
				"c46410aef77ff7fbe2f44815ac5ca8fb351d190f"),
			gi.NewTreeNode(
				gi.NtBlob,
				"testdocument.txt",
				"d6455f8dde4bd5300e65467b170f7485f4ab77e7"),
			gi.NewTreeNode(
				gi.NtTree,
				"docs",
				"633ec3e50dafa8ca1d9b87763fd2bf1f90dd77b1"),
			gi.NewTreeNode(
				gi.NtSubmodule,
				"ghapi-test",
				"017bdc26adbd7c544b2180ab947857ff98d8434f"),
		}
		ctx := context.Background()
		owner := "uu64"
		repo := "ghapi-test"
		ref := "main"
		recursive := false

		contents, err := gh.GetTree(ctx, owner, repo, ref, recursive)
		sort.Slice(expected, func(i, j int) bool {
			return *expected[i].Path < *expected[j].Path
		})

		if assert.NoError(t, err) {
			for i, content := range contents {
				assert.Equal(t, *expected[i].Path, *content.Path)
				assert.Equal(t, expected[i].Type, content.Type)
			}
		}
	})

	t.Run("can get all contents sorted by the path", func(t *testing.T) {
		expected := []*gi.TreeNode{
			gi.NewTreeNode(
				gi.NtBlob,
				"README.md",
				"6799dc110eaa15880d42c0b447013c30422dc527"),
			gi.NewTreeNode(
				gi.NtBlob,
				".gitmodules",
				"1a41f17d2c803f43f316adb1cf07e6b9cbfbda3d"),
			gi.NewTreeNode(
				gi.NtBlob,
				"LICENSE",
				"c46410aef77ff7fbe2f44815ac5ca8fb351d190f"),
			gi.NewTreeNode(
				gi.NtBlob,
				"testdocument.txt",
				"d6455f8dde4bd5300e65467b170f7485f4ab77e7"),
			gi.NewTreeNode(
				gi.NtBlob,
				"docs/testdocument.txt",
				"a0f31e800f7bb4493ad94210b9f1770f6334531f"),
			gi.NewTreeNode(
				gi.NtTree,
				"docs",
				"633ec3e50dafa8ca1d9b87763fd2bf1f90dd77b1"),
			gi.NewTreeNode(
				gi.NtSubmodule,
				"ghapi-test",
				"017bdc26adbd7c544b2180ab947857ff98d8434f"),
		}
		ctx := context.Background()
		owner := "uu64"
		repo := "ghapi-test"
		ref := "main"
		recursive := true

		contents, err := gh.GetTree(ctx, owner, repo, ref, recursive)
		sort.Slice(expected, func(i, j int) bool {
			return *expected[i].Path < *expected[j].Path
		})

		if assert.NoError(t, err) {
			for i, content := range contents {
				assert.Equal(t, *expected[i].Path, *content.Path)
				assert.Equal(t, expected[i].Type, content.Type)
			}
		}
	})

	t.Run("get an error when non-existent repository is specified", func(t *testing.T) {
		ctx := context.Background()
		owner := "uu64"
		repo := "non-existent"
		ref := "main"
		recursive := false

		_, err := gh.GetTree(ctx, owner, repo, ref, recursive)
		assert.Error(t, err)
	})
}

func TestGetBlob(t *testing.T) {
	gh := NewGithub()

	t.Run("can get a blob", func(t *testing.T) {
		ctx := context.Background()
		owner := "uu64"
		repo := "ghapi-test"
		// SHA of docs/testdocument.txt
		sha := "a0f31e800f7bb4493ad94210b9f1770f6334531f"
		expected := "This is a test document.\n"

		blob, err := gh.GetBlob(ctx, owner, repo, sha)
		if assert.NoError(t, err) {
			assert.Equal(t, expected, *blob)
		}
	})

	t.Run("get an error when non-existent SHA is specified", func(t *testing.T) {
		ctx := context.Background()
		owner := "uu64"
		repo := "non-existent"
		sha := "test"

		_, err := gh.GetBlob(ctx, owner, repo, sha)
		assert.Error(t, err)
	})
}
