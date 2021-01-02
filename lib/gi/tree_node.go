package gi

// NodeType indicates the type of TreeNode object.
type NodeType int

const (
	// NtBlob indicates that the Content Object is a blob.
	NtBlob NodeType = iota
	// NtTree indicates that the Content Object is a tree.
	NtTree
	// NtSubmodule indicates that the Content Object is a submodule.
	NtSubmodule
)

// TreeNode is the object that represents a content stored in the repository.
// NOTE: https://git-scm.com/book/en/v2/Git-Internals-Git-Objects/#_tree_objects
type TreeNode struct {
	Type NodeType
	Path *string
	SHA  *string
}

// NewTreeNode returns a new TreeNode object.
func NewTreeNode(nodeType NodeType, path string, sha string) *TreeNode {
	return &TreeNode{
		Type: nodeType,
		Path: &path,
		SHA:  &sha,
	}
}
