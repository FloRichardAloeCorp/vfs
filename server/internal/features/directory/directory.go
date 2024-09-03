package directory

import (
	"fmt"

	"github.com/FloRichardAloeCorp/vfs/server/internal/features/file"
	"github.com/FloRichardAloeCorp/vfs/vfs"
)

type DirectoryFeatures interface {
	Create(path string) error
	ListFiles(path string) ([]file.FileInfo, error)
	Delete(path string) error
}

func NewDirectoryFeatures(connectorType string, connector any) (DirectoryFeatures, error) {
	var directoryRepo iDirectoryRepository
	switch connectorType {
	case "vfs":
		db, castable := connector.(*vfs.VFS)
		if !castable {
			return nil, fmt.Errorf("can't cast connector to vfs")
		}
		directoryRepo = NewDirectoryRepository(db)
	}

	return newDirectoryController(directoryRepo), nil
}
