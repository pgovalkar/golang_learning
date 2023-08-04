package core

import (
	"fmt"
	"os"
)

// Package to Read files , contains overriding methods that have different function signatures
// TO Loca;l files do not need token
func ReadFile(fileName string) ([]byte, error) {
	fi, err := os.Stat(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w. Please provide relative path", err)
	}

	if fi.IsDir() {
		return nil, fmt.Errorf("file path %s is a directory", fileName)
	}

	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", fileName)
	} else if err != nil {
		return nil, fmt.Errorf("error checking file %s: %w", fileName, err)
	}

	checkSize := fi.Size()
	if checkSize == 0 {
		return nil, fmt.Errorf("file %s is empty", fileName)
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", fileName, err)
	}

	return file, nil
}
