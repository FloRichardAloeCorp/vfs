package engine

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/FloRichardAloeCorp/vfs/vfs/pkg/node"
)

type Engine struct {
	*node.Node
}

func (engine *Engine) FindNode(path string) (*node.Node, error) {
	var node = engine.Node
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

func (engine *Engine) AddNode(parentPath string, child *node.Node) error {
	parent, err := engine.FindNode(parentPath)
	if err != nil {
		return err
	}

	if parent.Type != node.Directory {
		return ErrFileIsNotADirectory
	}

	if _, ok := parent.Children[child.Name]; ok {
		return ErrFileAlreadyExists
	}

	parent.Children[child.Name] = child

	updateLastUpdateDate := func(node *node.Node) {
		node.LastUpdate = time.Now().UTC()
	}

	if err := engine.UpdateAll(parentPath, updateLastUpdateDate); err != nil {
		return fmt.Errorf("can't update last update date of parent nodes: %w", err)
	}

	return nil
}

func (engine *Engine) DeleteNode(path string) error {
	if path == "/" {
		return ErrDelRoot
	}

	parentPath := filepath.Dir(path)
	parent, err := engine.FindNode(parentPath)
	if err != nil {
		return err
	}

	if _, ok := parent.Children[filepath.Base(path)]; !ok {
		return ErrUnknowFileOrDirectory
	}

	delete(parent.Children, filepath.Base(path))
	return nil
}

func (engine *Engine) UpdateAll(path string, updateFn func(node *node.Node)) error {
	if path == "/" {
		updateFn(engine.Node)
		return nil
	}

	// Check if path is valid before applying updates
	_, err := engine.FindNode(path)
	if err != nil {
		return err
	}

	// Apply updates on all nodes of the path
	pathParts := strings.Split(path, "/")
	if len(path) == 0 {
		return ErrTooShorPath
	}

	node := engine.Node
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

func (engine *Engine) UpdateOne(path string, updateFn func(node *node.Node)) error {
	node, err := engine.FindNode(path)
	if err != nil {
		return err
	}
	updateFn(node)

	return nil
}
