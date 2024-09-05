package vfs

import (
	"path/filepath"

	"github.com/FloRichardAloeCorp/vfs/vfs/pkg/node"
)

func (vfs *VFS) CreateDirectory(path string) error {
	directoryName := filepath.Base(path)
	directoryNode := node.NewDirectory(directoryName)

	return vfs.engine.AddNode(filepath.Dir(path), directoryNode)
}

func (vfs *VFS) ListFiles(path string) ([]node.Node, error) {
	dir, err := vfs.engine.FindNode(path)
	if err != nil {
		return nil, err
	}

	if dir.Type != node.Directory {
		return nil, ErrFileIsNotADirectory
	}

	nodes := make([]node.Node, 0, len(dir.Children))

	for _, node := range dir.Children {
		nodes = append(nodes, *node)
	}

	return nodes, nil
}

func (vfs *VFS) DeleteDirectory(path string) error {
	directory, err := vfs.engine.FindNode(path)
	if err != nil {
		return err
	}

	if directory.Type != node.Directory {
		return ErrFileIsNotADirectory
	}

	return vfs.engine.DeleteNode(path)
}
