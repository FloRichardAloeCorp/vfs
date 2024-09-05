package file

import "github.com/google/uuid"

type FileInfo struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Path string    `json:"path"`
	Type string    `json:"type"`
}
