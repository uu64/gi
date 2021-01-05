package core

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var gi *Gi

type fileMock struct {
	nt      NodeType
	path    string
	sha     string
	content string
}

var testFiles = []fileMock{
	{
		nt:      NtBlob,
		path:    ".gitignore",
		sha:     "abc",
		content: "This is .gitignore",
	},
	{
		nt:      NtBlob,
		path:    "golang.gitignore",
		sha:     "def",
		content: "This is golang.gitignore",
	},
	{
		nt:      NtBlob,
		path:    "javascript.gitignore",
		sha:     "fhi",
		content: "This is javascript.gitignore",
	},
	{
		nt:      NtBlob,
		path:    "shellscript.gitignore",
		sha:     "jkl",
		content: "This is shellscript.gitignore",
	},
}

type repoMock struct{}

func (m *repoMock) GetTree(ctx context.Context, recursive bool) ([]*TreeNode, error) {
	tree := []*TreeNode{}
	for _, file := range testFiles {
		tree = append(tree, NewTreeNode(file.nt, file.path, file.sha))
	}
	return tree, nil
}

func (m *repoMock) GetBlob(ctx context.Context, sha string) ([]byte, error) {
	for _, file := range testFiles {
		if file.sha == sha {
			return []byte(file.content), nil
		}
	}
	return nil, errors.New("failed to get content")
}

func init() {
	gi = NewGi(&repoMock{})
}

func TestListGitIgnorePath(t *testing.T) {
	t.Run("can only get gitignore template files", func(t *testing.T) {
		paths, _ := gi.ListGitIgnorePath()
		expected := []string{
			testFiles[1].path,
			testFiles[2].path,
			testFiles[3].path,
		}
		assert.Equal(t, expected, paths)
	})
}

func TestDownload(t *testing.T) {
	t.Run("can download a file", func(t *testing.T) {
		selected := []string{
			testFiles[1].path,
		}
		buffer := new(bytes.Buffer)
		gi.Download(selected, buffer)

		expected := fmt.Sprintf("# %s\n%s\n", testFiles[1].path, testFiles[1].content)
		assert.Equal(t, expected, buffer.String())
	})

	t.Run("can download multiple files", func(t *testing.T) {
		selected := []string{
			testFiles[1].path,
			testFiles[2].path,
		}
		buffer := new(bytes.Buffer)
		gi.Download(selected, buffer)

		expected := fmt.Sprintf("# %s\n%s\n# %s\n%s\n",
			testFiles[1].path, testFiles[1].content,
			testFiles[2].path, testFiles[2].content)
		assert.Equal(t, expected, buffer.String())
	})
}
