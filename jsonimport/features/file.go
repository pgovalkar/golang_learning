package features

import (
	"fmt"
	"golang_learning/jsonImport/core"
	"golang_learning/jsonImport/parser"
)

func File(file_name string, outputpath string) {
	//Read the json
	data, err := core.ReadFile(file_name)
	if err != nil {
		fmt.Println("Error Rading JSON:", err)
	}
	parser.Import(data, outputpath)
}
