package engine

import (
	"errors"
	"fmt"
)

var (
	ErrFileIsNotADirectory   = errors.New("target file is not a directory")
	ErrFileIsADirectory      = errors.New("target file is not directory")
	ErrUnknowFileOrDirectory = errors.New("unknow file or directory")
	ErrFileAlreadyExists     = errors.New("file already exists")
	ErrTooShorPath           = errors.New("path is too short")
	ErrDelRoot               = errors.New("can't delete root directory")
)

type InvalidFileError struct {
	BaseFileName     string
	ChildrenFileName string
}

func (e *InvalidFileError) Error() string {
	return fmt.Sprintf("file %s doesn't exists in %s directory", e.ChildrenFileName, e.BaseFileName)
}
