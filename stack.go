package cit

type Template struct {
	AWSTemplateFormatVersion string                 `json:"AWSTemplateFormatVersion,omitempty"`
	Transform                map[string]interface{} `json:"Transform,omitempty"`
	Description              string                 `json:"Description,omitempty"`
	Metadata                 map[string]interface{} `json:"Metadata,omitempty"`
	Parameters               map[string]interface{} `json:"Parameters,omitempty"`
	Mappings                 map[string]interface{} `json:"Mappings,omitempty"`
	Conditions               map[string]interface{} `json:"Conditions,omitempty"`
	Resources                Resources              `json:"Resources,omitempty"`
	Rules                    map[string]interface{} `json:"Rules,omitempty"`
	Outputs                  map[string]interface{} `json:"Outputs,omitempty"`
}

type Resources map[string]Resource

type Resource struct {
	Metadata map[string]string `json:"Metadata,omitempty"`
}
