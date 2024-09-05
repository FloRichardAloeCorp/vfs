package node

type NodeType int

const (
	File NodeType = iota
	Directory
)

func (t *NodeType) String() string {
	switch *t {
	case File:
		return "file"
	case Directory:
		return "directory"
	default:
		return "unknown"
	}
}
