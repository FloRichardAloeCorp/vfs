package vfs

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type VFS struct {
	*Node
}

func NewVFS() *VFS {
	return &VFS{
		&Node{
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

	updateLastUpdateDate := func(node *Node) {
		node.LastUpdate = time.Now().UTC()
	}

	if err := vfs.updateAll(parentPath, updateLastUpdateDate); err != nil {
		return fmt.Errorf("can't update last update date of parent nodes: %w", err)
	}

	return nil
}

func (vfs *VFS) findNode(path string) (*Node, error) {
	var node = vfs.Node
	var err error

	if path == "/" {
		return node, nil
	}

	pathParts := strings.Split(path, "/")
	if len(path) == 0 {
		return nil, ErrTooShorPath
	}

	for _, part := range pathParts[1:] {
		node, err = node.FindChild(part)
		if err != nil {
			return nil, err
		}
	}

	return node, nil
}

func (vfs *VFS) deleteNode(path string) error {
	if path == "/" {
		return ErrDelRoot
	}

	parentPath := filepath.Dir(path)
	parent, err := vfs.findNode(parentPath)
	if err != nil {
		return err
	}

	if _, ok := parent.Children[filepath.Base(path)]; !ok {
		return ErrUnknowFileOrDirectory
	}

	delete(parent.Children, filepath.Base(path))
	return nil
}

func (vfs *VFS) updateAll(path string, updateFn func(node *Node)) error {
	if path == "/" {
		updateFn(vfs.Node)
		return nil
	}

	// Check if path is valid before applying updates
	_, err := vfs.findNode(path)
	if err != nil {
		return err
	}

	// Apply updates on all nodes of the path
	pathParts := strings.Split(path, "/")
	if len(path) == 0 {
		return ErrTooShorPath
	}

	node := vfs.Node
	updateFn(node)

	for _, part := range pathParts[1:] {
		node, err = node.FindChild(part)
		if err != nil {
			return err
		}
		updateFn(node)
	}

	return nil
}
