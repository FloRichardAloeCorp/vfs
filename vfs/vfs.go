package vfs

import (
	"strings"

	"github.com/google/uuid"
)

type VFS struct {
	Root *Node
}

func NewVFS() *VFS {
	return &VFS{
		Root: &Node{
			ID:       uuid.New(),
			Name:     "/",
			Type:     Directory,
			Children: make(map[string]*Node),
		},
	}
}

func (vfs *VFS) addNode(parentPath string, child *Node) error {
	parent, err := vfs.findNode(parentPath)
	if err != nil {
		return err
	}

	if parent.Type != Directory {
		return ErrFileIsNotADirectory
	}

	if _, ok := parent.Children[child.Name]; ok {
		return ErrFileAlreadyExists
	}

	parent.Children[child.Name] = child

	return nil
}

func (vfs *VFS) findNode(path string) (*Node, error) {
	if path == "/" {
		return vfs.Root, nil
	}

	pathParts := strings.Split(path, "/")
	if len(path) == 0 {
		return nil, ErrTooShorPath
	}

	// exluding root node
	node, err := vfs.Root.FindChild(pathParts[1])
	if err != nil {
		return nil, err
	}

	if node.Type == File {
		return node, nil
	}

	for _, part := range pathParts[2:] {
		node, err = node.FindChild(part)
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}
