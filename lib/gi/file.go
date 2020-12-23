package gi

// ContentType indicates the type of Content object.
type ContentType int

const (
	// CtFile indicates that the Content Object is a file.
	CtFile ContentType = iota
	// CtDirectory indicates that the Content Object is a directory.
	CtDirectory
	// CtSymLink indicates that the Content Object is a symbolic link.
	CtSymLink
	// CtSubmodule indicates that the Content Object is a submodule.
	CtSubmodule
)

// Content is the object that represents a file stored in the repository.
type Content struct {
	Type ContentType
	Path string
}
