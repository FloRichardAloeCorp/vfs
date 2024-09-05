package vfs

import (
	"testing"
	"time"

	"github.com/FloRichardAloeCorp/vfs/vfs/internal/engine"
	"github.com/FloRichardAloeCorp/vfs/vfs/pkg/node"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newTestVFS() *VFS {
	createdAt := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	return &VFS{
		engine: &engine.Engine{
			Node: &node.Node{
				ID:        uuid.New(),
				Name:      "/",
				Type:      node.Directory,
				CreatedAt: createdAt,
				Children: map[string]*node.Node{
					"dir1": {
						ID:        uuid.New(),
						Name:      "dir1",
						Type:      node.Directory,
						CreatedAt: createdAt,
						Children: map[string]*node.Node{
							"file1.txt": {
								ID:        uuid.New(),
								Name:      "file1.txt",
								Type:      node.File,
								CreatedAt: createdAt,
								Content:   []byte("hello word 1"),
							},
							"file2.txt": {
								ID:        uuid.New(),
								Name:      "file2.txt",
								Type:      node.File,
								CreatedAt: createdAt,
								Content:   []byte("hello word 2"),
							},
						},
					},
					"dir2": {
						ID:        uuid.New(),
						Name:      "dir2",
						Type:      node.Directory,
						CreatedAt: createdAt,
						Children: map[string]*node.Node{
							"dir3": {
								ID:        uuid.New(),
								Name:      "dir3",
								Type:      node.Directory,
								CreatedAt: createdAt,
								Children: map[string]*node.Node{
									"dir4": {
										ID:        uuid.New(),
										Name:      "dir4",
										Type:      node.Directory,
										CreatedAt: createdAt,
										Children:  map[string]*node.Node{},
									},
								},
							},
						},
					},
					"file3.txt": {
						ID:      uuid.New(),
						Name:    "file3.txt",
						Type:    node.File,
						Content: []byte("hello word 3"),
					},
				},
			},
		},
	}
}

func TestNew(t *testing.T) {
	vfs := New()
	assert.NotNil(t, vfs)
	assert.NotNil(t, vfs.engine)
	assert.Equal(t, "/", vfs.engine.Name)
	assert.Equal(t, node.Directory, vfs.engine.Type)
}
