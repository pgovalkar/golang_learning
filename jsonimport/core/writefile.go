package core

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

// Package to Write files , contains overriding methods that have different function signatures
func WriteFile(file_name string, data []byte, outputpath string) error {
	if err := os.MkdirAll(outputpath, 0755); err != nil {
		return errors.WithMessagef(err, "directory creation failed %s", outputpath)
	}
	if err := os.WriteFile(filepath.Join(outputpath, file_name), data, 0644); err != nil {
		return errors.WithMessagef(err, "error writing file %q", file_name)
	}
	return nil
}

// func ensureDir(dirName string) error {
// 	err := os.MkdirAll(dirName, os.ModePerm)
// 	if err == nil {
// 		return nil
// 	}

// 	err = os.Chmod(dirName, 0755)
// 	if err != nil {
// 		fmt.Printf("Error changing directory permission: %s\n", err.Error())
// 		return err
// 	}

// 	if os.IsExist(err) {
// 		// check that the existing path is a directory
// 		info, err := os.Stat(dirName)
// 		if err != nil {
// 			return err
// 		}
// 		if !info.IsDir() {
// 			return errors.New("path exists but is not a directory")
// 		}
// 		return nil
// 	}
// 	return err
// }
