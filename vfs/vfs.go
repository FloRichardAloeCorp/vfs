package vfs

import (
	"github.com/FloRichardAloeCorp/vfs/vfs/internal/engine"
	"github.com/FloRichardAloeCorp/vfs/vfs/pkg/node"
)

type VFS struct {
	engine *engine.Engine
}

func New() *VFS {
	return &VFS{
		&engine.Engine{
			Node: node.NewDirectory("/"),
		},
	}
}
