package cit

// Template Content
// See
// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html
// and
// https://github.com/awslabs/goformation/blob/d7193b5126b26a12ad1c810c7c481023cef849d8/cloudformation/template.go#L14
type Template struct {
	AWSTemplateFormatVersion string                 `json:"AWSTemplateFormatVersion,omitempty"`
	Description              string                 `json:"Description,omitempty"`
	Metadata                 map[string]interface{} `json:"Metadata,omitempty"`
	Parameters               map[string]interface{} `json:"Parameters,omitempty"`
	Rules                    map[string]interface{} `json:"Rules,omitempty"`
	Mappings                 map[string]interface{} `json:"Mappings,omitempty"`
	CDKMetadata              map[string]interface{} `json:"CDKMetadata,omitempty"`
	Conditions               map[string]interface{} `json:"Conditions,omitempty"`
	Transform                map[string]interface{} `json:"Transform,omitempty"`
	Resources                Resources              `json:"Resources,omitempty"`
	Outputs                  map[string]interface{} `json:"Outputs,omitempty"`


}

type Resources map[string]Resource

type Resource struct {
	Metadata map[string]interface{} `json:"Metadata,omitempty"`
}
