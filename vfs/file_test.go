package vfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestVfsReadFile(t *testing.T) {
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
			expectedContent: instance.Root.Children["dir1"].Children["file2.txt"].Content,
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
			content, err := instance.ReadFile(testCase.path)
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
		{
			name:       "Fail case: file is a directory",
			shouldFail: true,
			path:       "/dir1",
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
