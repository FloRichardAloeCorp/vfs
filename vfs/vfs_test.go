package vfs

import (
	"path/filepath"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newTestVFS() *VFS {
	return &VFS{
		Root: &Node{
			ID:   uuid.New(),
			Name: "/",
			Type: Directory,
			Children: map[string]*Node{
				"dir1": {
					ID:   uuid.New(),
					Name: "dir1",
					Type: Directory,
					Children: map[string]*Node{
						"file1.txt": {
							ID:      uuid.New(),
							Name:    "file1.txt",
							Type:    File,
							Content: []byte("hello word 1"),
						},
						"file2.txt": {
							ID:      uuid.New(),
							Name:    "file2.txt",
							Type:    File,
							Content: []byte("hello word 2"),
						},
					},
				},
				"dir2": {
					ID:   uuid.New(),
					Name: "dir2",
					Type: Directory,
					Children: map[string]*Node{
						"dir3": {
							ID:   uuid.New(),
							Name: "dir3",
							Type: Directory,
							Children: map[string]*Node{
								"dir4": {
									ID:       uuid.New(),
									Name:     "dir4",
									Type:     Directory,
									Children: map[string]*Node{},
								},
							},
						},
					},
				},
				"file3.txt": {
					ID:      uuid.New(),
					Name:    "file3.txt",
					Type:    File,
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
		expectedNode *Node
	}

	instance := newTestVFS()

	var testCases = [...]testData{
		{
			name:         "Success case",
			shouldFail:   false,
			path:         "/dir1/file2.txt",
			expectedNode: instance.Root.Children["dir1"].Children["file2.txt"],
		},
		{
			name:         "Success case: root",
			shouldFail:   false,
			path:         filepath.Join("/"),
			expectedNode: instance.Root,
		},
		{
			name:         "Success case: File directly under root",
			shouldFail:   false,
			path:         "/file3.txt",
			expectedNode: instance.Root.Children["file3.txt"],
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
			node, err := instance.findNode(testCase.path)
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
		child      *Node
	}

	var testCases = [...]testData{
		{
			name:       "Success case",
			shouldFail: false,
			parentPath: "/dir1",
			child: &Node{
				ID:      uuid.New(),
				Name:    "test.txt",
				Type:    File,
				Content: []byte("hi"),
			},
		},
		{
			name:       "Fail case: invalid parent path",
			shouldFail: true,
			parentPath: filepath.Join("/", "invalid"),
			child: &Node{
				ID:      uuid.New(),
				Name:    "test.txt",
				Type:    File,
				Content: []byte("hi"),
			},
		},
		{
			name:       "Fail case: parent is not a directory",
			shouldFail: true,
			parentPath: "/dir1/file2.txt",
			child: &Node{
				ID:      uuid.New(),
				Name:    "test.txt",
				Type:    File,
				Content: []byte("hi"),
			},
		},
		{
			name:       "Fail case: file already exists",
			shouldFail: true,
			parentPath: "/dir1",
			child: &Node{
				ID:      uuid.New(),
				Name:    "file2.txt",
				Type:    File,
				Content: []byte("hi"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			instance := newTestVFS()
			err := instance.addNode(testCase.parentPath, testCase.child)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				expected, err := instance.findNode(filepath.Join(testCase.parentPath, testCase.child.Name))
				assert.NoError(t, err)
				assert.Equal(t, testCase.child, expected)
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
			err := instance.deleteNode(testCase.path)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				node, err := instance.findNode(testCase.path)
				assert.Error(t, err)
				assert.Nil(t, node)
			}
		})
	}
}
