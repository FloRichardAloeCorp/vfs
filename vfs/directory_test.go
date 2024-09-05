package vfs

import (
	"path/filepath"
	"testing"

	"github.com/FloRichardAloeCorp/vfs/vfs/pkg/node"
	"github.com/stretchr/testify/assert"
)

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

			nodes, err := instance.ListFiles(testCase.path)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.ElementsMatch(t, testCase.expectedNodes, nodes)
			}
		})
	}
}

func TestVfsDeleteDirectory(t *testing.T) {
	type testData struct {
		name       string
		shouldFail bool
		path       string
	}

	var testCases = [...]testData{
		{
			name:       "Success case",
			shouldFail: false,
			path:       "/dir1",
		},
		{
			name:       "Fail case: invalid path",
			shouldFail: true,
			path:       "/invalid",
		},
		{
			name:       "Fail case: directory is afile",
			shouldFail: true,
			path:       "/dir1/file1.txt",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			instance := newTestVFS()
			err := instance.DeleteDirectory(testCase.path)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
