package engine

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/FloRichardAloeCorp/vfs/vfs/pkg/node"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newTestVFS() *Engine {
	createdAt := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	return &Engine{
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
	}
}

func TestVfsFindNode(t *testing.T) {
	type testData struct {
		name         string
		shouldFail   bool
		path         string
		expectedNode *node.Node
	}

	instance := newTestVFS()

	var testCases = [...]testData{
		{
			name:         "Success case",
			shouldFail:   false,
			path:         "/dir1/file2.txt",
			expectedNode: instance.Node.Children["dir1"].Children["file2.txt"],
		},
		{
			name:         "Success case: root",
			shouldFail:   false,
			path:         filepath.Join("/"),
			expectedNode: instance.Node,
		},
		{
			name:         "Success case: File directly under root",
			shouldFail:   false,
			path:         "/file3.txt",
			expectedNode: instance.Node.Children["file3.txt"],
		},
		{
			name:       "Fail case: no path provided",
			shouldFail: true,
			path:       "",
		},
		{
			name:       "Fail case: invalid path at root",
			shouldFail: true,
			path:       filepath.Join("/", "invalid"),
		},
		{
			name:       "Fail case: invalid path",
			shouldFail: true,
			path:       "/dir1/invalid",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			node, err := instance.FindNode(testCase.path)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedNode, node)
			}
		})
	}
}

func TestVfsAddNode(t *testing.T) {
	type testData struct {
		name       string
		shouldFail bool
		parentPath string
		child      *node.Node
	}

	var testCases = [...]testData{
		{
			name:       "Success case",
			shouldFail: false,
			parentPath: "/dir1",
			child: &node.Node{
				ID:      uuid.New(),
				Name:    "test.txt",
				Type:    node.File,
				Content: []byte("hi"),
			},
		},
		{
			name:       "Fail case: invalid parent path",
			shouldFail: true,
			parentPath: filepath.Join("/", "invalid"),
			child: &node.Node{
				ID:      uuid.New(),
				Name:    "test.txt",
				Type:    node.File,
				Content: []byte("hi"),
			},
		},
		{
			name:       "Fail case: parent is not a directory",
			shouldFail: true,
			parentPath: "/dir1/file2.txt",
			child: &node.Node{
				ID:      uuid.New(),
				Name:    "test.txt",
				Type:    node.File,
				Content: []byte("hi"),
			},
		},
		{
			name:       "Fail case: file already exists",
			shouldFail: true,
			parentPath: "/dir1",
			child: &node.Node{
				ID:      uuid.New(),
				Name:    "file2.txt",
				Type:    node.File,
				Content: []byte("hi"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			instance := newTestVFS()
			err := instance.AddNode(testCase.parentPath, testCase.child)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				expected, err := instance.FindNode(filepath.Join(testCase.parentPath, testCase.child.Name))
				assert.NoError(t, err)
				assert.Equal(t, testCase.child, expected)
				assert.NotEqual(t, instance.Node.LastUpdate, newTestVFS().Node.LastUpdate)
			}
		})
	}
}

func TestVfsDeleteNode(t *testing.T) {
	type testData struct {
		name       string
		shouldFail bool
		path       string
	}

	var testCases = [...]testData{
		{
			name:       "Succes case",
			shouldFail: false,
			path:       "/dir1/file1.txt",
		},
		{
			name:       "Succes case: first level element",
			shouldFail: false,
			path:       "/dir1",
		},
		{
			name:       "Succes case: longer path",
			shouldFail: false,
			path:       "/dir2/dir3/dir4",
		},
		{
			name:       "Fail case: deleting root",
			shouldFail: true,
			path:       "/",
		},
		{
			name:       "Fail case: no path provided",
			shouldFail: true,
			path:       "",
		},
		{
			name:       "Fail case: invalid path at root",
			shouldFail: true,
			path:       filepath.Join("/", "invalid"),
		},
		{
			name:       "Fail case: invalid path",
			shouldFail: true,
			path:       "/dir1/invalid",
		},
		{
			name:       "Succes case: invalid longer path",
			shouldFail: true,
			path:       "/dir2/invalid/invalid",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			instance := newTestVFS()
			err := instance.DeleteNode(testCase.path)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				node, err := instance.FindNode(testCase.path)
				assert.Error(t, err)
				assert.Nil(t, node)
			}
		})
	}
}

func TestVfsUpdateAllSuccessCases(t *testing.T) {
	updateFn := func(node *node.Node) {
		node.CreatedAt = time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	}

	instance := newTestVFS()
	err := instance.UpdateAll("/dir1/file1.txt", updateFn)
	assert.NoError(t, err)
	assert.Equal(t, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), instance.Node.CreatedAt)
	assert.Equal(t, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), instance.Node.Children["dir1"].CreatedAt)
	assert.Equal(t, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), instance.Node.Children["dir1"].Children["file1.txt"].CreatedAt)

	instance = newTestVFS()
	err = instance.UpdateAll("/", updateFn)
	assert.NoError(t, err)
	assert.Equal(t, time.Date(0, 0, 0, 0, 0, 0, 0, time.Local), instance.Node.CreatedAt)
}

func TestVfsUpdateAllFailCases(t *testing.T) {
	type testData struct {
		name     string
		path     string
		updateFn func(node *node.Node)
	}

	var testCases = [...]testData{
		{
			name: "Fail case: invalid path",
			path: "/dir1/invalid",
			updateFn: func(node *node.Node) {
				node.CreatedAt = time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
			},
		},
		{
			name: "Fail case: no path",
			path: "",
			updateFn: func(node *node.Node) {
				node.CreatedAt = time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			instance := newTestVFS()
			err := instance.UpdateAll(testCase.path, testCase.updateFn)
			assert.Error(t, err)
		})
	}
}
