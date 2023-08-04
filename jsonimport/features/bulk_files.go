package features

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func Bulk(path string, outputpath string) error {
	items, _ := os.ReadDir(path)
	for _, item := range items {
		if item.IsDir() {
			return fmt.Errorf("path: %s is a directory, Please just add files", path)
		} else {
			filepath, err := filepath.Abs(path)
			if err != nil {
				errors.WithMessagef(err, "failed to get absolute path")
			}
			filename := filepath + "/" + item.Name()
			fmt.Println("Processed file: " + filename)
			File(filename, outputpath)
		}
	}
	return nil
}
