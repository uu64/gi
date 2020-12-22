package gi

type ContentType int

const (
	File ContentType = iota
	Directory
	SymLink
	Submodule
)

// Content is the object that represents a file stored in the repository.
type Content struct {
	Type ContentType
	Path string
}
