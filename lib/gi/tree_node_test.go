package gi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTreeNode(t *testing.T) {
	t.Run("can get a new TreeNode object", func(t *testing.T) {
		nodeType := NtBlob
		path := "testfile.txt"

		node := NewTreeNode(nodeType, path)
		assert.Equal(t, nodeType, node.Type)
		assert.Equal(t, path, *node.Path)
	})
}
