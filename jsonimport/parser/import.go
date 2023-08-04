package parser

import (
	"encoding/json"
	"fmt"
	"jsonImport/clients/wavefront"
	"jsonImport/core"
)

func Import(data []byte, outputpath string) {
	//unmarshal the file as per struct
	var dashboard wavefront.Dashboard
	//var dashboard wavefront.WaveFrontDashboardSpec
	err := json.Unmarshal([]byte(data), &dashboard)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)

	}
	fmt.Println(dashboard.Url)
	//Marshal the file
	output, err := json.MarshalIndent(dashboard, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	//writing sub-section key/values to file
	err = core.WriteFile(dashboard.Url, output, "json_temp")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
	}

	//Converting Json to YAML
	YamlParser(dashboard, dashboard.Url, outputpath)

}
