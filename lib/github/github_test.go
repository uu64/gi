package github

import (
	"context"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uu64/gi/lib/core"
)

var giTestToken = os.Getenv("GI_TEST_TOKEN")

func TestGetTree(t *testing.T) {
	t.Run("can get contents sorted by path (not recursively)", func(t *testing.T) {
		expected := []*core.TreeNode{
			core.NewTreeNode(
				core.NtBlob,
				"README.md",
				"6799dc110eaa15880d42c0b447013c30422dc527"),
			core.NewTreeNode(
				core.NtBlob,
				".gitmodules",
				"1a41f17d2c803f43f316adb1cf07e6b9cbfbda3d"),
			core.NewTreeNode(
				core.NtBlob,
				"LICENSE",
				"c46410aef77ff7fbe2f44815ac5ca8fb351d190f"),
			core.NewTreeNode(
				core.NtBlob,
				"testdocument.txt",
				"d6455f8dde4bd5300e65467b170f7485f4ab77e7"),
			core.NewTreeNode(
				core.NtTree,
				"docs",
				"633ec3e50dafa8ca1d9b87763fd2bf1f90dd77b1"),
			core.NewTreeNode(
				core.NtSubmodule,
				"ghapi-test",
				"017bdc26adbd7c544b2180ab947857ff98d8434f"),
		}

		gh := NewRepository("uu64", "ghapi-test", "main", giTestToken)
		ctx := context.Background()
		recursive := false

		contents, err := gh.GetTree(ctx, recursive)
		sort.Slice(expected, func(i, j int) bool {
			return *expected[i].Path < *expected[j].Path
		})

		if assert.NoError(t, err) {
			for i, content := range contents {
				assert.Equal(t, expected[i].Path, content.Path)
				assert.Equal(t, expected[i].Type, content.Type)
				assert.Equal(t, expected[i].SHA, content.SHA)
			}
		}
	})

	t.Run("can get all contents sorted by the path", func(t *testing.T) {
		expected := []*core.TreeNode{
			core.NewTreeNode(
				core.NtBlob,
				"README.md",
				"6799dc110eaa15880d42c0b447013c30422dc527"),
			core.NewTreeNode(
				core.NtBlob,
				".gitmodules",
				"1a41f17d2c803f43f316adb1cf07e6b9cbfbda3d"),
			core.NewTreeNode(
				core.NtBlob,
				"LICENSE",
				"c46410aef77ff7fbe2f44815ac5ca8fb351d190f"),
			core.NewTreeNode(
				core.NtBlob,
				"testdocument.txt",
				"d6455f8dde4bd5300e65467b170f7485f4ab77e7"),
			core.NewTreeNode(
				core.NtBlob,
				"docs/testdocument.txt",
				"a0f31e800f7bb4493ad94210b9f1770f6334531f"),
			core.NewTreeNode(
				core.NtTree,
				"docs",
				"633ec3e50dafa8ca1d9b87763fd2bf1f90dd77b1"),
			core.NewTreeNode(
				core.NtSubmodule,
				"ghapi-test",
				"017bdc26adbd7c544b2180ab947857ff98d8434f"),
		}

		gh := NewRepository("uu64", "ghapi-test", "main", giTestToken)
		ctx := context.Background()
		recursive := true

		contents, err := gh.GetTree(ctx, recursive)
		sort.Slice(expected, func(i, j int) bool {
			return *expected[i].Path < *expected[j].Path
		})

		if assert.NoError(t, err) {
			for i, content := range contents {
				assert.Equal(t, expected[i].Path, content.Path)
				assert.Equal(t, expected[i].Type, content.Type)
				assert.Equal(t, expected[i].SHA, content.SHA)
			}
		}
	})

	t.Run("can get all contents in feature branch", func(t *testing.T) {
		expected := []*core.TreeNode{
			core.NewTreeNode(
				core.NtBlob,
				"README.md",
				"6799dc110eaa15880d42c0b447013c30422dc527"),
			core.NewTreeNode(
				core.NtBlob,
				".gitmodules",
				"1a41f17d2c803f43f316adb1cf07e6b9cbfbda3d"),
			core.NewTreeNode(
				core.NtBlob,
				"LICENSE",
				"c46410aef77ff7fbe2f44815ac5ca8fb351d190f"),
			core.NewTreeNode(
				core.NtBlob,
				"testdocument.txt",
				"d6455f8dde4bd5300e65467b170f7485f4ab77e7"),
			core.NewTreeNode(
				core.NtBlob,
				"docs/testdocument.txt",
				"802f4ed1f1e64021f0918422ad75a8545e344010"),
			core.NewTreeNode(
				core.NtTree,
				"docs",
				"fdea391afb1439f651e2f24f6523b8fe88c028dd"),
			core.NewTreeNode(
				core.NtSubmodule,
				"ghapi-test",
				"017bdc26adbd7c544b2180ab947857ff98d8434f"),
		}

		gh := NewRepository("uu64", "ghapi-test", "feature/not-main", giTestToken)
		ctx := context.Background()
		recursive := true

		contents, err := gh.GetTree(ctx, recursive)
		sort.Slice(expected, func(i, j int) bool {
			return *expected[i].Path < *expected[j].Path
		})

		if assert.NoError(t, err) {
			for i, content := range contents {
				assert.Equal(t, expected[i].Path, content.Path)
				assert.Equal(t, expected[i].Type, content.Type)
				assert.Equal(t, expected[i].SHA, content.SHA)
			}
		}
	})

	t.Run("get an error when non-existent repository is specified", func(t *testing.T) {
		gh := NewRepository("uu64", "not-exist", "main", giTestToken)
		ctx := context.Background()
		recursive := false

		_, err := gh.GetTree(ctx, recursive)
		assert.Error(t, err)
	})
}

func TestGetBlob(t *testing.T) {
	t.Run("can get a blob", func(t *testing.T) {
		gh := NewRepository("uu64", "ghapi-test", "main", giTestToken)
		ctx := context.Background()
		// SHA of docs/testdocument.txt
		sha := "a0f31e800f7bb4493ad94210b9f1770f6334531f"
		expected := "This is a test document.\n"

		blob, err := gh.GetBlob(ctx, sha)
		if assert.NoError(t, err) {
			assert.Equal(t, []byte(expected), blob)
		}
	})

	t.Run("can get a blob in feature branch", func(t *testing.T) {
		gh := NewRepository("uu64", "ghapi-test", "feature/not-main", giTestToken)
		ctx := context.Background()
		// SHA of docs/testdocument.txt
		sha := "802f4ed1f1e64021f0918422ad75a8545e344010"
		expected := "This is a test document!\n"

		blob, err := gh.GetBlob(ctx, sha)
		if assert.NoError(t, err) {
			assert.Equal(t, []byte(expected), blob)
		}
	})

	t.Run("get an error when non-existent SHA is specified", func(t *testing.T) {
		gh := NewRepository("uu64", "not-exist", "main", giTestToken)
		ctx := context.Background()
		sha := "test"

		_, err := gh.GetBlob(ctx, sha)
		assert.Error(t, err)
	})
}
