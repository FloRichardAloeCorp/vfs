package directory

import (
	"fmt"

	"github.com/FloRichardAloeCorp/vfs/server/internal/features/file"
)

var (
	_ DirectoryFeatures = (*directoryController)(nil)
)

type directoryController struct {
	directoryRepository iDirectoryRepository
}

func newDirectoryController(directoryRepository iDirectoryRepository) *directoryController {
	return &directoryController{
		directoryRepository: directoryRepository,
	}
}

func (c *directoryController) Create(path string) error {
	if err := c.directoryRepository.Create(path); err != nil {
		return fmt.Errorf("can't create directory: %w", err)
	}
	return nil
}

func (c *directoryController) ListFiles(path string) ([]file.FileInfo, error) {
	files, err := c.directoryRepository.ListFiles(path)
	if err != nil {
		return nil, fmt.Errorf("cant' list files: %w", err)
	}
	return files, nil
}

func (c *directoryController) Delete(path string) error {
	if err := c.directoryRepository.Delete(path); err != nil {
		return fmt.Errorf("can't delete directory: %w", err)
	}
	return nil
}
