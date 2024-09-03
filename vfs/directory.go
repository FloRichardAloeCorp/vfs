package vfs

import "path/filepath"

func (vfs *VFS) CreateDirectory(path string) error {
	directoryName := filepath.Base(path)
	directoryNode := NewDirectoryNode(directoryName)

	return vfs.addNode(filepath.Dir(path), directoryNode)
}

func (vfs *VFS) ListFiles(path string) ([]Node, error) {
	dir, err := vfs.findNode(path)
	if err != nil {
		return nil, err
	}

	if dir.Type != Directory {
		return nil, ErrFileIsNotADirectory
	}

	nodes := make([]Node, 0, len(dir.Children))

	for _, node := range dir.Children {
		nodes = append(nodes, *node)
	}

	return nodes, nil
}
