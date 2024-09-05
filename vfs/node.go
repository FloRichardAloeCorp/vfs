package vfs

import (
	"time"

	"github.com/google/uuid"
)

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

type Node struct {
	ID         uuid.UUID
	Name       string
	Type       NodeType
	Content    []byte
	Children   map[string]*Node
	CreatedAt  time.Time
	LastUpdate time.Time
}

func NewFileNode(name string, content []byte) *Node {
	now := time.Now().UTC()
	return &Node{
		ID:         uuid.New(),
		Name:       name,
		Type:       File,
		Content:    content,
		CreatedAt:  now,
		LastUpdate: now,
	}
}

func NewDirectoryNode(name string) *Node {
	now := time.Now().UTC()
	return &Node{
		ID:         uuid.New(),
		Name:       name,
		Type:       Directory,
		Children:   make(map[string]*Node),
		CreatedAt:  now,
		LastUpdate: now,
	}
}

func (n *Node) FindChild(name string) (*Node, error) {
	if n.Type != Directory {
		return nil, ErrFileIsNotADirectory
	}

	if _, ok := n.Children[name]; !ok {
		return nil, &InvalidFileError{BaseFileName: n.Name, ChildrenFileName: name}
	}

	return n.Children[name], nil
}
