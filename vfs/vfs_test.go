package vfs

import (
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/FloRichardAloeCorp/vfs/vfs/internal/engine"
	"github.com/FloRichardAloeCorp/vfs/vfs/pkg/node"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func newTestVFS() *vfs {
	createdAt := time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)
	return &vfs{
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
}

func TestVfsCreateFile(t *testing.T) {
	type testData struct {
		name       string
		shouldFail bool
		path       string
		content    []byte
	}

	var testCases = [...]testData{
		{
			name:       "Success case",
			shouldFail: false,
			path:       "/dir1/test.txt",
			content:    []byte("Hello word"),
		},
		{
			name:       "Fail case: invalid path",
			shouldFail: true,
			path:       "/invalid/test",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			instance := newTestVFS()
			err := instance.CreateFile(testCase.path, testCase.content)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestVfsCreateDirectory(t *testing.T) {
	type testData struct {
		name       string
		shouldFail bool
		path       string
	}

	var testCases = [...]testData{
		{
			name:       "Succes case",
			shouldFail: false,
			path:       "/dir1/test",
		},
		{
			name:       "Fail case: path already exists",
			shouldFail: true,
			path:       "/dir1",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			instance := newTestVFS()
			err := instance.CreateDirectory(testCase.path)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestVfsReadFileContent(t *testing.T) {
	type testData struct {
		name            string
		shouldFail      bool
		path            string
		expectedContent []byte
	}

	instance := newTestVFS()

	var testCases = [...]testData{
		{
			name:            "Success case",
			shouldFail:      false,
			path:            "/dir1/file2.txt",
			expectedContent: instance.engine.Node.Children["dir1"].Children["file2.txt"].Content,
		},
		{
			name:       "Fail case: invalid path",
			shouldFail: true,
			path:       "/invalid",
		},
		{
			name:       "Fail case: File is directory",
			shouldFail: true,
			path:       "/dir1",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			content, err := instance.ReadFileContent(testCase.path)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedContent, content)
			}
		})
	}
}

func TestVfsReadFileInfo(t *testing.T) {
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
			expectedNode: instance.engine.Node.Children["dir1"].Children["file2.txt"],
		},
		{
			name:       "Fail case: invalid path",
			shouldFail: true,
			path:       "/invalid",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			node, err := instance.ReadFileInfo(testCase.path)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedNode, node)
			}
		})
	}
}

func TestVfsListFiles(t *testing.T) {
	type testData struct {
		name          string
		shouldFail    bool
		path          string
		expectedNodes []node.Node
	}

	instance := newTestVFS()

	var testCases = [...]testData{
		{
			name:       "Succes case",
			shouldFail: false,
			path:       "/dir1",
			expectedNodes: []node.Node{
				*instance.engine.Node.Children["dir1"].Children["file1.txt"],
				*instance.engine.Node.Children["dir1"].Children["file2.txt"],
			},
		},
		{
			name:       "Fail case: invalid path",
			shouldFail: true,
			path:       filepath.Join("/", "invalid"),
		},
		{
			name:       "Fail case: file is not a directory",
			shouldFail: true,
			path:       "/dir1/file1.txt",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			nodes, err := instance.ListChilren(testCase.path)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, testCase.expectedNodes, nodes)
			}
		})
	}
}

func TestVfsDeleteFile(t *testing.T) {
	type testData struct {
		name       string
		shouldFail bool
		path       string
	}

	var testCases = [...]testData{
		{
			name:       "Success case",
			shouldFail: false,
			path:       "/dir1/file1.txt",
		},
		{
			name:       "Fail case: invalid path",
			shouldFail: true,
			path:       "/invalid",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			instance := newTestVFS()
			err := instance.DeleteFile(testCase.path)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestVfsRenameFile(t *testing.T) {
	type testData struct {
		name       string
		shouldFail bool
		path       string
		newName    string
	}

	var testCases = [...]testData{
		{
			name:       "Success case",
			shouldFail: false,
			path:       "/dir1/file1.txt",
			newName:    "modified.txt",
		},
		{
			name:       "Fail case: empty new name",
			shouldFail: true,
			path:       "/dir1/file1.txt",
			newName:    "",
		},
		{
			name:       "Fail case: invalid path",
			shouldFail: true,
			path:       "/invalid",
			newName:    "modified.txt",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			instance := newTestVFS()
			err := instance.RenameFile(testCase.path, testCase.newName)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				updatedFile, err := instance.engine.FindNode(strings.Replace(testCase.path, filepath.Base(testCase.path), testCase.newName, 1))
				assert.NoError(t, err)
				assert.Equal(t, testCase.newName, updatedFile.Name)
			}
		})
	}
}
