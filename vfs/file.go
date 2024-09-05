package vfs

import (
	"path/filepath"

	"github.com/FloRichardAloeCorp/vfs/vfs/pkg/node"
)

func (vfs *VFS) CreateFile(path string, content []byte) error {
	fileName := filepath.Base(path)
	fileNode := node.NewFile(fileName, content)

	return vfs.engine.AddNode(filepath.Dir(path), fileNode)
}

func (vfs *VFS) ReadFile(path string) ([]byte, error) {
	file, err := vfs.engine.FindNode(path)
	if err != nil {
		return nil, err
	}

	if file.Type != node.File {
		return nil, ErrFileIsADirectory
	}

	return file.Content, nil
}

func (vfs *VFS) ReadFileInfo(path string) (*node.Node, error) {
	file, err := vfs.engine.FindNode(path)
	if err != nil {
		return nil, err
	}

	if file.Type != node.File {
		return nil, ErrFileIsADirectory
	}

	return file, nil
}

func (vfs *VFS) DeleteFile(path string) error {
	file, err := vfs.engine.FindNode(path)
	if err != nil {
		return err
	}

	if file.Type != node.File {
		return ErrFileIsADirectory
	}

	return vfs.engine.DeleteNode(path)
}
