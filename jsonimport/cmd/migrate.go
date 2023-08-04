package cmd

import (
	"fmt"
	"jsonImport/features"
	"os"

	"github.com/spf13/cobra"
)

var dashboard, directory, file, token, outputpath string
var importCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Import Dashboard from Json to YAML Format",
	Run:   Migrate,
}

func init() {
	rootCmd.AddCommand(importCmd)
	//update name to id
	importCmd.Flags().StringVarP(&dashboard, "dashboard", "d", "", "Import from Wavefront,use dashboard option with Dashboard Name")
	importCmd.Flags().StringVarP(&token, "token", "t", "", "Wavefront token")
	importCmd.MarkFlagRequired("service")
	importCmd.Flags().StringVarP(&directory, "directory", "D", "", "For Bulk import, use directory option")
	importCmd.Flags().StringVarP(&file, "file", "f", "", "Import dashboard as file, use file option")
	importCmd.Flags().StringVarP(&outputpath, "outputpath", "o", "yaml_output", "Output directory to write yaml files")

	if dashboard != "" && directory != "" && file != "" {
		fmt.Println("Please provide exactly one of dashboard or directory or file option ")
		os.Exit(1)
	}
}


//TODO Dry run option, 
func Migrate(cmd *cobra.Command, args []string) {
	if directory != "" {
		features.Bulk(directory, outputpath)
	}
	if file != "" {
		features.File(file, outputpath)
	}
	if dashboard != "" {
		if token == "" {
			fmt.Println("Please provide token using -t option")
		}
		features.ApiRequest(dashboard, token, outputpath)
	}

}
