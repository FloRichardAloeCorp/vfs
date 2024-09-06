package file

import "fmt"

var (
	_ FileFeatures = (*fileController)(nil)
)

type fileController struct {
	fileRepository iFileRepository
}

func newFileController(fileRepository iFileRepository) *fileController {
	return &fileController{
		fileRepository: fileRepository,
	}
}

func (c *fileController) Create(path string, content []byte) error {
	if err := c.fileRepository.CreateFile(path, content); err != nil {
		return fmt.Errorf("can't create new file: %w", err)
	}
	return nil
}

func (c *fileController) Read(path string) ([]byte, error) {
	content, err := c.fileRepository.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("can't read file: %w", err)
	}
	return content, nil
}

func (c *fileController) ReadInfo(path string) (*FileInfo, error) {
	info, err := c.fileRepository.ReadInfo(path)
	if err != nil {
		return nil, fmt.Errorf("can't read file info: %w", err)
	}
	return info, nil
}

func (c *fileController) UpdateName(path string, newName string) error {
	return nil
}

func (c *fileController) Delete(path string) error {
	if err := c.fileRepository.DeleteFile(path); err != nil {
		return fmt.Errorf("can't delete file: %w", err)
	}
	return nil
}
