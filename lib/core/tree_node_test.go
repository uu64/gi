package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTreeNode(t *testing.T) {
	t.Run("can get a new TreeNode object", func(t *testing.T) {
		nodeType := NtBlob
		path := "testfile.txt"
		sha := "0123456789"

		node := NewTreeNode(nodeType, path, sha)
		assert.Equal(t, nodeType, node.Type)
		assert.Equal(t, path, *node.Path)
		assert.Equal(t, sha, *node.SHA)
	})
}
