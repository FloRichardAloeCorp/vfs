package vfs

import (
	"errors"
)

var (
	ErrFileIsNotADirectory = errors.New("target file is not a directory")
	ErrFileIsADirectory    = errors.New("target file is not directory")
	ErrEmptyName           = errors.New("name is empty")
)
