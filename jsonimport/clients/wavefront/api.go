package wavefront

const (
	Lte ComparisonOperator = "lte"
	Gte ComparisonOperator = "gte"
	Lt  ComparisonOperator = "lt"
	Gt  ComparisonOperator = "gt"
	Eq  ComparisonOperator = "eq"
	Neq ComparisonOperator = "neq"

	SeverityInfo   Severity = "INFO"
	SeveritySmoke  Severity = "SMOKE"
	SeverityWarn   Severity = "WARN"
	SeveritySevere Severity = "SEVERE"

	AlertTypeClassic   AlertType = "CLASSIC"
	AlertTypeThreshold AlertType = "THRESHOLD"
)

type (
	ComparisonOperator string
	Severity           string
	AlertType          string
)

type DashboardResponse struct {
	Response Dashboard `json:"response"`
}

type ResponseStatus struct {
	Status Status `json:"status"`
}
type Status struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Code    int16  `json:"code"`
}

//TODO
// type AlertResponse struct {
// 	Response Dashboard `json:"response"`
// }

type ACL struct {
	CanView   []interface{} `json:"canView"`
	CanModify []string      `json:"canModify"`
}

type Tags struct {
	CustomerTags []string `json:"customerTags,omitempty"`
}
type Dashboard struct {
	Name string `json:"name"`
	ACL  struct {
		CanView   []interface{} `json:"canView"`
		CanModify []string      `json:"canModify"`
	} `json:"acl,omitempty"`
	Description                   string    `json:"description,omitempty"`
	Url                           string    `json:"url"`
	EventQuery                    string    `json:"eventQuery,omitempty"`
	DefaultTimeWindow             string    `json:"defaultTimeWindow,omitempty"`
	DisplayDescription            bool      `json:"displayDescription,omitempty"`
	DisplaySectionTableOfContents bool      `json:"displaySectionTableOfContents,omitempty"`
	DisplayQueryParameters        bool      `json:"displayQueryParameters,omitempty"`
	Tags                          Tags      `json:"tags,omitempty"`
	Sections                      []Section `json:"sections"`
	//ParameterDetails []ParameterDetails `json:"parameterDetails,omitempty" yaml:"parameterDetails,omitempty"`
	ParameterDetails map[string]struct {
		ParameterType           string            `json:"parameterType"`
		Order                   int               `json:"order"`
		DefaultValue            string            `json:"defaultValue"`
		HideFromView            bool              `json:"hideFromView"`
		Label                   string            `json:"label"`
		ValuesToReadableStrings map[string]string `json:"valuesToReadableStrings"`
		TagKey                  string            `json:"tagKey,omitempty"`
		QueryValue              string            `json:"queryValue,omitempty"`
		DynamicFieldType        string            `json:"dynamicFieldType,omitempty"`
	}
}

type ParameterDetails struct {
	ParameterType           string            `json:"parameterType"`
	Order                   int               `json:"order"`
	DefaultValue            string            `json:"defaultValue"`
	HideFromView            bool              `json:"hideFromView"`
	Label                   string            `json:"label"`
	ValuesToReadableStrings map[string]string `json:"valuesToReadableStrings"`
	TagKey                  string            `json:"tagKey,omitempty"`
	QueryValue              string            `json:"queryValue,omitempty"`
	DynamicFieldType        string            `json:"dynamicFieldType,omitempty"`
}
type Section struct {
	Name string `json:"name"`
	Rows []Row  `json:"rows"`
}

type Row struct {
	Charts []Chart `json:"charts"`
}

type Chart struct {
	ChartSettings ChartSettings `json:"chartSettings"`
	Name          string        `json:"name"`
	Sources       []Source      `json:"sources" yaml:"queries"`
	Summarization string        `json:"summarization,omitempty"`
	//Description   string        `json:"description,omitempty"`
	Unit string `json:"units,omitempty"`
}

type ChartSettings struct {
	Type string `json:"type"`
	// AutoColumnTags            bool   `json:"autoColumnTags,omitempty"`
	// SparklineDecimalPrecision int    `json:"sparklineDecimalPrecision,omitempty"`
	// DefaultSortColumn         string `json:"defaultSortColumn,omitempty"`
	// SparklineDisplayColor     string `json:"sparklineDisplayColor,omitempty"`
	// SparklineLineColor        string `json:"sparklineLineColor,omitempty"`
	// SparklineFillColor        string `json:"sparklineFillColor,omitempty"`
}

type Source struct {
	Name     string `json:"name"`
	Query    string `json:"query"`
	Disabled bool   `json:"disabled,omitempty"`
}

type YamlDashboard struct {
	APIVersion string    `yaml:"apiVersion"`
	Kind       string    `yaml:"kind"`
	Metadata   Metadata  `yaml:"metadata"`
	Spec       Dashboard `yaml:"spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
}

// Alerts

type AlertResponse struct {
	Response AlertResponseSpec `json:"response"`
}

type AlertResponseSpec struct {
	Items []WFAlertSpec `json:"items"`
}

type WFAlertSpec struct {
	AlertType string `json:"alertType"`
	// Condition                  string   `json:"condition"`
	// Name                       string   `json:"name"`
	// DisplayExpression          string   `json:"displayExpression"`
	//DisplayName string    `json:"displayName,omitempty" validate:"required,min=5"`
	//Query       string    `json:"query,omitempty" validate:"required,terraformescaped"`
	//Condition           *Condition          `json:"condition,omitempty" validate:"required_without=SeverityConditions"`
	//SeverityConditions  *SeverityConditions `json:"severityConditions,omitempty" validate:"required_without=Condition"`
}


// type WaveFrontAlertBasic struct {
// 	ACL                   `json:",inline,omitempty"`
// 	AlertType             AlertType             `json:"alertType,omitempty" validate:"omitempty,oneof=CLASSIC THRESHOLD"`
// 	Targets               []string              `json:"targets,omitempty" validate:"omitempty,dive,min=5"`
// 	ThresholdTargets      *AlertThresholdTarget `json:"thresholdTargets,omitempty" validate:"omitempty"`
// 	AdditionalInformation string                `json:"additionalInformation,omitempty"`
// 	Severity              Severity              `json:"severity,omitempty" validate:"omitempty,oneof=INFO SMOKE WARN SEVERE"`
// 	Minutes               int                   `json:"minutes,omitempty" validate:"min=0"`
// 	ResolveAfterMinutes   int                   `json:"resolveAfterMinutes,omitempty" validate:"min=0"`
// 	Tags                  []string              `json:"tags,omitempty"`
// }


// type SeverityConditions struct {
// 	Info   *Condition `json:"info,omitempty" validate:"required_without_all=Warn Severe"`
// 	Warn   *Condition `json:"warn,omitempty" validate:"required_without_all=Info Severe"`
// 	Severe *Condition `json:"severe,omitempty" validate:"required_without_all=Info Warn"`
// }

// type Condition struct {
// 	Operator  ComparisonOperator `json:"operator,omitempty" validate:"required,oneof=lte gte lt gt eq neq"`
// 	Threshold float64                  `json:"threshold,omitempty"`
// }
