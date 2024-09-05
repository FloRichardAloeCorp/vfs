package vfs

import (
	"path/filepath"

	"github.com/FloRichardAloeCorp/vfs/vfs/internal/engine"
	"github.com/FloRichardAloeCorp/vfs/vfs/pkg/node"
)

var _ VFS = (*vfs)(nil)

type VFS interface {
	CreateFile(path string, content []byte) error
	CreateDirectory(path string) error

	ReadFileContent(path string) ([]byte, error)
	ReadFileInfo(path string) (*node.Node, error)
	ListChilren(path string) ([]node.Node, error)

	DeleteFile(path string) error
}

type vfs struct {
	engine *engine.Engine
}

func New() VFS {
	return &vfs{
		&engine.Engine{
			Node: node.NewDirectory("/"),
		},
	}
}

func (vfs *vfs) CreateFile(path string, content []byte) error {
	fileName := filepath.Base(path)
	fileNode := node.NewFile(fileName, content)

	return vfs.engine.AddNode(filepath.Dir(path), fileNode)
}

func (vfs *vfs) CreateDirectory(path string) error {
	directoryName := filepath.Base(path)
	directoryNode := node.NewDirectory(directoryName)

	return vfs.engine.AddNode(filepath.Dir(path), directoryNode)
}

func (vfs *vfs) ReadFileContent(path string) ([]byte, error) {
	file, err := vfs.engine.FindNode(path)
	if err != nil {
		return nil, err
	}

	if file.Type != node.File {
		return nil, ErrFileIsADirectory
	}

	return file.Content, nil
}

func (vfs *vfs) ReadFileInfo(path string) (*node.Node, error) {
	file, err := vfs.engine.FindNode(path)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (vfs *vfs) ListChilren(path string) ([]node.Node, error) {
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

func (vfs *vfs) DeleteFile(path string) error {
	return vfs.engine.DeleteNode(path)
}
