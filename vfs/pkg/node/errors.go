package node

import (
	"errors"
)

var (
	ErrFileIsNotADirectory = errors.New("target file is not a directory")
	ErrNoResult            = errors.New("unknown file or directory")
)
