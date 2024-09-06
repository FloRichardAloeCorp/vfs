package file

import (
	"fmt"

	"github.com/FloRichardAloeCorp/vfs/vfs"
)

type FileFeatures interface {
	Create(path string, content []byte) error
	Read(path string) ([]byte, error)
	ReadInfo(path string) (*FileInfo, error)
	UpdateName(path string, newName string) error
	Delete(path string) error
}

func NewFileFeatures(connectorType string, connector any) (FileFeatures, error) {
	var fileRepo iFileRepository
	switch connectorType {
	case "vfs":
		db, castable := connector.(vfs.VFS)
		if !castable {
			return nil, fmt.Errorf("can't cast connector to vfs")
		}
		fileRepo = NewFileRepository(db)
	}

	return newFileController(fileRepo), nil
}
