package file

import (
	"time"

	"github.com/google/uuid"
)

type FileInfo struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	LastUpdate time.Time `json:"last_update"`
}
