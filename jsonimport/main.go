package main

import (
	"fmt"
	"jsonImport/features"
	//"jsonImport/cmd"
	//"jsonImport/core"
	//"jsonImport/clients/wavefront"
)

func main() {
	//cmd.Execute()
	// f, err := core.ReadFile("/Users/pgovalkar/testing/gcs-dashboard.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = core.WriteFile("gcs-dahboard.yaml", f, "")

	const (
		id           string = "Instances"
		tag          string = "box-oac.dashboard"
		endpoint     string = "alert/"
		defaultWFUrl string = "https://box.wavefront.com/api/v2/"
		token        string = "887ae0e7-e64a-457e-bb48-ea8ea7516256"
	)
	// c, err := wavefront.New(token, defaultWFUrl)
	// //c.GetDashboard(id)
	// c.GetAlert(tag)
	err := features.ApiRequest(id, token, "yaml_output")
	if err != nil {
		fmt.Println()
	}
	//features.ApiRequest(id,token,"yaml_output")

}
