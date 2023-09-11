package parser

import (
	//"fmt"
	"bytes"
	"fmt"
	"golang_learning/jsonImport/clients/wavefront"
	"golang_learning/jsonImport/core"

	"gopkg.in/yaml.v3"
)

func YamlParser(dashboard wavefront.Dashboard, file_name string, outputpath string) error {
	//TODO Bulk check for required fields then parse to yaml
	// yamlData, err := yaml.Marshal(map[string]interface{}{
	// 	"apiVersion": "oac.box.com/v1alpha",
	// 	"kind":       "WaveFrontDashboard",
	// 	"metadata": map[string]string{
	// 		"name": file_name,
	// 	},

	// 	"spec": map[string]interface{}{
	// 		"displayName":                   dashboard.Name,
	// 		"description":                   dashboard.Description,
	// 		"url":                           dashboard.Url,
	// 		"displaySectionTableOfContents": dashboard.DisplaySectionTableOfContents,
	// 		"displayQueryParameters":        dashboard.DisplayQueryParameters,
	// 		"tags":                          dashboard.Tags,
	// 		"parameterDetails":              dashboard.ParameterDetails,
	// 		"sections": func() interface{} {
	// 			sections := make([]interface{}, len(dashboard.Sections))
	// 			for i, s := range dashboard.Sections {
	// 				rows := make([]interface{}, len(s.Rows))
	// 				for j, r := range s.Rows {
	// 					charts := make([]interface{}, len(r.Charts))
	// 					for k, c := range r.Charts {
	// 						sources := make([]interface{}, len(c.Sources))
	// 						for l, src := range c.Sources {
	// 							sources[l] = map[string]string{
	// 								"name":  src.Name,
	// 								"query": src.Query,
	// 							}
	// 						}
	// 						charts[k] = map[string]interface{}{
	// 							"name":    c.Name,
	// 							"type":    c.ChartSettings.Type,
	// 							"queries": sources,
	// 						}
	// 					}
	// 					rows[j] = map[string]interface{}{
	// 						"charts": charts,
	// 					}
	// 				}
	// 				sections[i] = map[string]interface{}{
	// 					"name": s.Name,
	// 					"rows": rows,
	// 				}
	// 			}
	// 			return sections
	// 		}(),
	// 	},
	// })

	// if err != nil {
	// 	return fmt.Errorf("error")
	// }
	// err = core.WriteFile(dashboard.Url, yamlData, "yaml_output")
	// if err != nil {
	// 	return fmt.Errorf("error")
	// }
	// using yaml.NewEncoder

	fmt.Println("in yaml parser")
	yamldata := wavefront.YamlDashboard{
		APIVersion: "oac.box.com/v1alpha",
		Kind:       "WaveFrontDashboard",
		Metadata: wavefront.Metadata{
			Name: dashboard.Url,
		},
		Spec: dashboard,
	}
	var b bytes.Buffer

	if err := yaml.NewEncoder(&b).Encode(yamldata); err != nil {
		return fmt.Errorf("error parsing json to yaml")
	}

	fmt.Println(b.String())
	//err := os.WriteFile(outputpath+"/"+dashboard.Url+".yaml", b.Bytes(), 0644)
	err := core.WriteFile(dashboard.Url+".yaml", b.Bytes(), ".temp")
	if err != nil {
		return fmt.Errorf("error writing file %s to path %s", dashboard.Url+".yaml", outputpath)
	}
	return nil
}

// var b bytes.Buffer

// if err := yaml.NewEncoder(&b).Encode(yamldata); err != nil {
// 	return nil, fmt.Errorf(err.Error())
// }

// fmt.Println(b.String())
