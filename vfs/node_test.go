package vfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeTypeString(t *testing.T) {
	type testData struct {
		name        string
		nt          NodeType
		expectedRes string
	}

	var testCases = [...]testData{
		{
			name:        "Success case: file",
			nt:          File,
			expectedRes: "file",
		},
		{
			name:        "Success case: directory",
			nt:          Directory,
			expectedRes: "directory",
		},
		{
			name:        "Success case: unknown",
			nt:          99999,
			expectedRes: "unknown",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.nt.String()
			assert.Equal(t, testCase.expectedRes, res)
		})
	}
}

func TestNewFileNode(t *testing.T) {
	type testData struct {
		name         string
		fileName     string
		content      []byte
		expectedNode *Node
	}

	var testCases = [...]testData{
		{
			name:     "Succes case",
			fileName: "test",
			content:  []byte("hello word"),
			expectedNode: &Node{
				Name:     "test",
				Content:  []byte("hello word"),
				Children: nil,
				Type:     File,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			node := NewFileNode(testCase.fileName, testCase.content)
			assert.Equal(t, testCase.expectedNode.Name, node.Name)
			assert.Equal(t, testCase.expectedNode.Type, node.Type)
			assert.Equal(t, testCase.expectedNode.Content, node.Content)
			assert.Equal(t, testCase.expectedNode.Children, node.Children)
		})
	}
}

func TestNewDirectoryNode(t *testing.T) {
	type testData struct {
		name         string
		dirName      string
		expectedNode *Node
	}

	var testCases = [...]testData{
		{
			name:    "Success case",
			dirName: "test",
			expectedNode: &Node{
				Name:     "test",
				Children: make(map[string]*Node),
				Type:     Directory,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			node := NewDirectoryNode(testCase.dirName)
			assert.Equal(t, testCase.expectedNode.Name, node.Name)
			assert.Equal(t, testCase.expectedNode.Type, node.Type)
			assert.Equal(t, testCase.expectedNode.Content, node.Content)
			assert.Equal(t, testCase.expectedNode.Children, node.Children)
		})
	}
}

func TestNodeFindChild(t *testing.T) {
	type testData struct {
		name         string
		instance     *Node
		shouldFail   bool
		childName    string
		expectedNode *Node
	}

	instance := newTestVFS().Root

	var testCases = [...]testData{
		{
			name:         "Succes case",
			shouldFail:   false,
			instance:     instance,
			childName:    "dir1",
			expectedNode: instance.Children["dir1"],
		},
		{
			name:       "Fail case: invalid name",
			shouldFail: true,
			instance:   instance,
			childName:  "invalid",
		},
		{
			name:       "Fail case: instance is not a directory",
			shouldFail: true,
			instance: &Node{
				Type: File,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			node, err := testCase.instance.FindChild(testCase.childName)
			if testCase.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.expectedNode, node)
			}
		})
	}
}
