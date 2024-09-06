package file

import (
	"fmt"

	"github.com/FloRichardAloeCorp/vfs/vfs"
)

var (
	_ iFileRepository = (*fileRepository)(nil)
)

type iFileRepository interface {
	CreateFile(path string, content []byte) error
	ReadFile(path string) ([]byte, error)
	ReadInfo(path string) (*FileInfo, error)
	UpdateName(path string, newName string) error
	DeleteFile(path string) error
}

type fileRepository struct {
	vfs vfs.VFS
}

func NewFileRepository(vfs vfs.VFS) iFileRepository {
	return &fileRepository{
		vfs: vfs,
	}
}

func (r *fileRepository) CreateFile(path string, content []byte) error {
	if err := r.vfs.CreateFile(path, content); err != nil {
		return fmt.Errorf("can't create file: %w", err)
	}
	return nil
}

func (r *fileRepository) ReadFile(path string) ([]byte, error) {
	content, err := r.vfs.ReadFileContent(path)
	if err != nil {
		return nil, fmt.Errorf("can't read file content: %w", err)
	}
	return content, nil
}

func (r *fileRepository) ReadInfo(path string) (*FileInfo, error) {
	node, err := r.vfs.ReadFileInfo(path)
	if err != nil {
		return nil, fmt.Errorf("can't read file info: %w", err)
	}

	info := &FileInfo{
		ID:         node.ID,
		Name:       node.Name,
		Path:       path,
		Type:       node.Type.String(),
		CreatedAt:  node.CreatedAt,
		LastUpdate: node.LastUpdate,
	}

	return info, nil
}

func (r *fileRepository) UpdateName(path string, newName string) error {
	err := r.vfs.RenameFile(path, newName)
	if err != nil {
		return fmt.Errorf("can't rename file: %w", err)
	}
	return nil
}

func (r *fileRepository) DeleteFile(path string) error {
	err := r.vfs.DeleteFile(path)
	if err != nil {
		return fmt.Errorf("can't delete file: %w", err)
	}
	return nil
}
