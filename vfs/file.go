package vfs

import (
	"path/filepath"
)

func (vfs *VFS) CreateFile(path string, content []byte) error {
	fileName := filepath.Base(path)
	fileNode := NewFileNode(fileName, content)

	return vfs.addNode(filepath.Dir(path), fileNode)
}

func (vfs *VFS) ReadFile(path string) ([]byte, error) {
	node, err := vfs.findNode(path)
	if err != nil {
		return nil, err
	}

	if node.Type != File {
		return nil, ErrFileIsADirectory
	}

	return node.Content, nil
}

func (vfs *VFS) ReadFileInfo(path string) (*Node, error) {
	node, err := vfs.findNode(path)
	if err != nil {
		return nil, err
	}

	if node.Type != File {
		return nil, ErrFileIsADirectory
	}

	return node, nil
}
