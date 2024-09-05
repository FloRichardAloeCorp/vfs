package directory

import (
	"fmt"
	"path/filepath"

	"github.com/FloRichardAloeCorp/vfs/server/internal/features/file"
	"github.com/FloRichardAloeCorp/vfs/vfs"
)

var (
	_ iDirectoryRepository = (*directoryRepository)(nil)
)

type iDirectoryRepository interface {
	Create(path string) error
	ListFiles(path string) ([]file.FileInfo, error)
	Delete(path string) error
}

type directoryRepository struct {
	vfs *vfs.VFS
}

func NewDirectoryRepository(vfs *vfs.VFS) iDirectoryRepository {
	return &directoryRepository{
		vfs: vfs,
	}
}
func (r *directoryRepository) Create(path string) error {
	if err := r.vfs.CreateDirectory(path); err != nil {
		return fmt.Errorf("can't create repository: %w", err)
	}
	return nil
}

func (r *directoryRepository) ListFiles(path string) ([]file.FileInfo, error) {
	nodes, err := r.vfs.ListFiles(path)
	if err != nil {
		return nil, fmt.Errorf("can't list files in directory: %w", err)
	}

	files := make([]file.FileInfo, 0, len(nodes))
	for _, node := range nodes {
		files = append(files, file.FileInfo{
			ID:         node.ID,
			Name:       node.Name,
			Path:       filepath.Join(path, node.Name),
			Type:       node.Type.String(),
			CreatedAt:  node.CreatedAt,
			LastUpdate: node.LastUpdate,
		})
	}

	return files, nil
}

func (r *directoryRepository) Delete(path string) error {
	err := r.vfs.DeleteDirectory(path)
	if err != nil {
		return fmt.Errorf("can't delete directory: %w", err)
	}
	return nil
}
