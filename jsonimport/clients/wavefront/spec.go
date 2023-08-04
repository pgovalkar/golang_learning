package wavefront
// type Severity string
// type AlertType string

// const (
// 	SeverityInfo   Severity = "INFO"
// 	SeveritySmoke  Severity = "SMOKE"
// 	SeverityWarn   Severity = "WARN"
// 	SeveritySevere Severity = "SEVERE"

// 	AlertTypeClassic   AlertType = "CLASSIC"
// 	AlertTypeThreshold AlertType = "THRESHOLD"
// )

// type ACL struct {
// 	CanView   []string `json:"canView,omitempty"`
// 	CanModify []string `json:"canModify,omitempty"`
// }

// // type WaveFrontDashboard struct {
// // 	specs.Header `json:",inline,omitempty" validate:"required"`
// // 	Spec         WaveFrontDashboardSpec `json:"spec,omitempty" validate:"required"`
// // }

// type WaveFrontDashboardSpec struct {
// 	ACL                           `json:"acl,inline,omitempty"`
// 	DisplayName                   string             `json:"displayName,omitempty" validate:"required,min=3" yaml:"name"`
// 	Description                   string             `json:"description,omitempty" validate:"required" yaml:"description"`
// 	Url                           string             `json:"url,omitempty" validate:"required,min=3"`
// 	DisplayQueryParameters        bool               `json:"displayQueryParameters,omitempty"`
// 	DisplaySectionTableOfContents bool               `json:"displaySectionTableOfContents,omitempty"`
// 	Tags                          []string           `json:"tags,omitempty"`
// 	ParameterDetails              []ParameterDetails `json:"parameterDetails,omitempty" validate:"omitempty,dive"`
// 	Sections                      []Section          `json:"sections,omitempty" validate:"required,min=1,dive"`
// }

// type ParameterDetails struct {
// 	Name                    string              `json:"name,omitempty" validate:"required"`
// 	DefaultValue            string              `json:"defaultValue,omitempty" validate:"required_unless=ParameterType SIMPLE"`
// 	SimpleValue             string              `json:"simpleValue,omitempty" validate:"required_if=ParameterType SIMPLE"`
// 	HideFromView            bool                `json:"hideFromView,omitempty"`
// 	Label                   string              `json:"label,omitempty" validate:"required"`
// 	ParameterType           string              `json:"parameterType,omitempty" validate:"required,oneof=LIST SIMPLE DYNAMIC"`
// 	ValuesToReadableStrings *map[string]*string `json:"valuesToReadableStrings,omitempty" validate:"required_unless=ParameterType SIMPLE,omitempty,min=1"`
// 	TagKey                  string              `json:"tagKey,omitempty"`
// 	QueryValue              string              `json:"queryValue,omitempty" validate:"required_if=ParameterType DYNAMIC"`
// 	DynamicFieldType        string              `json:"dynamicFieldType,omitempty" validate:"omitempty,oneof=SOURCE SOURCE_TAG METRIC_NAME TAG_KEY MATCHING_SOURCE_TAG"`
// }

// type Section struct {
// 	Name string `json:"name,omitempty"`
// 	Rows []Row  `json:"rows,omitempty" validate:"required,min=1,dive"`
// }

// type Row struct {
// 	Charts []Chart `json:"charts,omitempty" validate:"required,min=1,dive"`
// }

// type Chart struct {
// 	Name                 string  `json:"name,omitempty" validate:"required"`
// 	Type                 string  `json:"type,omitempty" validate:"required,oneof=line scatterplot stacked-area stacked-column table markdown-widget sparkline globe nodemap top-k histogram heatmap gauge pie logs-table geo-map"`
// 	Units                string  `json:"units,omitempty"`
// 	Queries              []Query `json:"queries,omitempty" validate:"required_unless=Type markdown-widget,omitempty,min=1,dive"`
// 	PlainMarkdownContent string  `json:"plainMarkdownContent,omitempty" validate:"required_if=Type markdown-widget"`
// }

// type Query struct {
// 	Name  string `json:"name,omitempty" validate:"required"`
// 	Query string `json:"query,omitempty" validate:"required,terraformescaped"`
// 	Hide  bool   `json:"hide,omitempty"`
// }


